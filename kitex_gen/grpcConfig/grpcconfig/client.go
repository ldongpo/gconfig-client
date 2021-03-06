// Code generated by Kitex v0.1.0. DO NOT EDIT.

package grpcconfig

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/ldongpo/gconfig-client/kitex_gen/grpcConfig"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Get(ctx context.Context, Req *grpcConfig.Request, callOptions ...callopt.Option) (r *grpcConfig.Response, err error)
	Put(ctx context.Context, Req *grpcConfig.PutRequest, callOptions ...callopt.Option) (r *grpcConfig.PutResponse, err error)
	Del(ctx context.Context, Req *grpcConfig.DelRequest, callOptions ...callopt.Option) (r *grpcConfig.DelResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kGrpcConfigClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kGrpcConfigClient struct {
	*kClient
}

func (p *kGrpcConfigClient) Get(ctx context.Context, Req *grpcConfig.Request, callOptions ...callopt.Option) (r *grpcConfig.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Get(ctx, Req)
}

func (p *kGrpcConfigClient) Put(ctx context.Context, Req *grpcConfig.PutRequest, callOptions ...callopt.Option) (r *grpcConfig.PutResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Put(ctx, Req)
}

func (p *kGrpcConfigClient) Del(ctx context.Context, Req *grpcConfig.DelRequest, callOptions ...callopt.Option) (r *grpcConfig.DelResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Del(ctx, Req)
}
