package rpc

import (
	"encoding/json"
	"grpc-route/rpc/proto"
	"unsafe"
)

type Response struct {
	StatusCode int
	Header	map[interface{}]interface{}
	Data	map[interface{}]interface{}
}


func NewResponse(code int, header map[interface{}]interface{}, data map[interface{}]interface{}) *Response{
	return &Response{
		StatusCode: code,
		Header:     header,
		Data:       data,
	}
}

func LoadProtoResponse(r *proto.RpcResponse) (*Response, error){
	var dataUnmarshal map[interface{}]interface{}
	if err := json.Unmarshal([]byte(r.Data), &dataUnmarshal); err != nil {
		return nil, err
	}
	var headerUnmarshal map[interface{}]interface{}
	if err := json.Unmarshal([]byte(r.Header), &headerUnmarshal); err != nil {
		return nil, err
	}
	return &Response{
		StatusCode: int(r.StatusCode),
		Header:     headerUnmarshal,
		Data:       dataUnmarshal,
	}, nil
}


func(r *Response)DumpProtoResponse() (*proto.RpcResponse, error){
	headerMarshalByte, err := json.Marshal(r.Header)
	if err != nil{
		return nil, err
	}
	header := *(*string)(unsafe.Pointer(&headerMarshalByte))

	dataDictByte, err := json.Marshal(r.Data)
	if err != nil{
		return nil, err
	}
	data := *(*string)(unsafe.Pointer(&dataDictByte))

	return &proto.RpcResponse{
		StatusCode: int64(r.StatusCode),
		Header:     header,
		Data:       data,
	}, nil

}

