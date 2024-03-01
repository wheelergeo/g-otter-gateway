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

type UserCommonKvCreateService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUserCommonKvCreateService(Context context.Context, RequestContext *app.RequestContext) *UserCommonKvCreateService {
	return &UserCommonKvCreateService{RequestContext: RequestContext, Context: Context}
}

func (h *UserCommonKvCreateService) Run(req *user.UserCommonKvCreateReq) (resp *user.UserCommonKvCreateResp, err error) {
	//hlog.CtxInfof(h.Context, "baggage: %v", baggage.FromContext(h.Context).String())
	var userCommonKvResp *userservice.RpcUserCommonKvCreateResp
	eg, ctx := errgroup.WithContext(h.Context)
	eg.Go(func() error {
		resp, err := rpc.UserClient.RpcUserCommonKvCreate(ctx,
			&userservice.RpcUserCommonKvCreateReq{
				Key:   req.Key,
				Value: req.Value,
			},
		)

		if resp == nil {
			err = errors.New("response empty")
		}
		if err != nil {
			hlog.CtxErrorf(ctx, "call user common kv create method error: %s", err.Error())
			return err
		}
		userCommonKvResp = resp
		return nil
	})
	if err = eg.Wait(); err != nil {
		return
	}

	resp = &user.UserCommonKvCreateResp{
		Status: &user.BaseResp{
			Code:    userCommonKvResp.Status.Code,
			Message: userCommonKvResp.Status.Message,
		},
	}
	return
}
