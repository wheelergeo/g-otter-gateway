package admin

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	admin "github.com/wheelergeo/g-otter-gateway/biz/model/user/admin"
	"github.com/wheelergeo/g-otter-gateway/biz/service"
	"github.com/wheelergeo/g-otter-gateway/biz/utils"
)

// UserVerify .
// @router /user/verify [GET]
func UserVerify(ctx context.Context, c *app.RequestContext) {
	var err error
	var req admin.UserVerifyReq
	err = c.BindAndValidate(&req)
	hlog.Info(req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewUserVerifyService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
