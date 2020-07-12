package client

import (
	"context"
	"net"
	"net/url"

	"github.com/golang/protobuf/ptypes/empty"
	tcc "github.com/ljb021027/tcc/proto/go"
	"github.com/ljb021027/tcc/util"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type RmService struct {
	tc         tcc.TcClient
	tccService TccService
}

func NewRmService(tc tcc.TcClient, tccService TccService) (*RmService, error) {
	rmService := RmService{
		tc:         tc,
		tccService: tccService,
	}
	return &rmService, nil
}

func (rmService *RmService) Listen() {
	s := grpc.NewServer()
	tcc.RegisterRmServer(s, rmService)
	parse, err := url.Parse(rmService.tccService.GetRmResource().Uri)
	if err != nil {
		log.Fatal(err)
	}
	listen, err := net.Listen("tcp", ":"+parse.Port())
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("listen:%s", parse.Port())
	s.Serve(listen)
}

func (rmService *RmService) Prepare(ctx context.Context, xid *tcc.Xid) (*empty.Empty, error) {
	branch := tcc.Branch{
		Xid:        xid,
		BranchId:   util.GenerateBranchId(),
		RmResource: rmService.tccService.GetRmResource(),
	}
	_, err := rmService.tc.RegisterBranch(ctx, &branch)
	if err != nil {
		return nil, err
	}
	err = rmService.tccService.Try(branch.Param)
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

func (rmService *RmService) Commit(ctx context.Context, branch *tcc.Branch) (*empty.Empty, error) {
	err := rmService.tccService.Commit(branch.Param)
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

func (rmService *RmService) Cancel(ctx context.Context, branch *tcc.Branch) (*empty.Empty, error) {
	err := rmService.tccService.Cancel(branch.Param)
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}
