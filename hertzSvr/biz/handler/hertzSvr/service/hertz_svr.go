// Code generated by hertz generator.

package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"hertzSvr/biz/handler/hertzSvr/utils"
	"hertzSvr/biz/model/hertzSvr/service"
)

// Request .
// @router /request [POST]
func Request(ctx context.Context, c *app.RequestContext) {
	var err error
	var req service.SvrRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	// 从body中获取服务名（后续考虑通过路由层进行优化）
	svcName := req.SvrName
	cli := clients[svcName]

	resp := utils.GetHTTPGenericResponse(ctx, c, "", cli)

	c.JSON(consts.StatusOK, resp)
}

// RegisterIDL .
// @router /registerIDL [POST]
func RegisterIDL(ctx context.Context, c *app.RequestContext) {
	var err error
	var req service.RegisterIDL
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(service.SvrResponse)

	c.JSON(consts.StatusOK, resp)
}