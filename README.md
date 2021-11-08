# grpc-router
让程序内部像使用http的方式一样使用grpc,
grpc-router基于google开发的grpc开发的一个grpc路由工具,目前还有python版本的grpc-router版本，但是尚未发布

###### Install
`go get github.com/PrimaryCY/grpc-router`

## Quick Start

**代码示例:**

### grpc-router-server:

```go
package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"github.com/PrimaryCY/grpc-router/coordinate"
	"github.com/PrimaryCY/grpc-router/rpc"
	"github.com/PrimaryCY/grpc-router/tools"
)

func main() {

    manager := rpc.NewManagerRpc()

	conf := api.DefaultConfig()
	conf.Address = "http://127.0.0.1:8500"
	cd, err := coordinate.NewConsulCoordinateManager(conf)
	if err != nil{
		panic(err)
	}
	server := rpc.NewServiceRpc(manager, cd)
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

	fmt.Printf("start consul server, ip: %s port: %d \n", tools.LocalIP(), port)
	if err := server.Walk(port); err!=nil{
		panic(err)
	}

}

```

### grpc-router-client:

```go
package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"github.com/PrimaryCY/grpc-router/coordinate"
	"github.com/PrimaryCY/grpc-router/rpc"
)

func main() {
    manager := rpc.NewManagerRpc()

	conf := api.DefaultConfig()
	conf.Address = "http://127.0.0.1:8500"
	cd, err := coordinate.NewConsulCoordinateManager(conf)
	if err != nil{
		panic(err)
	}
	client := rpc.NewClient(manager, cd)
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
        fmt.Println(err)
    }
    fmt.Printf("response: %#v", response)
}

```
                            
                      