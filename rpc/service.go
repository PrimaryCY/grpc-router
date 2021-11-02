package rpc

import (
	"grpc-route/coordinate"
)

type Service struct {
	coordinateManager coordinate.BaseCoordinateManager
	rpcManager BaseRpcManager
	ready bool
}

func NewService(rpcManager BaseRpcManager, coordinateManager coordinate.BaseCoordinateManager) *Service  {
	return &Service{
		coordinateManager: coordinateManager,
		rpcManager:        rpcManager,
		ready:             false,
	}
}


func (s *Service)Register(coordinateService *coordinate.Service)(id string, err error){
	if id, err = s.coordinateManager.RegisterRpc(coordinateService); err != nil{
		return
	}
	s.ready = true
	return
}

func (s *Service)Walk(port int) {
	if !s.ready{
		return
	}
	s.rpcManager.Walk(port)
}

func (s *Service)AfterRequest(f func(*Request)) {
	s.rpcManager.addAfterRequest(f)
}

func (s *Service)BeforeRequest(f func(*Request)) {
	s.rpcManager.addBeforeRequest(f)
}

func (s *Service)AddHandler(pkg string, name string, f func(*Context)) {
	s.rpcManager.Bound(pkg, name, f)
}
