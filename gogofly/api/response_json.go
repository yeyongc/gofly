package api

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

// 响应数据类型
type ResponseJson struct {
	Status int         `json:"-"`
	Code   int         `json:"code,omitempty"`
	Msg    string      `json:"msg,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

func (resp *ResponseJson) IsEmpty() bool {
	return reflect.DeepEqual(resp, ResponseJson{})
}

func HttpResponse(ctx *gin.Context, status int, resp ResponseJson) {
	if resp.IsEmpty() {
		ctx.AbortWithStatus(status)
	} else {
		ctx.AbortWithStatusJSON(status, resp)
	}

}

func Ok(ctx *gin.Context, resp ResponseJson) {
	HttpResponse(ctx, http.StatusOK, resp)
}

func Fail(ctx *gin.Context, resp ResponseJson) {
	HttpResponse(ctx, http.StatusBadRequest, resp)
}

func ServerFail(ctx *gin.Context, resp ResponseJson) {
	HttpResponse(ctx, http.StatusInternalServerError, resp)
}
