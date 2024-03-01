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

type UserCommonPostRetrieveService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUserCommonPostRetrieveService(Context context.Context, RequestContext *app.RequestContext) *UserCommonPostRetrieveService {
	return &UserCommonPostRetrieveService{RequestContext: RequestContext, Context: Context}
}

func (h *UserCommonPostRetrieveService) Run(req *user.UserCommonPostRetrieveReq) (resp *user.UserCommonPostRetrieveResp, err error) {
	//hlog.CtxInfof(h.Context, "baggage: %v", baggage.FromContext(h.Context).String())
	var userCommonPostResp *userservice.RpcUserCommonPostRetrieveResp
	eg, ctx := errgroup.WithContext(h.Context)
	eg.Go(func() error {
		resp, err := rpc.UserClient.RpcUserCommonPostRetrieve(ctx,
			&userservice.RpcUserCommonPostRetrieveReq{
				PostCode: req.PostCode,
				PostName: req.PostName,
				Level:    req.Level,
				Status:   req.Status,
				Id:       req.Id,
			},
		)

		if resp == nil {
			err = errors.New("response empty")
		}
		if err != nil {
			hlog.CtxErrorf(ctx, "call user common post retrieve method error: %s",
				err.Error())
			return err
		}
		userCommonPostResp = resp
		return nil
	})
	if err = eg.Wait(); err != nil {
		return
	}

	resp = &user.UserCommonPostRetrieveResp{
		Status: &user.BaseResp{
			Code:    userCommonPostResp.Status.Code,
			Message: userCommonPostResp.Status.Message,
		},
		List: resp.List,
	}
	return
}
