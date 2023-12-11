package api

import (
	"gogofly/api/action"
	"gogofly/api/goflyhttp"

	"go.uber.org/zap"
)

type (
	Response = goflyhttp.Response
	Request  = goflyhttp.Request
	Action   = action.Action
)

var (
	HandleActionAndGenerateResponse = action.HandleActionAndGenerateResponse
	SendResponse                    = goflyhttp.SendResponse
)

type BaseApi struct {
	Logger *zap.SugaredLogger // 日志
}
