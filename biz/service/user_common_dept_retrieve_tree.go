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

type UserCommonDeptRetrieveTreeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUserCommonDeptRetrieveTreeService(Context context.Context, RequestContext *app.RequestContext) *UserCommonDeptRetrieveTreeService {
	return &UserCommonDeptRetrieveTreeService{RequestContext: RequestContext, Context: Context}
}

func (h *UserCommonDeptRetrieveTreeService) Run(req *user.UserCommonDeptRetrieveTreeReq) (resp *user.UserCommonDeptRetrieveTreeResp, err error) {
	//hlog.CtxInfof(h.Context, "baggage: %v", baggage.FromContext(h.Context).String())
	var userCommonDeptResp *userservice.RpcUserCommonDeptRetrieveTreeResp
	eg, ctx := errgroup.WithContext(h.Context)
	eg.Go(func() error {
		resp, err := rpc.UserClient.RpcUserCommonDeptRetrieveTree(ctx,
			&userservice.RpcUserCommonDeptRetrieveTreeReq{
				AncestorId: req.AncestorId,
			},
		)

		if resp == nil {
			err = errors.New("response empty")
		}
		if err != nil {
			hlog.CtxErrorf(ctx, "call user common dept retrieve tree method error: %s", err.Error())
			return err
		}
		userCommonDeptResp = resp
		return nil
	})
	if err = eg.Wait(); err != nil {
		return
	}

	resp = &user.UserCommonDeptRetrieveTreeResp{
		Status: &user.BaseResp{
			Code:    userCommonDeptResp.Status.Code,
			Message: userCommonDeptResp.Status.Message,
		},
		List: resp.List,
	}
	return
}
