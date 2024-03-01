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

type UserCommonRoleRetrieveService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUserCommonRoleRetrieveService(Context context.Context, RequestContext *app.RequestContext) *UserCommonRoleRetrieveService {
	return &UserCommonRoleRetrieveService{RequestContext: RequestContext, Context: Context}
}

func (h *UserCommonRoleRetrieveService) Run(req *user.UserCommonRoleRetrieveReq) (resp *user.UserCommonRoleRetrieveResp, err error) {
	//hlog.CtxInfof(h.Context, "baggage: %v", baggage.FromContext(h.Context).String())
	var userCommonRoleResp *userservice.RpcUserCommonRoleRetrieveResp
	eg, ctx := errgroup.WithContext(h.Context)
	eg.Go(func() error {
		resp, err := rpc.UserClient.RpcUserCommonRoleRetrieve(ctx,
			&userservice.RpcUserCommonRoleRetrieveReq{
				Status: req.Status,
				Level:  req.Level,
				Name:   req.Name,
				Id:     req.Id,
			},
		)

		if resp == nil {
			err = errors.New("response empty")
		}
		if err != nil {
			hlog.CtxErrorf(ctx, "call user common role retrieve method error: %s",
				err.Error())
			return err
		}
		userCommonRoleResp = resp
		return nil
	})
	if err = eg.Wait(); err != nil {
		return
	}

	resp = &user.UserCommonRoleRetrieveResp{
		Status: &user.BaseResp{
			Code:    userCommonRoleResp.Status.Code,
			Message: userCommonRoleResp.Status.Message,
		},
		List: resp.List,
	}
	return
}
