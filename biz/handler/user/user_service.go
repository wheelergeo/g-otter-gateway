package user

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	user "github.com/wheelergeo/g-otter-gateway/biz/model/user"
	userService "github.com/wheelergeo/g-otter-gateway/biz/service/user"
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

	resp, err := userService.NewUserAdminLoginService(ctx, c).Run(&req)

	if err != nil {
		utils.SendInternalErrResponse(ctx, c)
		return
	}
	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
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

	resp, err := userService.NewUserCommonDeptCreateService(ctx, c).Run(&req)

	if err != nil {
		utils.SendInternalErrResponse(ctx, c)
		return
	}
	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
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

	resp, err := userService.NewUserCommonDeptRetrieveService(ctx, c).Run(&req)

	if err != nil {
		utils.SendInternalErrResponse(ctx, c)
		return
	}
	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
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

	resp, err := userService.NewUserCommonDeptRetrieveTreeService(ctx, c).Run(&req)

	if err != nil {
		utils.SendInternalErrResponse(ctx, c)
		return
	}
	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
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

	resp, err := userService.NewUserCommonDeptUpdateService(ctx, c).Run(&req)

	if err != nil {
		utils.SendInternalErrResponse(ctx, c)
		return
	}
	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
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

	resp, err := userService.NewUserCommonDeptDeleteService(ctx, c).Run(&req)

	if err != nil {
		utils.SendInternalErrResponse(ctx, c)
		return
	}
	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// UserCommonRoleCreate .
// @router /user/common/role/create [POST]
func UserCommonRoleCreate(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UserCommonRoleCreateReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := userService.NewUserCommonRoleCreateService(ctx, c).Run(&req)

	if err != nil {
		utils.SendInternalErrResponse(ctx, c)
		return
	}
	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// UserCommonRoleRetrieve .
// @router /user/common/role/retrieve [GET]
func UserCommonRoleRetrieve(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UserCommonRoleRetrieveReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := userService.NewUserCommonRoleRetrieveService(ctx, c).Run(&req)

	if err != nil {
		utils.SendInternalErrResponse(ctx, c)
		return
	}
	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// UserCommonRoleUpdate .
// @router /user/common/role/update [PUT]
func UserCommonRoleUpdate(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UserCommonRoleUpdateReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := userService.NewUserCommonRoleUpdateService(ctx, c).Run(&req)

	if err != nil {
		utils.SendInternalErrResponse(ctx, c)
		return
	}
	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// UserCommonRoleDelete .
// @router /user/common/role/delete [DELETE]
func UserCommonRoleDelete(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UserCommonRoleDeleteReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := userService.NewUserCommonRoleDeleteService(ctx, c).Run(&req)

	if err != nil {
		utils.SendInternalErrResponse(ctx, c)
		return
	}
	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// UserCommonPostCreate .
// @router /user/common/post/create [POST]
func UserCommonPostCreate(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UserCommonPostCreateReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := userService.NewUserCommonPostCreateService(ctx, c).Run(&req)

	if err != nil {
		utils.SendInternalErrResponse(ctx, c)
		return
	}
	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// UserCommonPostRetrieve .
// @router /user/common/post/retrieve [GET]
func UserCommonPostRetrieve(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UserCommonPostRetrieveReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := userService.NewUserCommonPostRetrieveService(ctx, c).Run(&req)

	if err != nil {
		utils.SendInternalErrResponse(ctx, c)
		return
	}
	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// UserCommonPostUpdate .
// @router /user/common/post/update [PUT]
func UserCommonPostUpdate(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UserCommonPostUpdateReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := userService.NewUserCommonPostUpdateService(ctx, c).Run(&req)

	if err != nil {
		utils.SendInternalErrResponse(ctx, c)
		return
	}
	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// UserCommonPostDelete .
// @router /user/common/post/delete [DELETE]
func UserCommonPostDelete(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UserCommonPostDeleteReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := userService.NewUserCommonPostDeleteService(ctx, c).Run(&req)

	if err != nil {
		utils.SendInternalErrResponse(ctx, c)
		return
	}
	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// UserCommonKvCreate .
// @router /user/common/kv/create [POST]
func UserCommonKvCreate(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UserCommonKvCreateReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := userService.NewUserCommonKvCreateService(ctx, c).Run(&req)

	if err != nil {
		utils.SendInternalErrResponse(ctx, c)
		return
	}
	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// UserCommonKvRetrieve .
// @router /user/common/kv/retrieve [GET]
func UserCommonKvRetrieve(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UserCommonKvRetrieveReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := userService.NewUserCommonKvRetrieveService(ctx, c).Run(&req)

	if err != nil {
		utils.SendInternalErrResponse(ctx, c)
		return
	}
	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// UserCommonKvUpdate .
// @router /user/common/kv/update [PUT]
func UserCommonKvUpdate(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UserCommonKvUpdateReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := userService.NewUserCommonKvUpdateService(ctx, c).Run(&req)

	if err != nil {
		utils.SendInternalErrResponse(ctx, c)
		return
	}
	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// UserCommonKvDelete .
// @router /user/common/kv/delete [DELETE]
func UserCommonKvDelete(ctx context.Context, c *app.RequestContext) {
	var err error
	var req user.UserCommonKvDeleteReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := userService.NewUserCommonKvDeleteService(ctx, c).Run(&req)

	if err != nil {
		utils.SendInternalErrResponse(ctx, c)
		return
	}
	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
