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

type UserCommonDeptCreateService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUserCommonDeptCreateService(Context context.Context, RequestContext *app.RequestContext) *UserCommonDeptCreateService {
	return &UserCommonDeptCreateService{RequestContext: RequestContext, Context: Context}
}

func (h *UserCommonDeptCreateService) Run(req *user.UserCommonDeptCreateReq) (resp *user.UserCommonDeptCreateResp, err error) {
	//hlog.CtxInfof(h.Context, "baggage: %v", baggage.FromContext(h.Context).String())
	var userCommonDeptResp *userservice.RpcUserCommonDeptCreateResp
	eg, ctx := errgroup.WithContext(h.Context)
	eg.Go(func() error {
		resp, err := rpc.UserClient.RpcUserCommonDeptCreate(ctx,
			&userservice.RpcUserCommonDeptCreateReq{
				ParentId:   req.ParentId,
				AncestorId: req.AncestorId,
				DeptName:   req.DeptName,
				Leader:     req.Leader,
				PhoneNum:   req.PhoneNum,
				Email:      req.Email,
				Status:     req.Status,
			},
		)

		if resp == nil {
			err = errors.New("response empty")
		}
		if err != nil {
			hlog.CtxErrorf(ctx, "call user common dept create method error: %s", err.Error())
			return err
		}
		userCommonDeptResp = resp
		return nil
	})
	if err = eg.Wait(); err != nil {
		return
	}

	resp = &user.UserCommonDeptCreateResp{
		Status: &user.BaseResp{
			Code:    userCommonDeptResp.Status.Code,
			Message: userCommonDeptResp.Status.Message,
		},
	}
	return
}
