package rpc

import (
	"errors"
	"grpc-route/coordinate"
)

type ServiceRpc struct {
	coordinateManager coordinate.BaseCoordinateManager
	rpcManager BaseRpcManager
	ready bool
	serviceId string
}

func NewServiceRpc(rpcManager BaseRpcManager, coordinateManager coordinate.BaseCoordinateManager) *ServiceRpc  {
	return &ServiceRpc{
		coordinateManager: coordinateManager,
		rpcManager:        rpcManager,
		ready:             false,
	}
}


func (s *ServiceRpc)Register(coordinateService *coordinate.Service)(id string, err error){
	if id, err = s.coordinateManager.RegisterRpc(coordinateService); err != nil{
		return
	}
	s.ready = true
	s.serviceId = id
	return
}

func (s *ServiceRpc)DeRegister(id string)(err error){
	if err = s.coordinateManager.DeregisterRpc(id); err != nil{
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
	return s.rpcManager.Walk(port)
}

func (s *ServiceRpc)AfterRequest(f func(*Context)) {
	s.rpcManager.addAfterRequest(f)
}

func (s *ServiceRpc)BeforeRequest(f func(*Context)) {
	s.rpcManager.addBeforeRequest(f)
}

func (s *ServiceRpc)AddHandler(pkg string, name string, f func(*Context)) {
	s.rpcManager.Bound(pkg, name, f)
}
