package service

import (
	"context"
	"errors"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	user "github.com/wheelergeo/g-otter-gateway/biz/model/user"
	"github.com/wheelergeo/g-otter-gateway/biz/rpc"
	userservice "github.com/wheelergeo/g-otter-gen/user"
	"golang.org/x/sync/errgroup"
)

type UserCommonRoleDeleteService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUserCommonRoleDeleteService(Context context.Context, RequestContext *app.RequestContext) *UserCommonRoleDeleteService {
	return &UserCommonRoleDeleteService{RequestContext: RequestContext, Context: Context}
}

func (h *UserCommonRoleDeleteService) Run(req *user.UserCommonRoleDeleteReq) (resp *user.UserCommonRoleDeleteResp, err error) {
	//hlog.CtxInfof(h.Context, "baggage: %v", baggage.FromContext(h.Context).String())
	var userCommonRoleResp *userservice.RpcUserCommonRoleDeleteResp
	eg, ctx := errgroup.WithContext(h.Context)
	eg.Go(func() error {
		resp, err := rpc.UserClient.RpcUserCommonRoleDelete(ctx,
			&userservice.RpcUserCommonRoleDeleteReq{
				Ids: req.Ids,
			},
		)

		if resp == nil {
			err = errors.New("response empty")
		}
		if err != nil {
			hlog.CtxErrorf(ctx, "call user common role delete method error: %s", err.Error())
			return err
		}
		userCommonRoleResp = resp
		return nil
	})
	if err = eg.Wait(); err != nil {
		return
	}

	resp = &user.UserCommonRoleDeleteResp{
		Status: &user.BaseResp{
			Code:    userCommonRoleResp.Status.Code,
			Message: userCommonRoleResp.Status.Message,
		},
	}
	return
}
