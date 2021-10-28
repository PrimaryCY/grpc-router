package rpc

import (
	"context"
)

type Context struct {
	Context context.Context
	Request *Request
	Res *Response
}

func NewContext(ctx context.Context, req *Request) *Context {
	return &Context{
		Context:  ctx,
		Request:  req,
		Res: nil,
	}
}


func (m *Context)Response(code int, header map[interface{}]interface{}, data map[interface{}]interface{}){
	m.Res = NewResponse(code, header, data)
}