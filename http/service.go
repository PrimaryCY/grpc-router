package http

import "grpc-route/coordinate"

type ServiceHttp struct {
	coordinateManager coordinate.BaseCoordinateManager
	ready bool
	serviceId string
}

func NewServiceHttp(coordinateManager coordinate.BaseCoordinateManager) *ServiceHttp  {
	return &ServiceHttp{
		coordinateManager: coordinateManager,
		ready:             false,
	}
}

func (s *ServiceHttp)Register(coordinateService *coordinate.Service)(id string, err error){
	if id, err = s.coordinateManager.RegisterHttp(coordinateService); err != nil{
		return
	}
	s.ready = true
	return
}

func (s *ServiceHttp)DeRegister(id string)(err error){
	if err = s.coordinateManager.DeregisterRpc(id); err != nil{
		return
	}
	s.ready = false
	s.serviceId = ""
	return
}