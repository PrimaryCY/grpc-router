package rpc

import (
	"grpc-route/coordinate"
)

type ClientRpc struct {
	CoordinateManager coordinate.BaseCoordinateManager
	RpcManager BaseRpcManager
}

func NewClient(rpcManager BaseRpcManager, coordinateManager coordinate.BaseCoordinateManager) *ClientRpc {
	return &ClientRpc{
		CoordinateManager: coordinateManager,
		RpcManager:rpcManager,
	}
}

func (c *ClientRpc)RpcCallBu(buName string, request *Request)(*Response, error)  {
	service, err := c.CoordinateManager.GetRpcService(buName)
	if err != nil{
		return nil, err
	}
	return c.RpcManager.RpcBUCall(service.Ip, service.Port, request)

}

