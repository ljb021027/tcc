package client

import (
	tcc "github.com/ljb021027/tcc/proto/go"
)

type TccService interface {
	Try(param *tcc.Param) error
	Commit(param *tcc.Param) error
	Cancel(param *tcc.Param) error
	GetRmResource() *tcc.RmResource
}
