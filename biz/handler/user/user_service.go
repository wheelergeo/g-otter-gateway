package user

import (
	"context"
	"net/http"

	"github.com/cloudwego/dynamicgo/thrift/base"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	user "github.com/wheelergeo/g-otter-gateway/biz/model/user"
	"github.com/wheelergeo/g-otter-gateway/biz/service"
	"github.com/wheelergeo/g-otter-gateway/biz/utils"
)

// UserLogin .
// @router /login [POST]
func UserLogin(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UserLoginReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewUserLoginService(ctx, c).Run(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &base.BaseResp{
			StatusMessage: "micro user internal error",
			StatusCode:    http.StatusInternalServerError,
			Extra:         nil,
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// UserAdminLogin .
// @router /login/admin [POST]
func UserAdminLogin(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UserAdminLoginReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewUserAdminLoginService(ctx, c).Run(&req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, &base.BaseResp{
			StatusMessage: "micro user internal error",
			StatusCode:    http.StatusInternalServerError,
			Extra:         nil,
		})
		return
	}
	c.JSON(http.StatusOK, resp)
}
