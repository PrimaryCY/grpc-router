package main

import (
	"fmt"
	"github.com/PrimaryCY/grpc-router/coordinate"
	"github.com/PrimaryCY/grpc-router/rpc"
)

func KubernetesClient() *rpc.ClientRpc {
	manager := rpc.NewManagerRpc()

	cd, err := coordinate.NewKubernetesCoordinateManager(5000)
	if err != nil{
		panic(err)
	}
	return rpc.NewClient(manager, cd)
}


func main() {
	client := KubernetesClient()
	for i := 0; i<1;i++  {
		response, err := client.RpcCallBu("example", &rpc.Request{
			Header: map[string]interface{}{
				"token": "==!jxz",
			},
			Params: map[string]interface{}{
				"data": map[string]int{
					"num": i,
				},
			},
			Files:    []byte{1, 2},
			FuncName: "testFuncName",
			Package:  "pkg",
		})
		if err != nil{
			fmt.Println(err)
		}
		fmt.Printf("response: %#v", response)
	}
}
