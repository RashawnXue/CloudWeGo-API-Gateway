// Code generated by Kitex v0.6.1. DO NOT EDIT.

package hertzsvr

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	service "kitexSvr-serviceB/kitex_gen/kitex/service"
)

func serviceInfo() *kitex.ServiceInfo {
	return hertzSvrServiceInfo
}

var hertzSvrServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "HertzSvr"
	handlerType := (*service.HertzSvr)(nil)
	methods := map[string]kitex.MethodInfo{
		"Mul": kitex.NewMethodInfo(mulHandler, newHertzSvrMulArgs, newHertzSvrMulResult, false),
		"Div": kitex.NewMethodInfo(divHandler, newHertzSvrDivArgs, newHertzSvrDivResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "service",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.6.1",
		Extra:           extra,
	}
	return svcInfo
}

func mulHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*service.HertzSvrMulArgs)
	realResult := result.(*service.HertzSvrMulResult)
	success, err := handler.(service.HertzSvr).Mul(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newHertzSvrMulArgs() interface{} {
	return service.NewHertzSvrMulArgs()
}

func newHertzSvrMulResult() interface{} {
	return service.NewHertzSvrMulResult()
}

func divHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*service.HertzSvrDivArgs)
	realResult := result.(*service.HertzSvrDivResult)
	success, err := handler.(service.HertzSvr).Div(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newHertzSvrDivArgs() interface{} {
	return service.NewHertzSvrDivArgs()
}

func newHertzSvrDivResult() interface{} {
	return service.NewHertzSvrDivResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Mul(ctx context.Context, request *service.Request) (r *service.Response, err error) {
	var _args service.HertzSvrMulArgs
	_args.Request = request
	var _result service.HertzSvrMulResult
	if err = p.c.Call(ctx, "Mul", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Div(ctx context.Context, request *service.Request) (r *service.Response, err error) {
	var _args service.HertzSvrDivArgs
	_args.Request = request
	var _result service.HertzSvrDivResult
	if err = p.c.Call(ctx, "Div", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}