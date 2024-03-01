package service

import (
	"context"
	"errors"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	user "github.com/wheelergeo/g-otter-gateway/biz/model/user"
	"github.com/wheelergeo/g-otter-gateway/biz/rpc"
	userservice "github.com/wheelergeo/g-otter-gen/user"
	"github.com/wheelergeo/g-otter-pkg/utils"
	"golang.org/x/sync/errgroup"
)

type UserAdminLoginService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUserAdminLoginService(Context context.Context, RequestContext *app.RequestContext) *UserAdminLoginService {
	return &UserAdminLoginService{RequestContext: RequestContext, Context: Context}
}

func (h *UserAdminLoginService) Run(req *user.UserAdminLoginReq) (resp *user.UserAdminLoginResp, err error) {
	//hlog.CtxInfof(h.Context, "baggage: %v", baggage.FromContext(h.Context).String())
	var userLoginResp *userservice.RpcWebLoginResp
	eg, ctx := errgroup.WithContext(h.Context)
	eg.Go(func() error {
		ip := h.RequestContext.ClientIP()
		agent := string(h.RequestContext.UserAgent())
		resp, err := rpc.UserClient.RpcWebLogin(ctx, &userservice.RpcWebLoginReq{
			PhoneNum: req.PhoneNum,
			Password: req.Password,
			Ip:       ip,
			Location: utils.HttpGetClientLocation(ip),
			Browser:  utils.HttpGetClientBrowser(agent),
			Os:       utils.HttpGetClientOs(agent),
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

	resp = &user.UserAdminLoginResp{
		Token: userLoginResp.Token,
		Status: &user.BaseResp{
			Code:    userLoginResp.Status.Code,
			Message: userLoginResp.Status.Message,
		},
	}
	return
}
