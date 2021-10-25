package rpc

import (
	"google.golang.org/grpc"
	"sync"
)

type BaseRpcManager interface {
	Walk(int, struct{})
	Bound(pkg string, name string, f func())
}


type ManagerRpc struct {
	server *grpc.Server
	funcTable map[string]map[string]func()
}


// 新建manager rpc类，采用懒汉加锁单例模式
var mgr *ManagerRpc
var mu sync.Mutex
func NewManagerRpc(option ...grpc.ServerOption) *ManagerRpc {
	if mgr == nil{
		mu.Lock()
		defer mu.Unlock()
		defaultOption := []grpc.ServerOption{
			grpc.MaxSendMsgSize(512 * 1024 * 1024),
			grpc.MaxRecvMsgSize(512 * 1024 * 1024),
		}
		opts := append(defaultOption, option...)

		mgr = &ManagerRpc{
			server: grpc.NewServer(opts...),
		}
	}

	return mgr
}

func (m *ManagerRpc)RpcCallBu(){

}

func (m *ManagerRpc) Bound(pkg string, name string, f func()) {
	m.funcTable[pkg][name] = f
}
