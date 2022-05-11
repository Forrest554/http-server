package main

import "net/http"

type Server interface {
	Route(pattern string, handlerFunc http.HandlerFunc)
	Start(address string) error
}

// 结构体实现Server接口
type sdkHttpServer struct {
	Name string
}

func (s *sdkHttpServer) Route(pattern string, handlerFunc http.HandlerFunc) {
	http.HandleFunc(pattern, handlerFunc)
}

func (s sdkHttpServer) Start(address string) error {
	return http.ListenAndServe(address, nil)
}

func NewServer(name string) Server {
	return &sdkHttpServer{Name: name}
}
