package http

import "github.com/PrimaryCY/grpc-router/coordinate"

type ServiceHttp struct {
	CoordinateManager coordinate.BaseCoordinateManager
	ready bool
	serviceId string
}

func NewServiceHttp(coordinateManager coordinate.BaseCoordinateManager) *ServiceHttp  {
	return &ServiceHttp{
		CoordinateManager: coordinateManager,
		ready:             false,
	}
}

func (s *ServiceHttp)Register(coordinateService *coordinate.Service)(id string, err error){
	if id, err = s.CoordinateManager.RegisterHttp(coordinateService); err != nil{
		return
	}
	s.ready = true
	return
}

func (s *ServiceHttp)DeRegister(id string)(err error){
	if err = s.CoordinateManager.DeregisterRpc(id); err != nil{
		return
	}
	s.ready = false
	s.serviceId = ""
	return
}