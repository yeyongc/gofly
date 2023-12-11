package goflyhttp

import (
	"github.com/gin-gonic/gin"
)

type Request struct {
	BindUri bool // 是否从uri中获取参数
	DTO     interface{}
}

// 根据DTO解析Request数据
func ParseRequest(ctx *gin.Context, req *Request) error {
	var err error
	if req.BindUri {
		ctx.ShouldBindUri(req.DTO)
	}
	err = ctx.ShouldBind(req.DTO)
	return err
}
