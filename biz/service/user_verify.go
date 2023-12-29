package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	admin "github.com/wheelergeo/g-otter-gateway/biz/model/user/admin"
	"github.com/wheelergeo/g-otter-gateway/biz/rpc"
	"github.com/wheelergeo/g-otter-gen/user"
)

type UserVerifyService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUserVerifyService(Context context.Context, RequestContext *app.RequestContext) *UserVerifyService {
	return &UserVerifyService{RequestContext: RequestContext, Context: Context}
}

func (h *UserVerifyService) Run(req *admin.UserVerifyReq) (resp *admin.UserVerifyResp, err error) {
	defer func() {
		hlog.CtxInfof(h.Context, "req = %+v", req)
		hlog.CtxInfof(h.Context, "resp = %+v", resp)
	}()

	_, err = rpc.UserClient.Verify(h.Context, &user.VerifyReq{
		PhoneNum: req.PhoneNum,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	resp = &admin.UserVerifyResp{
		Token: "11111111",
		Status: &admin.BaseResp{
			Message: "Succcss",
			Code:    200,
		},
	}
	return
}
