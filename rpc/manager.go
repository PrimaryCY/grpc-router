package rpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"grpc-route/rpc/proto"
	"net"
	"sync"
)

type BaseRpcManager interface {
	Walk(int, struct{})
	Bound(pkg string, name string, f func())
}


type ManagerRpc struct {
	proto.UnimplementedRouteServer
	GRpcServer *grpc.Server
	funcTable map[string]map[string]func(*Context)
}


// manager rpc类，懒汉加锁单例模式
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
			GRpcServer: grpc.NewServer(opts...),
		}
	}

	return mgr
}

func (m *ManagerRpc)Walk(port int, service interface{}) error{
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		return err
	}

	proto.RegisterRouteServer(m.GRpcServer, &ManagerRpc{})
	reflection.Register(m.GRpcServer)

	if err = m.GRpcServer.Serve(lis); err != nil {
		return err
	}
	return nil
}

// 接收其他服务请求
func(m *ManagerRpc)RpcCallBU(ctx context.Context, req *proto.RpcRequest) (*proto.RpcResponse, error){
	r, err := LoadProtoRequest(req)
	if err != nil{
		return nil, err
	}
	c := NewContext(ctx, r)

	m.execute(c)

	response, err :=c.Res.DumpProtoResponse()
	if err != nil{
		return nil, err
	}
	return response, nil
}


func (m *ManagerRpc)execute(ctx *Context){
	_, pkgHas := m.funcTable[ctx.Request.RawRequest.Package]
	Func, FuncHas := m.funcTable[ctx.Request.RawRequest.Package][ctx.Request.RawRequest.FunctionName]
	if pkgHas && FuncHas {
		// todo: 缺少中间件逻辑
		Func(ctx)
	} else{
		ctx.Res.StatusCode = RPC_404_NOT_FOUND
		ctx.Res.Data = map[interface{}]interface{}{
			"err": fmt.Sprintf("function or package not fount, package: %s, function: %s", ctx.Request.RawRequest.Package, ctx.Request.RawRequest.FunctionName),
		}
	}

}

// 向其他服务发送请求
func (m *ManagerRpc)RpcBUCall (
	buIp string,
	buPort int,
	request *Request)(*Response, error){

	ctx := NewContext(context.Background(), request)
	protoReq, err := ctx.Request.DumpProtoRequest()
	if err != nil{
		return  nil, err
	}

	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", buIp, buPort), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	c := proto.NewRouteClient(conn)
	protoResponse, err := c.RpcCallBU(ctx.Context, protoReq)
	if err != nil{
		return nil, err
	}

	ctx.Res, err = LoadProtoResponse(protoResponse)
	if err != nil{
		return nil, err
	}

	return ctx.Res, nil
}

func (m *ManagerRpc) Bound(pkg string, name string, f func(*Context)) {
	m.funcTable[pkg][name] = f
}
