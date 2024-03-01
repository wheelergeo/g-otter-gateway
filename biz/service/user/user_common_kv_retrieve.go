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

type UserCommonKvRetrieveService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUserCommonKvRetrieveService(Context context.Context, RequestContext *app.RequestContext) *UserCommonKvRetrieveService {
	return &UserCommonKvRetrieveService{RequestContext: RequestContext, Context: Context}
}

func (h *UserCommonKvRetrieveService) Run(req *user.UserCommonKvRetrieveReq) (resp *user.UserCommonKvRetrieveResp, err error) {
	//hlog.CtxInfof(h.Context, "baggage: %v", baggage.FromContext(h.Context).String())
	var userCommonKvResp *userservice.RpcUserCommonKvRetrieveResp
	eg, ctx := errgroup.WithContext(h.Context)
	eg.Go(func() error {
		resp, err := rpc.UserClient.RpcUserCommonKvRetrieve(ctx,
			&userservice.RpcUserCommonKvRetrieveReq{
				Key: req.Key,
			},
		)

		if resp == nil {
			err = errors.New("response empty")
		}
		if err != nil {
			hlog.CtxErrorf(ctx, "call user common kv retrieve method error: %s",
				err.Error())
			return err
		}
		userCommonKvResp = resp
		return nil
	})
	if err = eg.Wait(); err != nil {
		return
	}

	resp = &user.UserCommonKvRetrieveResp{
		Status: &user.BaseResp{
			Code:    userCommonKvResp.Status.Code,
			Message: userCommonKvResp.Status.Message,
		},
		Value: resp.Value,
	}
	return
}
