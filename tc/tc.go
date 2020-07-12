package tc

import (
	"context"
	"net"
	"strconv"

	"github.com/golang/protobuf/ptypes/empty"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/ljb021027/tcc/proto/go"
	"github.com/ljb021027/tcc/tc/storage"
	"github.com/ljb021027/tcc/util"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TcService struct {
	grpcCliCache *util.GrpcCliCache
	gtMap        map[string][]*tcc.Branch
}

func NewTcService() *TcService {
	return &TcService{
		gtMap:        make(map[string][]*tcc.Branch, 0),
		grpcCliCache: util.GloalGrpcCliCache,
	}
}

func (t *TcService) NewGlobalTransaction(context.Context, *empty.Empty) (*tcc.Xid, error) {
	sequences := util.GenerateSeq()
	xid := tcc.Xid{
		Sequences: sequences,
	}
	log.Infof("NewGlobalTransaction xid:%+v", xid)
	t.gtMap[sequences] = make([]*tcc.Branch, 0)
	return &xid, nil
}

func (t *TcService) RegisterBranch(ctx context.Context, branch *tcc.Branch) (*tcc.Report, error) {
	log := util.Action(ctx, "RegisterBranch")
	log.Info("RegisterBranch Branch:%+v", branch)
	report := tcc.Report{
		ReportStatus: tcc.ReportStatus_FAIL,
	}
	err := storage.SaveBranchRecord(branch.BranchId, branch.Xid.Sequences, storage.StatusTry)
	if err != nil {
		if err == storage.AlreadyExistsError {
			log.Error("already exists")
			return &report, nil
		}
		return nil, err
	}
	rmResources, err := t.getBranchs(branch.Xid)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	rmResources = append(rmResources, branch)
	report.ReportStatus = tcc.ReportStatus_SUCCESS
	log.Info("RegisterBranch Branch:%+v success", branch)
	return &report, nil
}

func (t *TcService) getBranchs(xid *tcc.Xid) ([]*tcc.Branch, error) {
	sequences := xid.Sequences
	rmResources, ok := t.gtMap[sequences]
	if !ok {
		return nil, status.Newf(codes.Unavailable, "xid:%v not open GlobalTransaction", xid).Err()
	}
	return rmResources, nil
}

func (t *TcService) Commit(ctx context.Context, xid *tcc.Xid) (*tcc.Report, error) {
	log := util.Action(ctx, "Commit")
	report := tcc.Report{
		ReportStatus: 0,
	}
	branchs, err := t.getBranchs(xid)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	for _, branch := range branchs {
		record, err := storage.QueryBranchRecord(branch.BranchId, branch.Xid.Sequences)
		if err != nil {
			return nil, err
		}
		//没有记录
		if record == 0 {
			log.Warnf("not found record branch:%+v", branch)
			continue
		}
		//已经提交
		if record == storage.StatusCommit {
			log.Warnf("already commit branch:%+v", branch)
			continue
		}
		//已经回滚??
		if record == storage.StatusCancel {
			log.Errorf("already cancel?? branch:%+v", branch)
			continue
		}

		cli, err := t.grpcCliCache.GetCli(branch.RmResource.Uri)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		rmClient := tcc.NewRmClient(cli)
		_, err = rmClient.Commit(ctx, branch)
		if err != nil {
			log.Errorf("branch:%+v commit error", branch)
			return nil, err
		}
	}
	return &report, nil
}

func (t *TcService) RollBack(ctx context.Context, xid *tcc.Xid) (*tcc.Report, error) {
	log := util.Action(ctx, "RollBack")
	branchs, err := t.getBranchs(xid)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	report := tcc.Report{
		ReportStatus: tcc.ReportStatus_FAIL,
	}
	for _, branch := range branchs {
		record, err := storage.QueryBranchRecord(branch.BranchId, branch.Xid.Sequences)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		//没找到记录 空回滚
		if record == 0 {
			log.Warnf("not found record branch:%+v", branch)
			//先插入记录
			err := storage.SaveBranchRecord(branch.BranchId, branch.Xid.Sequences, storage.StatusCancel)
			if err != nil {
				//记录又存在了
				if err == storage.AlreadyExistsError {
					log.Error("already exists")
					return &report, nil
				}
				return nil, err
			}
			continue
		}

		cli, err := t.grpcCliCache.GetCli(branch.RmResource.Uri)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		rmClient := tcc.NewRmClient(cli)
		_, err = rmClient.Cancel(ctx, branch)
		if err != nil {
			log.Errorf("branch:%+v Cancel error", branch)
			return nil, err
		}
	}
	return &report, nil
}

func (t *TcService) StartListen(port int) {
	unary := []grpc.UnaryServerInterceptor{
		grpc_validator.UnaryServerInterceptor(),
		grpc_ctxtags.UnaryServerInterceptor(),
	}
	opts := []grpc.ServerOption{
		grpc_middleware.WithUnaryServerChain(unary...),
	}
	s := grpc.NewServer(opts...)
	tcc.RegisterTcServer(s, t)
	endpoint := "0.0.0.0:" + strconv.Itoa(port)
	listen, err := net.Listen("tcp", endpoint)
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("listen:%s", endpoint)
	s.Serve(listen)
}
