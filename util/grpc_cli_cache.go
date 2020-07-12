package util

import (
	"context"
	"errors"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

var GloalGrpcCliCache *GrpcCliCache

//GrpcCliCache
type GrpcCliCache struct {
	//key-addr,value-grpcConn
	grpcConnMap *SyncMap
	initTimeOut int
	dialTimeOut int
}

func InitGrpcCli() {
	GloalGrpcCliCache = NewGrpcCli(5, 5)
}

var kacp = keepalive.ClientParameters{
	Time:                10 * time.Second, // send pings every 10 seconds if there is no activity
	Timeout:             time.Second,      // wait 1 second for ping ack before considering the connection dead
	PermitWithoutStream: true,             // send pings even without active streams
}

//NewGrpcCli
func NewGrpcCli(initTimeOut, dialTimeOut int) *GrpcCliCache {
	return &GrpcCliCache{
		grpcConnMap: &SyncMap{},
		initTimeOut: initTimeOut,
		dialTimeOut: dialTimeOut,
	}
}

//GetCli 通过target address 和初始化传入的func 建立grpc client并缓存起来
func (gc *GrpcCliCache) GetCli(serviceUri string) (*grpc.ClientConn, error) {
	service, err := parseService(serviceUri)
	if err != nil {
		return nil, err
	}
	value, err := gc.grpcConnMap.ComputeIfAbsent(serviceUri, func(key interface{}) (i interface{}, e error) {
		ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(gc.initTimeOut)*time.Second)
		defer cancelFunc()
		timeOut := gc.dialTimeOut
		if service.dialTimeOut != 0 {
			timeOut = service.dialTimeOut
		}
		optiong := GetUnaryGrpcClientBaseOptiong(time.Duration(timeOut) * time.Second)
		optiong = append(optiong, grpc.WithKeepaliveParams(kacp))
		conn, err := grpc.DialContext(ctx, service.addr, optiong...)
		if err != nil {
			return nil, err
		}
		return conn, nil
	})
	if err != nil {
		return nil, err
	}
	cli, isok := value.(*grpc.ClientConn)
	if isok == false {
		return nil, errors.New("ClientConn type error")
	}
	return cli, nil
}

//Delete
func (gc *GrpcCliCache) Delete(key string) {
	gc.grpcConnMap.Cache.Delete(key)
}

//Destory
func (gc *GrpcCliCache) Destory() {
	gc.grpcConnMap.Cache.Range(func(key, value interface{}) bool {
		conn := value.(*grpc.ClientConn)
		if conn != nil {
			conn.Close()
		}
		return false
	})
}

func GetUnaryGrpcClientBaseOptiong(timeout time.Duration) []grpc.DialOption {
	return []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(
			grpc_middleware.ChainUnaryClient(
				ClientUnaryTimeOut(timeout),
			)),
	}
}

//ClientUnaryTimeOut
func ClientUnaryTimeOut(timeout time.Duration) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		ctx, cancelFunc := context.WithTimeout(ctx, timeout)
		defer cancelFunc()
		err := invoker(ctx, method, req, reply, cc, opts...)
		return err
	}
}
