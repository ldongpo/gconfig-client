// Code generated by Kitex v0.0.5. DO NOT EDIT.

package grpcconfig

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	"github.com/cloudwego/kitex/pkg/streaming"
	"github.com/ldongpo/gconfig-client/kitex_gen/grpcConfig"
	"google.golang.org/protobuf/proto"
)

func serviceInfo() *kitex.ServiceInfo {
	return grpcConfigServiceInfo
}

var grpcConfigServiceInfo = newServiceInfo()

func newServiceInfo() *kitex.ServiceInfo {
	serviceName := "GrpcConfig"
	handlerType := (*grpcConfig.GrpcConfig)(nil)
	methods := map[string]kitex.MethodInfo{
		"Get": kitex.NewMethodInfo(getHandler, newGetArgs, newGetResult, false),
		"Put": kitex.NewMethodInfo(putHandler, newPutArgs, newPutResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "grpcConfig",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.0.5",
		Extra:           extra,
	}
	return svcInfo
}

func getHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(grpcConfig.Request)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(grpcConfig.GrpcConfig).Get(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *GetArgs:
		success, err := handler.(grpcConfig.GrpcConfig).Get(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*GetResult)
		realResult.Success = success
	}
	return nil
}
func newGetArgs() interface{} {
	return &GetArgs{}
}

func newGetResult() interface{} {
	return &GetResult{}
}

type GetArgs struct {
	Req *grpcConfig.Request
}

func (p *GetArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in GetArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *GetArgs) Unmarshal(in []byte) error {
	msg := new(grpcConfig.Request)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var GetArgs_Req_DEFAULT *grpcConfig.Request

func (p *GetArgs) GetReq() *grpcConfig.Request {
	if !p.IsSetReq() {
		return GetArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *GetArgs) IsSetReq() bool {
	return p.Req != nil
}

type GetResult struct {
	Success *grpcConfig.Response
}

var GetResult_Success_DEFAULT *grpcConfig.Response

func (p *GetResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in GetResult")
	}
	return proto.Marshal(p.Success)
}

func (p *GetResult) Unmarshal(in []byte) error {
	msg := new(grpcConfig.Response)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *GetResult) GetSuccess() *grpcConfig.Response {
	if !p.IsSetSuccess() {
		return GetResult_Success_DEFAULT
	}
	return p.Success
}

func (p *GetResult) SetSuccess(x interface{}) {
	p.Success = x.(*grpcConfig.Response)
}

func (p *GetResult) IsSetSuccess() bool {
	return p.Success != nil
}

func putHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(grpcConfig.PutRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(grpcConfig.GrpcConfig).Put(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *PutArgs:
		success, err := handler.(grpcConfig.GrpcConfig).Put(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*PutResult)
		realResult.Success = success
	}
	return nil
}
func newPutArgs() interface{} {
	return &PutArgs{}
}

func newPutResult() interface{} {
	return &PutResult{}
}

type PutArgs struct {
	Req *grpcConfig.PutRequest
}

func (p *PutArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, fmt.Errorf("No req in PutArgs")
	}
	return proto.Marshal(p.Req)
}

func (p *PutArgs) Unmarshal(in []byte) error {
	msg := new(grpcConfig.PutRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var PutArgs_Req_DEFAULT *grpcConfig.PutRequest

func (p *PutArgs) GetReq() *grpcConfig.PutRequest {
	if !p.IsSetReq() {
		return PutArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *PutArgs) IsSetReq() bool {
	return p.Req != nil
}

type PutResult struct {
	Success *grpcConfig.PutResponse
}

var PutResult_Success_DEFAULT *grpcConfig.PutResponse

func (p *PutResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, fmt.Errorf("No req in PutResult")
	}
	return proto.Marshal(p.Success)
}

func (p *PutResult) Unmarshal(in []byte) error {
	msg := new(grpcConfig.PutResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PutResult) GetSuccess() *grpcConfig.PutResponse {
	if !p.IsSetSuccess() {
		return PutResult_Success_DEFAULT
	}
	return p.Success
}

func (p *PutResult) SetSuccess(x interface{}) {
	p.Success = x.(*grpcConfig.PutResponse)
}

func (p *PutResult) IsSetSuccess() bool {
	return p.Success != nil
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Get(ctx context.Context, Req *grpcConfig.Request) (r *grpcConfig.Response, err error) {
	var _args GetArgs
	_args.Req = Req
	var _result GetResult
	if err = p.c.Call(ctx, "Get", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Put(ctx context.Context, Req *grpcConfig.PutRequest) (r *grpcConfig.PutResponse, err error) {
	var _args PutArgs
	_args.Req = Req
	var _result PutResult
	if err = p.c.Call(ctx, "Put", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
