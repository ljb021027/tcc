package util

import (
	"github.com/jinzhu/configor"
)

var SConfig ServiceConfig
var CConfig ClientConfig

type ServiceConfig struct {
	TcPort int    `required:"true"`
	Mysql  string `required:"true"`
}
type ClientConfig struct {
	TcAddr string `required:"true"`
}

func InitServiceConfig() error {
	err := configor.Load(&SConfig, "./conf/service.yaml")
	if err != nil {
		return err
	}
	return nil
}

func InitClientConfig() error {
	err := configor.Load(&CConfig, "./conf/service.yaml")
	if err != nil {
		return err
	}
	return nil
}
