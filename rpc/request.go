package rpc

import (
	"encoding/json"
	"grpc-route/rpc/proto"
	"unsafe"
)

type Request struct {
	RawRequest 	 *proto.RpcRequest
	Header       map[string]interface{}
	Params       map[string]interface{}
	Files		 []byte
	FuncName	string
	Package 	string
}

func NewRequest(funcName string,
	pkg string,
	files []byte,
	params map[string]interface{},
	header map[string]interface{},
	req *proto.RpcRequest) *Request {
	return &Request{
		RawRequest: req,
		Header:     header,
		Params:     params,
		Files:      files,
		FuncName:	funcName,
		Package:	pkg,
	}
}

func LoadProtoRequest(req *proto.RpcRequest) (*Request, error) {
	var header map[string]interface{}
	if err := json.Unmarshal([]byte(req.Header), &header); err != nil {
		return nil, err
	}

	var params map[string]interface{}
	if err := json.Unmarshal([]byte(req.Params), &params); err != nil {
		return nil, err
	}

	return &Request{
		RawRequest: req,
		Header:     header,
		Params:     params,
		Files:      req.Files,
		FuncName:	req.FunctionName,
		Package:	req.Package,
	}, nil
}

func (r *Request)DumpProtoRequest() (*proto.RpcRequest, error){

	if r.Params == nil{
		r.Params = map[string]interface{}{}
	}
	if r.Header == nil{
		r.Header = map[string]interface{}{}
	}
	headerMarshalByte, err := json.Marshal(r.Header)
	if err != nil{
		return nil, err
	}
	paramsMarshalByte, err := json.Marshal(r.Params)
	if err != nil{
		return nil, err
	}

	return &proto.RpcRequest{
		FunctionName: r.FuncName,
		Package:      r.Package,
		Header:    	  *(*string)(unsafe.Pointer(&headerMarshalByte)),
		Params:       *(*string)(unsafe.Pointer(&paramsMarshalByte)),
		Files:        r.Files,
	}, nil

}