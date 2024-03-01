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

type UserCommonPostUpdateService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUserCommonPostUpdateService(Context context.Context, RequestContext *app.RequestContext) *UserCommonPostUpdateService {
	return &UserCommonPostUpdateService{RequestContext: RequestContext, Context: Context}
}

func (h *UserCommonPostUpdateService) Run(req *user.UserCommonPostUpdateReq) (resp *user.UserCommonPostUpdateResp, err error) {
	//hlog.CtxInfof(h.Context, "baggage: %v", baggage.FromContext(h.Context).String())
	var userCommonPostResp *userservice.RpcUserCommonPostUpdateResp
	eg, ctx := errgroup.WithContext(h.Context)
	eg.Go(func() error {
		resp, err := rpc.UserClient.RpcUserCommonPostUpdate(ctx,
			&userservice.RpcUserCommonPostUpdateReq{
				Id:       req.Id,
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
			hlog.CtxErrorf(ctx, "call user common post update method error: %s",
				err.Error())
			return err
		}
		userCommonPostResp = resp
		return nil
	})
	if err = eg.Wait(); err != nil {
		return
	}

	resp = &user.UserCommonPostUpdateResp{
		Status: &user.BaseResp{
			Code:    userCommonPostResp.Status.Code,
			Message: userCommonPostResp.Status.Message,
		},
	}
	return
}
