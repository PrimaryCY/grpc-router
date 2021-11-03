package coordinate

import (
	"strings"
)

type KubernetesCoordinateManager struct {
	RpcPort int
}

func NewKubernetesCoordinateManager(rpcPort int) (*KubernetesCoordinateManager, error){
	return &KubernetesCoordinateManager{
		RpcPort:rpcPort,
	}, nil
}

func (c *KubernetesCoordinateManager)RegisterHttp(service *Service)(id string, err error)  {
	return "", nil
}

func (c *KubernetesCoordinateManager)RegisterRpc(service *Service)(id string, err error)  {
	return "", nil
}

func (c *KubernetesCoordinateManager)RegisterWs(service *Service)(id string, err error)  {
	return "", nil

}

func (c *KubernetesCoordinateManager)DeregisterHttp(id string) error {
	return nil
}

func (c *KubernetesCoordinateManager)DeregisterRpc(id string) error {
	return nil
}

func (c *KubernetesCoordinateManager)DeregisterWs(id string) error {
	return nil
}

func (c *KubernetesCoordinateManager) GetServices() ([]*Service, error) {
	return nil, nil
}

func (c *KubernetesCoordinateManager) GetHttpService(name string) (*Service, error) {
	return &Service{
		Id:    "",
		Ip:    strings.ToLower(name),
		Port:  0,
		Name:  name,
		Raw:   nil,
	},nil
}

func (c *KubernetesCoordinateManager) GetWsService(name string) (*Service, error) {
	return &Service{
		Id:    "",
		Ip:    strings.ToLower(name),
		Port:  0,
		Name:  name,
		Raw:   nil,
	},nil
}

func (c *KubernetesCoordinateManager) GetRpcService(name string) (*Service, error) {
	return &Service{
		Id:    "",
		Ip:    strings.ToLower(name),
		Port:  c.RpcPort,
		Name:  name,
		Raw:   nil,
	},nil
}