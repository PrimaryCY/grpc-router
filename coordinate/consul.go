package coordinate

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
	"strings"
)

type ConsulCoordinateManager struct {
	Client *api.Client
}

func NewConsulCoordinateManager(config *api.Config) (*ConsulCoordinateManager, error){
	client, err := api.NewClient(config)
	if err != nil{
		return nil, err
	}
	return &ConsulCoordinateManager{
		Client: client,
	}, nil
}

func (c *ConsulCoordinateManager) registerService(registration *api.AgentServiceRegistration) (id string, err error) {
	registration.Name = strings.ToLower(registration.Name)
	if err = c.Client.Agent().ServiceRegister(registration); err != nil {
		return
	}
	return
}

func (c *ConsulCoordinateManager)RegisterHttp(service *Service)(id string, err error)  {
	id = uuid.New().String()
	service.Id = id
	registration := new(api.AgentServiceRegistration)
	registration.ID = id
	registration.Name = service.Name
	registration.Port = service.Port
	registration.Tags = []string{"HTTP"}
	registration.Address = service.Ip

	registration.Check = &api.AgentServiceCheck{
		HTTP:                           fmt.Sprintf("http://%s:%d%s", registration.Address, registration.Port, "/check"),
		Timeout:                        "5s",
		Interval:                       "10s",
		DeregisterCriticalServiceAfter: "30s",
	}

	return c.registerService(registration)
}

func (c *ConsulCoordinateManager)RegisterRpc(service *Service)(id string, err error)  {
	id = uuid.New().String()
	service.Id = id
	registration := new(api.AgentServiceRegistration)
	registration.ID = id
	registration.Name = service.Name
	registration.Port = service.Port
	registration.Tags = []string{"RPC"}
	registration.Address = service.Ip

	registration.Check = &api.AgentServiceCheck{
		Timeout:                        "5s",
		Interval:                       "10s",
		DeregisterCriticalServiceAfter: "30s",
		// GRpc 支持，执行健康检查的地址，service 会传到 Health.Check 函数中
		GRPC:     fmt.Sprintf("%v:%v", service.Ip, service.Port),
	}

	return c.registerService(registration)
}

func (c *ConsulCoordinateManager)RegisterWs(service *Service)(id string, err error)  {
	id = uuid.New().String()
	registration := new(api.AgentServiceRegistration)
	registration.ID = id
	registration.Name = service.Name
	registration.Port = service.Port
	registration.Tags = []string{"WS"}
	registration.Address = service.Ip

	return c.registerService(registration)
}

func (c *ConsulCoordinateManager)DeregisterHttp(id string) error {
	return c.Client.Agent().ServiceDeregister(id)
}

func (c *ConsulCoordinateManager)DeregisterRpc(id string) error {
	return c.Client.Agent().ServiceDeregister(id)
}

func (c *ConsulCoordinateManager)DeregisterWs(id string) error {
	return c.Client.Agent().ServiceDeregister(id)
}

func (c *ConsulCoordinateManager) GetServices() ([]*Service, error) {
	services, err :=c.Client.Agent().Services()
	if err != nil{
		return nil, err
	}

	result := make([]*Service, 0, len(services))
	for _, v := range services{
		result = append(result, &Service{
			Id:    v.ID,
			Ip:    v.Address,
			Port:  v.Port,
			Name:  v.Service,
			Raw:v,
		})
	}
	return result, nil
}

func (c *ConsulCoordinateManager) GetHttpService(name string) (*Service, error) {
	services, err :=c.Client.Agent().ServicesWithFilter(fmt.Sprintf("Service == %s and Tags contains HTTP", name))
	var s *Service

	if err != nil{
		return s, err
	}

	for _, v := range services{
		s = &Service{
			Id:    v.ID,
			Ip:    v.Address,
			Port:  v.Port,
			Name:  v.Service,
			Raw: v,
		}
		break
	}
	return s, nil
}


func (c *ConsulCoordinateManager) GetWsService(name string) (*Service, error) {
	services, err :=c.Client.Agent().ServicesWithFilter(fmt.Sprintf("Service == %s and Tags contains Ws", name))
	var s *Service

	if err != nil{
		return s, err
	}

	for _, v := range services{
		s = &Service{
			Id:    v.ID,
			Ip:    v.Address,
			Port:  v.Port,
			Name:  v.Service,
			Raw: v,
		}
		break
	}
	return s, nil
}

func (c *ConsulCoordinateManager) GetRpcService(name string) (*Service, error) {
	services, err :=c.Client.Agent().ServicesWithFilter(fmt.Sprintf("Service == %s and Tags contains RPC", name))
	var s *Service

	if err != nil{
		return nil, err
	}

	for _, v := range services{
		s = &Service{
			Id:    v.ID,
			Ip:    v.Address,
			Port:  v.Port,
			Name:  v.Service,
			Raw: v,
		}
		break
	}
	return s, nil
}