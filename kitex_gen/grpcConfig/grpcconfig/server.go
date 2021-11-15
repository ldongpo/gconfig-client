// Code generated by Kitex v0.0.5. DO NOT EDIT.
package grpcconfig

import (
	"gconfig-client/kitex_gen/grpcConfig"
	"github.com/cloudwego/kitex/server"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler grpcConfig.GrpcConfig, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}
