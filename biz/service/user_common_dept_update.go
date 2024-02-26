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

type UserCommonDeptUpdateService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUserCommonDeptUpdateService(Context context.Context, RequestContext *app.RequestContext) *UserCommonDeptUpdateService {
	return &UserCommonDeptUpdateService{RequestContext: RequestContext, Context: Context}
}

func (h *UserCommonDeptUpdateService) Run(req *user.UserCommonDeptUpdateReq) (resp *user.UserCommonDeptUpdateResp, err error) {
	//hlog.CtxInfof(h.Context, "baggage: %v", baggage.FromContext(h.Context).String())
	var userCommonDeptResp *userservice.RpcUserCommonDeptUpdateResp
	eg, ctx := errgroup.WithContext(h.Context)
	eg.Go(func() error {
		resp, err := rpc.UserClient.RpcUserCommonDeptUpdate(ctx,
			&userservice.RpcUserCommonDeptUpdateReq{
				Id:         req.Id,
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
			hlog.CtxErrorf(ctx, "call user common dept update method error: %s",
				err.Error())
			return err
		}
		userCommonDeptResp = resp
		return nil
	})
	if err = eg.Wait(); err != nil {
		return
	}

	resp = &user.UserCommonDeptUpdateResp{
		Status: &user.BaseResp{
			Code:    userCommonDeptResp.Status.Code,
			Message: userCommonDeptResp.Status.Message,
		},
	}
	return
}
