package rpc

import (
	"grpc-route/coordinate"
)

type ClientRpc struct {
	coordinateManager coordinate.BaseCoordinateManager
	rpcManager BaseRpcManager
}

func NewClient(rpcManager BaseRpcManager, coordinateManager coordinate.BaseCoordinateManager) *ClientRpc {
	return &ClientRpc{
		coordinateManager: coordinateManager,
		rpcManager:rpcManager,
	}
}

func (c *ClientRpc)RpcCallBu(buName string, request *Request)(*Response, error)  {
	service, err := c.coordinateManager.GetRpcService(buName)
	if err != nil{
		return nil, err
	}
	return c.rpcManager.RpcBUCall(service.Ip, service.Port, request)

}

