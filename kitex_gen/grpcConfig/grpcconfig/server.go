// Code generated by Kitex v0.0.5. DO NOT EDIT.
package grpcconfig

import (
	"github.com/cloudwego/kitex/server"
	"github.com/ldongpo/gconfig-client/kitex_gen/grpcConfig"
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