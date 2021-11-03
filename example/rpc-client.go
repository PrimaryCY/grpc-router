package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"grpc-route/coordinate"
	"grpc-route/rpc"
)

func Client() *rpc.ClientRpc {
	manager := rpc.NewManagerRpc()

	conf := api.DefaultConfig()
	conf.Address = "http://127.0.0.1:8500"
	cd, err := coordinate.NewConsulCoordinateManager(conf)
	if err != nil{
		panic(err)
	}
	return rpc.NewClient(manager, cd)
}


func main() {
	client := Client()
	response, err := client.RpcCallBu("example", &rpc.Request{
		Header: map[string]interface{}{
			"token": "==!jxz",
		},
		Params: map[string]interface{}{
			"data": map[string]string{
				"name": "pony",
			},
		},
		Files:    []byte{1, 2},
		FuncName: "testFuncName",
		Package:  "pkg",
	})
	if err != nil{
		panic(err)
	}
	fmt.Printf("response: %#v", response)
	
}
