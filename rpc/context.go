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
		Res: &Response{},
	}
}


func (m *Context)Response(code int, header map[string]interface{}, data map[string]interface{}){
	m.Res = NewResponse(code, header, data)
}