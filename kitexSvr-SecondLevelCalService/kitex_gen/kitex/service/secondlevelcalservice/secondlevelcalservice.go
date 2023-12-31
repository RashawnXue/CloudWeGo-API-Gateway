// Code generated by Kitex v0.6.1. DO NOT EDIT.

package secondlevelcalservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	service "kitexSvr-SecondLevelCalService/kitex_gen/kitex/service"
)

func serviceInfo() *kitex.ServiceInfo {
	return secondLevelCalServiceServiceInfo
}

var secondLevelCalServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "SecondLevelCalService"
	handlerType := (*service.SecondLevelCalService)(nil)
	methods := map[string]kitex.MethodInfo{
		"Mul": kitex.NewMethodInfo(mulHandler, newSecondLevelCalServiceMulArgs, newSecondLevelCalServiceMulResult, false),
		"Div": kitex.NewMethodInfo(divHandler, newSecondLevelCalServiceDivArgs, newSecondLevelCalServiceDivResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "idlManager",
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
	realArg := arg.(*service.SecondLevelCalServiceMulArgs)
	realResult := result.(*service.SecondLevelCalServiceMulResult)
	success, err := handler.(service.SecondLevelCalService).Mul(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSecondLevelCalServiceMulArgs() interface{} {
	return service.NewSecondLevelCalServiceMulArgs()
}

func newSecondLevelCalServiceMulResult() interface{} {
	return service.NewSecondLevelCalServiceMulResult()
}

func divHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*service.SecondLevelCalServiceDivArgs)
	realResult := result.(*service.SecondLevelCalServiceDivResult)
	success, err := handler.(service.SecondLevelCalService).Div(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSecondLevelCalServiceDivArgs() interface{} {
	return service.NewSecondLevelCalServiceDivArgs()
}

func newSecondLevelCalServiceDivResult() interface{} {
	return service.NewSecondLevelCalServiceDivResult()
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
	var _args service.SecondLevelCalServiceMulArgs
	_args.Request = request
	var _result service.SecondLevelCalServiceMulResult
	if err = p.c.Call(ctx, "Mul", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Div(ctx context.Context, request *service.Request) (r *service.Response, err error) {
	var _args service.SecondLevelCalServiceDivArgs
	_args.Request = request
	var _result service.SecondLevelCalServiceDivResult
	if err = p.c.Call(ctx, "Div", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
