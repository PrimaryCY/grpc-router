package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"github.com/PrimaryCY/grpc-router/coordinate"
	"github.com/PrimaryCY/grpc-router/rpc"
	"github.com/PrimaryCY/grpc-router/tools"
)

func Server() *rpc.ServiceRpc {
	manager := rpc.NewManagerRpc()

	conf := api.DefaultConfig()
	conf.Address = "http://127.0.0.1:8500"
	cd, err := coordinate.NewConsulCoordinateManager(conf)
	if err != nil{
		panic(err)
	}
	return rpc.NewServiceRpc(manager, cd)
}

func main() {

	server := Server()
	port := 5000

	server.AddHandler("pkg", "testFuncName", func(context *rpc.Context) {
		fmt.Printf("receive params: %v \n", context.Request.Params)
		fmt.Printf("receive header: %v \n", context.Request.Header)
		fmt.Printf("receive files: %v \n", context.Request.Files)
		context.Response(200, map[string]interface{}{}, map[string]interface{}{
			"message": 200,
		})
	})

	server.AddHandler("pkg", "testFuncName2", func(context *rpc.Context) {
		fmt.Printf("receive params: %v \n", context.Request.Params)
		fmt.Printf("receive header: %v \n", context.Request.Header)
		fmt.Printf("receive files: %v \n", context.Request.Files)
		context.Response(200, map[string]interface{}{}, map[string]interface{}{
			"message": 200,
		})
	})

	server.BeforeRequest(func(context *rpc.Context) {
		fmt.Printf("before request! context: %v \n", context)
	})

	server.AfterRequest(func(context *rpc.Context) {
		fmt.Printf("after request! context: %v \n", context)
	})

	_, _ = server.Register(&coordinate.Service{
		Ip:   tools.LocalIP(),
		Port: port,
		Name: "example",
	})

	fmt.Printf("start consul server, ip: %s port: %d \n", tools.LocalIP(), port)
	if err := server.Walk(port); err!=nil{
		panic(err)
	}


}
