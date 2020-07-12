package test

import (
	"fmt"

	"github.com/ljb021027/tcc/proto/go"
)

type ServiceA struct {
	rmResource *tcc.RmResource
}

func NewServiceA(rmResource *tcc.RmResource) *ServiceA {
	return &ServiceA{rmResource: rmResource}
}

func (s *ServiceA) Try(param *tcc.Param) error {
	fmt.Println("serviceA try")
	return nil
}

func (s *ServiceA) Commit(param *tcc.Param) error {
	fmt.Println("serviceA Commit")
	return nil
}

func (s *ServiceA) Cancel(param *tcc.Param) error {
	fmt.Println("serviceA Cancel")
	return nil
}

func (s *ServiceA) GetRmResource() *tcc.RmResource {
	return s.rmResource
}

type ServiceB struct {
	rmResource *tcc.RmResource
}

func NewServiceB(rmResource *tcc.RmResource) *ServiceB {
	return &ServiceB{rmResource: rmResource}
}

func (s *ServiceB) Try(param *tcc.Param) error {
	fmt.Println("serviceB try")
	return nil
}

func (s *ServiceB) Commit(param *tcc.Param) error {
	fmt.Println("serviceB Commit")
	return nil
}

func (s *ServiceB) Cancel(param *tcc.Param) error {
	fmt.Println("serviceB Cancel")
	return nil
}

func (s *ServiceB) GetRmResource() *tcc.RmResource {
	return s.rmResource
}
