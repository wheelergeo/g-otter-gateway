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

type UserCommonDeptRetrieveService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUserCommonDeptRetrieveService(Context context.Context, RequestContext *app.RequestContext) *UserCommonDeptRetrieveService {
	return &UserCommonDeptRetrieveService{RequestContext: RequestContext, Context: Context}
}

func (h *UserCommonDeptRetrieveService) Run(req *user.UserCommonDeptRetrieveReq) (resp *user.UserCommonDeptRetrieveResp, err error) {
	//hlog.CtxInfof(h.Context, "baggage: %v", baggage.FromContext(h.Context).String())
	var userCommonDeptResp *userservice.RpcUserCommonDeptRetrieveResp
	eg, ctx := errgroup.WithContext(h.Context)
	eg.Go(func() error {
		resp, err := rpc.UserClient.RpcUserCommonDeptRetrieve(ctx,
			&userservice.RpcUserCommonDeptRetrieveReq{
				ParentId:   req.ParentId,
				AncestorId: req.AncestorId,
				DeptName:   req.DeptName,
				Status:     req.Status,
				Id:         req.Id,
			},
		)

		if resp == nil {
			err = errors.New("response empty")
		}
		if err != nil {
			hlog.CtxErrorf(ctx, "call user common dept retrieve method error: %s",
				err.Error())
			return err
		}
		userCommonDeptResp = resp
		return nil
	})
	if err = eg.Wait(); err != nil {
		return
	}

	resp = &user.UserCommonDeptRetrieveResp{
		Status: &user.BaseResp{
			Code:    userCommonDeptResp.Status.Code,
			Message: userCommonDeptResp.Status.Message,
		},
		List: resp.List,
	}
	return
}
