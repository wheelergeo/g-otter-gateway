package service

import (
	"context"
	"errors"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	user "github.com/wheelergeo/g-otter-gateway/biz/model/user"
	"github.com/wheelergeo/g-otter-gateway/biz/rpc"
	userservice "github.com/wheelergeo/g-otter-gen/user"
	"go.opentelemetry.io/otel/baggage"
	"golang.org/x/sync/errgroup"
)

type UserLoginService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUserLoginService(Context context.Context, RequestContext *app.RequestContext) *UserLoginService {
	return &UserLoginService{RequestContext: RequestContext, Context: Context}
}

func (h *UserLoginService) Run(req *user.UserLoginReq) (resp *user.UserLoginResp, err error) {
	hlog.CtxInfof(h.Context, "baggage: %v", baggage.FromContext(h.Context).String())
	var userLoginResp *userservice.LoginResp
	eg, ctx := errgroup.WithContext(h.Context)
	eg.Go(func() error {
		resp, err := rpc.UserClient.Login(ctx, &userservice.LoginReq{
			PhoneNum: req.PhoneNum,
			Password: req.Password,
		})

		if resp == nil {
			err = errors.New("response empty")
		}
		if err != nil {
			hlog.CtxErrorf(ctx, "call user login method error: %s", err.Error())
			return err
		}
		userLoginResp = resp
		return nil
	})
	if err = eg.Wait(); err != nil {
		return
	}

	resp = &user.UserLoginResp{
		Token: userLoginResp.Token,
		Status: &user.BaseResp{
			Code:    userLoginResp.Status.Code,
			Message: userLoginResp.Status.Message,
		},
	}
	return
}
