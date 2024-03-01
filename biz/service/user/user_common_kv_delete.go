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

type UserCommonKvDeleteService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUserCommonKvDeleteService(Context context.Context, RequestContext *app.RequestContext) *UserCommonKvDeleteService {
	return &UserCommonKvDeleteService{RequestContext: RequestContext, Context: Context}
}

func (h *UserCommonKvDeleteService) Run(req *user.UserCommonKvDeleteReq) (resp *user.UserCommonKvDeleteResp, err error) {
	//hlog.CtxInfof(h.Context, "baggage: %v", baggage.FromContext(h.Context).String())
	var userCommonKvResp *userservice.RpcUserCommonKvDeleteResp
	eg, ctx := errgroup.WithContext(h.Context)
	eg.Go(func() error {
		resp, err := rpc.UserClient.RpcUserCommonKvDelete(ctx,
			&userservice.RpcUserCommonKvDeleteReq{
				Keys: req.Keys,
			},
		)

		if resp == nil {
			err = errors.New("response empty")
		}
		if err != nil {
			hlog.CtxErrorf(ctx, "call user common kv delete method error: %s", err.Error())
			return err
		}
		userCommonKvResp = resp
		return nil
	})
	if err = eg.Wait(); err != nil {
		return
	}

	resp = &user.UserCommonKvDeleteResp{
		Status: &user.BaseResp{
			Code:    userCommonKvResp.Status.Code,
			Message: userCommonKvResp.Status.Message,
		},
	}
	return
}
