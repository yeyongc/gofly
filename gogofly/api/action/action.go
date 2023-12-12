package action

import (
	"fmt"
	"gogofly/api/goflyhttp"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Action struct {
	Req         *goflyhttp.Request
	Ctx         *gin.Context
	ErrCode     int
	ServiceFunc func() (any, error)
}

type (
	Response     = goflyhttp.Response
	ResponseData = goflyhttp.ResponseData
)

func HandleActionAndGenerateResponse(action Action) Response {
	var (
		err  error
		resp Response
	)

	//解析request
	err = goflyhttp.ParseRequest(action.Ctx, action.Req)
	if err != nil {
		fmt.Println(err.Error())
		data := ResponseData{
			Code: goflyhttp.ERR_CODE_ARGUMENT_INVALID,
			Payload: map[string]interface{}{
				"error": "无效请求参数",
			},
		}
		resp = Response{
			Code: http.StatusBadRequest,
			Data: data,
		}
		return resp
	}

	// 服务请求, 并生成对应response
	serviceData, err := action.ServiceFunc()

	if err != nil {
		data := ResponseData{
			Code: action.ErrCode,
			Payload: map[string]interface{}{
				"error": err.Error(),
			},
		}
		resp = Response{
			Code: http.StatusServiceUnavailable,
			Data: data,
		}
		return resp
	}
	resp = Response{
		Code: http.StatusOK,
		Data: ResponseData{
			Payload: serviceData,
			Code:    goflyhttp.SUCCESS_CODE,
		},
	}
	return resp
}
