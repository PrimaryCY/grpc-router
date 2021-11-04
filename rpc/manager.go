package rpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"github.com/PrimaryCY/grpc-router/rpc/proto"
	"net"
	"sync"
)

type BaseRpcManager interface {
	addAfterRequest(func(*Context))
	addBeforeRequest(func(*Context))

	RpcBUCall(buName string, buPort int, request *Request)(*Response, error)
	Walk(int) error
	Bound(pkg string, name string, f func(*Context))
}


type ManagerRpc struct {
	proto.UnimplementedRouterServer
	GRpcServer *grpc.Server
	BeforeRequest []func(*Context)
	AfterRequest []func(*Context)
	FuncTable *sync.Map
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

func (m *ManagerRpc)addAfterRequest(f func(*Context)) {
	m.AfterRequest = append(m.AfterRequest, f)
}

func (m *ManagerRpc)addBeforeRequest(f func(*Context)) {
	m.BeforeRequest = append(m.BeforeRequest, f)
}


func (m *ManagerRpc)Walk(port int) error{
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		return err
	}

	proto.RegisterRouterServer(m.GRpcServer, m)
	grpc_health_v1.RegisterHealthServer(m.GRpcServer, health.NewServer())

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
	pkgMapI, pkgHas := m.FuncTable.Load(ctx.Request.RawRequest.Package)
	pkgMap := pkgMapI.(*sync.Map)
	Func, FuncHas := pkgMap.Load(ctx.Request.RawRequest.FunctionName)
	if pkgHas && FuncHas {
		for _, f := range m.BeforeRequest{
			f(ctx)
		}
		if ctx.Res.StatusCode != 0{
			return
		}
		Func.(func(*Context))(ctx)

	} else{
		ctx.Res.StatusCode = RPC_404_NOT_FOUND
		ctx.Res.Data = map[string]interface{}{
			"err": fmt.Sprintf("function or package not fount, package: %s, function: %s", ctx.Request.RawRequest.Package, ctx.Request.RawRequest.FunctionName),
		}
	}

	for _, f := range m.AfterRequest{
		f(ctx)
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

	c := proto.NewRouterClient(conn)
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

func (m *ManagerRpc) Bound(pkg string, name string, f func(*Context)){
	/*
	{
		pkg: {
			"name": func(){}
			}
	}
	 */
	pkgMap := &sync.Map{}
	if m.FuncTable == nil{
		m.FuncTable = pkgMap
	}

	funcMap := &sync.Map{}
	_, ok := pkgMap.Load(pkg)
	if !ok{
		pkgMap.Store(pkg, funcMap)
	}

	funcMap.Store(name, f)
}
