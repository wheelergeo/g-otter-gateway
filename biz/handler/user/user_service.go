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

// UserCommonDeptCreate .
// @router /user/common/dept/create [POST]
func UserCommonDeptCreate(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UserCommonDeptCreateReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewUserCommonDeptCreateService(ctx, c).Run(&req)

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

// UserCommonDeptRetrieve .
// @router /user/common/dept/retrieve [GET]
func UserCommonDeptRetrieve(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UserCommonDeptRetrieveReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewUserCommonDeptRetrieveService(ctx, c).Run(&req)

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

// UserCommonDeptRetrieveTree .
// @router /user/common/dept/retrieveTree [GET]
func UserCommonDeptRetrieveTree(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UserCommonDeptRetrieveTreeReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewUserCommonDeptRetrieveTreeService(ctx, c).Run(&req)

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

// UserCommonDeptUpdate .
// @router /user/common/dept/update [PUT]
func UserCommonDeptUpdate(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UserCommonDeptUpdateReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewUserCommonDeptUpdateService(ctx, c).Run(&req)

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

// UserCommonDeptDelete .
// @router /user/common/dept/delete [DELETE]
func UserCommonDeptDelete(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UserCommonDeptDeleteReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewUserCommonDeptDeleteService(ctx, c).Run(&req)

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
