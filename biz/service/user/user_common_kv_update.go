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

type UserCommonKvUpdateService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUserCommonKvUpdateService(Context context.Context, RequestContext *app.RequestContext) *UserCommonKvUpdateService {
	return &UserCommonKvUpdateService{RequestContext: RequestContext, Context: Context}
}

func (h *UserCommonKvUpdateService) Run(req *user.UserCommonKvUpdateReq) (resp *user.UserCommonKvUpdateResp, err error) {
	//hlog.CtxInfof(h.Context, "baggage: %v", baggage.FromContext(h.Context).String())
	var userCommonKvResp *userservice.RpcUserCommonKvUpdateResp
	eg, ctx := errgroup.WithContext(h.Context)
	eg.Go(func() error {
		resp, err := rpc.UserClient.RpcUserCommonKvUpdate(ctx,
			&userservice.RpcUserCommonKvUpdateReq{
				Key:   req.Key,
				Value: req.Value,
			},
		)

		if resp == nil {
			err = errors.New("response empty")
		}
		if err != nil {
			hlog.CtxErrorf(ctx, "call user common kv update method error: %s",
				err.Error())
			return err
		}
		userCommonKvResp = resp
		return nil
	})
	if err = eg.Wait(); err != nil {
		return
	}

	resp = &user.UserCommonKvUpdateResp{
		Status: &user.BaseResp{
			Code:    userCommonKvResp.Status.Code,
			Message: userCommonKvResp.Status.Message,
		},
	}
	return
}
