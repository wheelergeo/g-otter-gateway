package utils

import (
	"context"
	"net/http"

	"github.com/cloudwego/dynamicgo/thrift/base"
	"github.com/cloudwego/hertz/pkg/app"
)

// SendErrResponse  pack error response
func SendErrResponse(ctx context.Context, c *app.RequestContext, code int, err error) {
	// todo edit custom code
	c.String(code, err.Error())
}

// SendInternalErrResponse  pack error response
func SendInternalErrResponse(ctx context.Context, c *app.RequestContext) {
	// todo edit custom code
	c.JSON(http.StatusInternalServerError, &base.BaseResp{
		StatusMessage: "micro user internal error",
		StatusCode:    http.StatusInternalServerError,
		Extra:         nil,
	})
}

// SendSuccessResponse  pack success response
func SendSuccessResponse(ctx context.Context, c *app.RequestContext, code int, data interface{}) {
	// todo edit custom code
	c.JSON(code, data)
}
