package pkg

import (
	"net/http"
)

type Server interface {
	Route(pattern string, handlerFunc func(ctx *Context))
	Start(address string) error
}

// 结构体实现Server接口
type sdkHttpServer struct {
	Name string
}

func (s *sdkHttpServer) Route(pattern string, handlerFunc func(ctx *Context)) {
	http.HandleFunc(pattern, func(writer http.ResponseWriter, request *http.Request) {
		ctx := NewContext(writer, request)
		handlerFunc(ctx)
	})
}

func (s *sdkHttpServer) Start(address string) error {
	return http.ListenAndServe(address, nil)
}

func NewServer(name string) Server {
	return &sdkHttpServer{Name: name}
}
