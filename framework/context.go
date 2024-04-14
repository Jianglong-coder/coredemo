package framework

import (
	"context"
	"net/http"
	"sync"
)

// 自定义Context
type Context struct {
	request        *http.Request
	responseWriter http.ResponseWriter
	//写保护机制
	writerMux *sync.Mutex
	//是否超时标记位
	hasTimeout bool
}

func (ctx *Context) BaseContext() context.Context {
	return ctx.request.Context()
}

func (ctx *Context) Done() <-chan struct{} {
	return ctx.BaseContext().Done()
}

// 对外暴露锁
func (ctx *Context) WriterMux() *sync.Mutex {
	return ctx.writerMux
}

// 将超时标记位设置位true
func (ctx *Context) setHasTimeout() {
	ctx.hasTimeout = true
}

func (ctx *Context) HasTimeout() bool {
	return ctx.hasTimeout
}
func (ctx *Context) Json(status int, obj interface{}) error {
	if ctx.HasTimeout() {
		return nil
	}
	return nil
}
