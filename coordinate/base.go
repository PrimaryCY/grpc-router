package coordinate

type BaseCoordinateManager interface {
	RegisterHttp(*Service) (id string, err error)
	RegisterWs(*Service) (id string, err error)
	RegisterRpc(*Service) (id string, err error)
	DeregisterHttp(string) error
	DeregisterWs(string) error
	DeregisterRpc(string) error
	GetServices() ([]*Service, error)
	GetHttpService(string) (*Service, error)
	GetRpcService(string) (*Service, error)
	GetWsService(string) (*Service, error)
}

type Service struct {
	Id    string
	Ip    string
	Port  int
	Name  string
	Raw interface{}
}

