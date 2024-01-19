package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	user "github.com/wheelergeo/g-otter-gateway/biz/model/user"
)

type UserLoginService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUserLoginService(Context context.Context, RequestContext *app.RequestContext) *UserLoginService {
	return &UserLoginService{RequestContext: RequestContext, Context: Context}
}

func (h *UserLoginService) Run(req *user.UserLoginReq) (resp *user.UserLoginResp, err error) {
	//hlog.CtxInfof(h.Context, "baggage: %v", baggage.FromContext(h.Context).String())
	return
}
