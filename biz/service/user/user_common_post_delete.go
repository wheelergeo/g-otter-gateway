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

type UserCommonPostDeleteService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUserCommonPostDeleteService(Context context.Context, RequestContext *app.RequestContext) *UserCommonPostDeleteService {
	return &UserCommonPostDeleteService{RequestContext: RequestContext, Context: Context}
}

func (h *UserCommonPostDeleteService) Run(req *user.UserCommonPostDeleteReq) (resp *user.UserCommonPostDeleteResp, err error) {
	//hlog.CtxInfof(h.Context, "baggage: %v", baggage.FromContext(h.Context).String())
	var userCommonPostResp *userservice.RpcUserCommonPostDeleteResp
	eg, ctx := errgroup.WithContext(h.Context)
	eg.Go(func() error {
		resp, err := rpc.UserClient.RpcUserCommonPostDelete(ctx,
			&userservice.RpcUserCommonPostDeleteReq{
				Ids: req.Ids,
			},
		)

		if resp == nil {
			err = errors.New("response empty")
		}
		if err != nil {
			hlog.CtxErrorf(ctx, "call user common post delete method error: %s", err.Error())
			return err
		}
		userCommonPostResp = resp
		return nil
	})
	if err = eg.Wait(); err != nil {
		return
	}

	resp = &user.UserCommonPostDeleteResp{
		Status: &user.BaseResp{
			Code:    userCommonPostResp.Status.Code,
			Message: userCommonPostResp.Status.Message,
		},
	}
	return
}
