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

type UserCommonRoleUpdateService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUserCommonRoleUpdateService(Context context.Context, RequestContext *app.RequestContext) *UserCommonRoleUpdateService {
	return &UserCommonRoleUpdateService{RequestContext: RequestContext, Context: Context}
}

func (h *UserCommonRoleUpdateService) Run(req *user.UserCommonRoleUpdateReq) (resp *user.UserCommonRoleUpdateResp, err error) {
	//hlog.CtxInfof(h.Context, "baggage: %v", baggage.FromContext(h.Context).String())
	var userCommonRoleResp *userservice.RpcUserCommonRoleUpdateResp
	eg, ctx := errgroup.WithContext(h.Context)
	eg.Go(func() error {
		resp, err := rpc.UserClient.RpcUserCommonRoleUpdate(ctx,
			&userservice.RpcUserCommonRoleUpdateReq{
				Id:        req.Id,
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
			hlog.CtxErrorf(ctx, "call user common role update method error: %s",
				err.Error())
			return err
		}
		userCommonRoleResp = resp
		return nil
	})
	if err = eg.Wait(); err != nil {
		return
	}

	resp = &user.UserCommonRoleUpdateResp{
		Status: &user.BaseResp{
			Code:    userCommonRoleResp.Status.Code,
			Message: userCommonRoleResp.Status.Message,
		},
	}
	return
}
