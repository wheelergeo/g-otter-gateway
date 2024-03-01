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

type UserCommonDeptDeleteService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUserCommonDeptDeleteService(Context context.Context, RequestContext *app.RequestContext) *UserCommonDeptDeleteService {
	return &UserCommonDeptDeleteService{RequestContext: RequestContext, Context: Context}
}

func (h *UserCommonDeptDeleteService) Run(req *user.UserCommonDeptDeleteReq) (resp *user.UserCommonDeptDeleteResp, err error) {
	//hlog.CtxInfof(h.Context, "baggage: %v", baggage.FromContext(h.Context).String())
	var userCommonDeptResp *userservice.RpcUserCommonDeptDeleteResp
	eg, ctx := errgroup.WithContext(h.Context)
	eg.Go(func() error {
		resp, err := rpc.UserClient.RpcUserCommonDeptDelete(ctx,
			&userservice.RpcUserCommonDeptDeleteReq{
				Ids: req.Ids,
			},
		)

		if resp == nil {
			err = errors.New("response empty")
		}
		if err != nil {
			hlog.CtxErrorf(ctx, "call user common dept delete method error: %s", err.Error())
			return err
		}
		userCommonDeptResp = resp
		return nil
	})
	if err = eg.Wait(); err != nil {
		return
	}

	resp = &user.UserCommonDeptDeleteResp{
		Status: &user.BaseResp{
			Code:    userCommonDeptResp.Status.Code,
			Message: userCommonDeptResp.Status.Message,
		},
	}
	return
}
