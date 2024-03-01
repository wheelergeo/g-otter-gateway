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

type UserCommonPostCreateService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUserCommonPostCreateService(Context context.Context, RequestContext *app.RequestContext) *UserCommonPostCreateService {
	return &UserCommonPostCreateService{RequestContext: RequestContext, Context: Context}
}

func (h *UserCommonPostCreateService) Run(req *user.UserCommonPostCreateReq) (resp *user.UserCommonPostCreateResp, err error) {
	//hlog.CtxInfof(h.Context, "baggage: %v", baggage.FromContext(h.Context).String())
	var userCommonPostResp *userservice.RpcUserCommonPostCreateResp
	eg, ctx := errgroup.WithContext(h.Context)
	eg.Go(func() error {
		resp, err := rpc.UserClient.RpcUserCommonPostCreate(ctx,
			&userservice.RpcUserCommonPostCreateReq{
				PostCode: req.PostCode,
				PostName: req.PostName,
				Level:    req.Level,
				Status:   req.Status,
				Remark:   req.Remark,
				RoleIds:  req.RoleIds,
			},
		)

		if resp == nil {
			err = errors.New("response empty")
		}
		if err != nil {
			hlog.CtxErrorf(ctx, "call user common post create method error: %s", err.Error())
			return err
		}
		userCommonPostResp = resp
		return nil
	})
	if err = eg.Wait(); err != nil {
		return
	}

	resp = &user.UserCommonPostCreateResp{
		Status: &user.BaseResp{
			Code:    userCommonPostResp.Status.Code,
			Message: userCommonPostResp.Status.Message,
		},
	}
	return
}
