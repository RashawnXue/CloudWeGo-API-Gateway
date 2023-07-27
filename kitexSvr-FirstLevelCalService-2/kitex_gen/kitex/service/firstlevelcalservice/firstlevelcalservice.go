// Code generated by Kitex v0.6.1. DO NOT EDIT.

package firstlevelcalservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	service "kitexSvr-FirstLevelCalService/kitex_gen/kitex/service"
)

func serviceInfo() *kitex.ServiceInfo {
	return firstLevelCalServiceServiceInfo
}

var firstLevelCalServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "FirstLevelCalService"
	handlerType := (*service.FirstLevelCalService)(nil)
	methods := map[string]kitex.MethodInfo{
		"Add": kitex.NewMethodInfo(addHandler, newFirstLevelCalServiceAddArgs, newFirstLevelCalServiceAddResult, false),
		"Sub": kitex.NewMethodInfo(subHandler, newFirstLevelCalServiceSubArgs, newFirstLevelCalServiceSubResult, false),
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

func addHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*service.FirstLevelCalServiceAddArgs)
	realResult := result.(*service.FirstLevelCalServiceAddResult)
	success, err := handler.(service.FirstLevelCalService).Add(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newFirstLevelCalServiceAddArgs() interface{} {
	return service.NewFirstLevelCalServiceAddArgs()
}

func newFirstLevelCalServiceAddResult() interface{} {
	return service.NewFirstLevelCalServiceAddResult()
}

func subHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*service.FirstLevelCalServiceSubArgs)
	realResult := result.(*service.FirstLevelCalServiceSubResult)
	success, err := handler.(service.FirstLevelCalService).Sub(ctx, realArg.Request)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newFirstLevelCalServiceSubArgs() interface{} {
	return service.NewFirstLevelCalServiceSubArgs()
}

func newFirstLevelCalServiceSubResult() interface{} {
	return service.NewFirstLevelCalServiceSubResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Add(ctx context.Context, request *service.Request) (r *service.Response, err error) {
	var _args service.FirstLevelCalServiceAddArgs
	_args.Request = request
	var _result service.FirstLevelCalServiceAddResult
	if err = p.c.Call(ctx, "Add", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Sub(ctx context.Context, request *service.Request) (r *service.Response, err error) {
	var _args service.FirstLevelCalServiceSubArgs
	_args.Request = request
	var _result service.FirstLevelCalServiceSubResult
	if err = p.c.Call(ctx, "Sub", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
