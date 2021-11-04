package rpc

import (
	"errors"
	"fmt"
	"google.golang.org/grpc/grpclog"
	"github.com/PrimaryCY/grpc-router/coordinate"
)

type ServiceRpc struct {
	CoordinateManager coordinate.BaseCoordinateManager
	RpcManager BaseRpcManager
	ready bool
	serviceId string
}

func NewServiceRpc(rpcManager BaseRpcManager, coordinateManager coordinate.BaseCoordinateManager) *ServiceRpc  {
	return &ServiceRpc{
		CoordinateManager: coordinateManager,
		RpcManager:        rpcManager,
		ready:             false,
	}
}


func (s *ServiceRpc)Register(coordinateService *coordinate.Service)(id string, err error){
	if id, err = s.CoordinateManager.RegisterRpc(coordinateService); err != nil{
		return
	}
	s.ready = true
	s.serviceId = id
	return
}

func (s *ServiceRpc)DeRegister(id string)(err error){
	if err = s.CoordinateManager.DeregisterRpc(id); err != nil{
		return
	}
	s.ready = false
	s.serviceId = ""
	return
}

func (s *ServiceRpc)Walk(port int) error {
	if !s.ready{
		return errors.New("rpc service not ready")
	}
	return s.RpcManager.Walk(port)
}

func (s *ServiceRpc)AfterRequest(f func(*Context)) {
	s.RpcManager.addAfterRequest(f)
}

func (s *ServiceRpc)BeforeRequest(f func(*Context)) {
	s.RpcManager.addBeforeRequest(f)
}

func (s *ServiceRpc)AddHandler(pkg string, name string, f func(*Context)) {
	grpclog.Info(fmt.Sprintf("pkg: %s -> name: %s -> func: %s", pkg, name, f))
	s.RpcManager.Bound(pkg, name, f)
}
