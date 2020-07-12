package client

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	tcc "github.com/ljb021027/tcc/proto/go"
	"github.com/ljb021027/tcc/util"
)

type Tm interface {
	StartTransaction(ctx context.Context) error
}

type TmService struct {
	tc  tcc.TcClient
	rms []*RmProxyClient
}

func NewTmService(tc tcc.TcClient, rms []*RmProxyClient) Tm {
	return &TmService{
		tc:  tc,
		rms: rms,
	}
}

//StartTransaction 开启一个全局事务
func (tmService *TmService) StartTransaction(ctx context.Context) error {
	log := util.Action(ctx, "tm")
	//tm -> tc 申请全局事务 并拿到xid
	xid, err := tmService.tc.NewGlobalTransaction(ctx, &empty.Empty{})
	if err != nil {
		log.Errorf("NewGlobalTransaction error:%s", err.Error())
		return err
	}
	log.Infof("xid:%+v NewGlobalTransaction success", xid)
	//向所有rm传播xid
	for _, rm := range tmService.rms {
		//rm 收到xid，会向tc注册这个xid下的分支事务
		_, err := rm.Prepare(ctx, xid)
		if err != nil {
			log.Errorf("xid:%+v rm:%+v prepare error:%s", xid, rm.GetRmResource(), err.Error())
			_, err := tmService.tc.RollBack(ctx, xid)
			if err != nil {
				log.Errorf("xid:%+v rm:%+v RollBack error:%s", xid, rm.GetRmResource(), err.Error())
			}
			return err
		}
		log.Infof("xid:%+v rm:%+v prepare success", xid, rm.GetRmResource())
	}

	//commit
	_, err = tmService.tc.Commit(ctx, xid)
	if err != nil {
		log.Errorf("xid:%+v Commit error:%s", xid, err.Error())
		return err
	}
	log.Infof("xid:%+v Commit success", xid)
	return nil
}
