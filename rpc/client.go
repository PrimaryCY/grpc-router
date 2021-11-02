package rpc

import "grpc-route/coordinate"

type Client struct {
	coordinateManager coordinate.BaseCoordinateManager
	rpcManager BaseRpcManager
}

func NewClient(rpcManager BaseRpcManager, coordinateManager coordinate.BaseCoordinateManager) *Client {
	return &Client{
		coordinateManager: coordinateManager,
		rpcManager:rpcManager,
	}
}

func (c *Client)RpcCallBu(buName string, request *Request)(*Response, error)  {
	service, err := c.coordinateManager.GetRpcService(buName)
	if err != nil{
		return nil, err
	}
	return c.rpcManager.RpcBUCall(service.Ip, service.Port, request)

}

