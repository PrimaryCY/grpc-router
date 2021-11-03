package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"grpc-route/coordinate"
	"grpc-route/rpc"
	"grpc-route/tools"
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

	_, _ = server.Register(&coordinate.Service{
		Ip:   tools.LocalIP(),
		Port: port,
		Name: "example",
	})

	fmt.Printf("start server, port: %d \n", port)
	if err := server.Walk(port); err!=nil{
		panic(err)
	}


}
