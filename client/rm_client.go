package client

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/ljb021027/tcc/proto/go"
	"github.com/ljb021027/tcc/util"
	"google.golang.org/grpc"
)

type RmProxyClient struct {
	rc         tcc.RmClient
	rmResource *tcc.RmResource
}

func NewRmProxyClient(rmResource *tcc.RmResource) (*RmProxyClient, error) {
	cli, err := util.GloalGrpcCliCache.GetCli(rmResource.Uri)
	if err != nil {
		return nil, err
	}
	return &RmProxyClient{
		rc:         tcc.NewRmClient(cli),
		rmResource: rmResource,
	}, nil

}

func (r *RmProxyClient) GetRmResource() *tcc.RmResource {
	return r.rmResource
}

func (r *RmProxyClient) Prepare(ctx context.Context, in *tcc.Xid, opts ...grpc.CallOption) (*empty.Empty, error) {
	return r.rc.Prepare(ctx, in, opts...)
}

func (r *RmProxyClient) Commit(ctx context.Context, in *tcc.Branch, opts ...grpc.CallOption) (*empty.Empty, error) {
	return r.rc.Commit(ctx, in, opts...)
}

func (r *RmProxyClient) Cancel(ctx context.Context, in *tcc.Branch, opts ...grpc.CallOption) (*empty.Empty, error) {
	return r.rc.Cancel(ctx, in, opts...)
}
