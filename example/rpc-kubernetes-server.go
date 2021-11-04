package main

import (
	"fmt"
	"github.com/PrimaryCY/grpc-router/coordinate"
	"github.com/PrimaryCY/grpc-router/rpc"
	"github.com/PrimaryCY/grpc-router/tools"
)

func KubernetesServer(port int) *rpc.ServiceRpc {
	manager := rpc.NewManagerRpc()

	cd, err := coordinate.NewKubernetesCoordinateManager(port)
	if err != nil{
		panic(err)
	}
	return rpc.NewServiceRpc(manager, cd)
}

func main() {
	port := 5000
	server := KubernetesServer(port)
	_, _ = server.Register(&coordinate.Service{})

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

	fmt.Printf("start kubernetes server, ip: %s port: %d \n", tools.LocalIP(), port)
	if err := server.Walk(port); err!=nil{
		panic(err)
	}


}
