package ws

import "grpc-route/coordinate"

type ServiceWs struct {
	coordinateManager coordinate.BaseCoordinateManager
	ready bool
	serviceId string
}

func NewServiceSock(coordinateManager coordinate.BaseCoordinateManager) *ServiceWs  {
	return &ServiceWs{
		coordinateManager: coordinateManager,
		ready:             false,
	}
}

func (s *ServiceWs)Register(coordinateService *coordinate.Service)(id string, err error){
	if id, err = s.coordinateManager.RegisterWs(coordinateService); err != nil{
		return
	}
	s.ready = true
	return
}

func (s *ServiceWs)DeRegister(id string)(err error){
	if err = s.coordinateManager.DeregisterRpc(id); err != nil{
		return
	}
	s.ready = false
	s.serviceId = ""
	return
}