package rpc

type BaseRpcService interface {
	register()
	walk()
}

type Service struct {

}