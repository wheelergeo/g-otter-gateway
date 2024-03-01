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

type UserCommonRoleCreateService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUserCommonRoleCreateService(Context context.Context, RequestContext *app.RequestContext) *UserCommonRoleCreateService {
	return &UserCommonRoleCreateService{RequestContext: RequestContext, Context: Context}
}

func (h *UserCommonRoleCreateService) Run(req *user.UserCommonRoleCreateReq) (resp *user.UserCommonRoleCreateResp, err error) {
	//hlog.CtxInfof(h.Context, "baggage: %v", baggage.FromContext(h.Context).String())
	var userCommonRoleResp *userservice.RpcUserCommonRoleCreateResp
	eg, ctx := errgroup.WithContext(h.Context)
	eg.Go(func() error {
		resp, err := rpc.UserClient.RpcUserCommonRoleCreate(ctx,
			&userservice.RpcUserCommonRoleCreateReq{
				Status:    req.Status,
				Level:     req.Level,
				Name:      req.Name,
				Remark:    req.Remark,
				DataScope: req.DataScope,
			},
		)

		if resp == nil {
			err = errors.New("response empty")
		}
		if err != nil {
			hlog.CtxErrorf(ctx, "call user common role create method error: %s", err.Error())
			return err
		}
		userCommonRoleResp = resp
		return nil
	})
	if err = eg.Wait(); err != nil {
		return
	}

	resp = &user.UserCommonRoleCreateResp{
		Status: &user.BaseResp{
			Code:    userCommonRoleResp.Status.Code,
			Message: userCommonRoleResp.Status.Message,
		},
	}
	return
}
