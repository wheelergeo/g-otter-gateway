package middleware

import (
	"context"
	"net/http"

	"github.com/cloudwego/dynamicgo/thrift/base"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/wheelergeo/g-otter-pkg/token"
)

func TokenMiddleware() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		authorization := string(ctx.GetHeader("Authorization"))
		if authorization == "" {
			ctx.JSON(http.StatusNonAuthoritativeInfo, &base.BaseResp{
				StatusMessage: "Header Authorization is null",
				StatusCode:    http.StatusNonAuthoritativeInfo,
				Extra:         nil,
			})
			ctx.Abort()
			return
		}

		claim, err := token.TokenAuth().Parse(authorization)
		if err != nil {
			ctx.JSON(http.StatusNonAuthoritativeInfo, &base.BaseResp{
				StatusMessage: "Parse Authorization failed",
				StatusCode:    http.StatusNonAuthoritativeInfo,
				Extra:         nil,
			})
			ctx.Abort()
			return
		}

		userCtx := token.ParseClaimAsUserContext(claim)
		token.SetContext(&c, &token.TokenContext{
			UserCtx: &userCtx,
		})

		ctx.Next(c)
		return
	}
}

func AuthMiddleware() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		ctx.Next(c)
	}
}
