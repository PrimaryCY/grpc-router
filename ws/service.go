package ws

import "github.com/PrimaryCY/grpc-router/coordinate"

type ServiceWs struct {
	CoordinateManager coordinate.BaseCoordinateManager
	ready bool
	serviceId string
}

func NewServiceSock(coordinateManager coordinate.BaseCoordinateManager) *ServiceWs  {
	return &ServiceWs{
		CoordinateManager: coordinateManager,
		ready:             false,
	}
}

func (s *ServiceWs)Register(coordinateService *coordinate.Service)(id string, err error){
	if id, err = s.CoordinateManager.RegisterWs(coordinateService); err != nil{
		return
	}
	s.ready = true
	return
}

func (s *ServiceWs)DeRegister(id string)(err error){
	if err = s.CoordinateManager.DeregisterRpc(id); err != nil{
		return
	}
	s.ready = false
	s.serviceId = ""
	return
}