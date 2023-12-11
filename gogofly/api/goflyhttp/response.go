package goflyhttp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code    int         `json:"code" form:"code"`
	Payload interface{} `json:"payload" form:"payload"`
}

type Response struct {
	Code int          // 响应状态码
	Data ResponseData // 响应数据
}

// 请求参数相关错误: 10XXX
const (
	// 自定义错误响应状态码----参数错误
	ERR_CODE_ARGUMENT_INVALID = 10601
)

// user service相关错误码: 105XX
const (
	// 自定义错误响应状态码----登录出错
	ERR_CODE_SERVICE_USER_LOGIN = 10501

	// 自定义错误响应状态码----注册或添加出错
	ERR_CODE_SERVICE_USER_ADD = 10502

	// 自定义错误响应状态码----用户更新出错
	ERR_CODE_SERVICE_USER_UPDATE = 10503

	// 自定义错误响应状态码----用户删除出错
	ERR_CODE_SERVICE_USER_DELETE = 10504

	// 自定义错误响应状态码----用户获取出错
	ERR_CODE_SERVICE_USER_GET = 10505

	// 自定义错误响应状态码----用户列表获取出错
	ERR_CODE_SERVICE_USER_LIST_GET = 10506
)

// address service 相关错误码: 104XX
const (
	// 自定义错误响应状态码----地址添加错误
	ERR_CODE_SERVICE_ADDRESS_ADD = 10401

	// 自定义错误响应状态码----地址更新错误
	ERR_CODE_SERVICE_ADDRESS_UPDATE = 10402

	// 自定义错误响应状态码----地址列表获取错误
	ERR_CODE_SERVICE_ADDRESS_GETALL = 10403

	// 自定义错误响应状态码----地址删除错误
	ERR_CODE_SERVICE_ADDRESS_DELETE = 10404
)

// token service 相关错误码：103XX
const (
	// 自定义错误响应状态码----无效token
	ERR_CODE_TOKEN_INVALID = 10201

	// 自定义错误响应状态码----重新生成token出错
	ERR_CODE_TOKEN_RENEW = 10202

	// 自定义错误响应状态码----token不匹配
	ERR_CODE_TOKEN_NOT_MATCHED = 10203
)

// service 成功: 101XX
const (
	SUCCESS_CODE = 10100
)

// 发送响应
func SendResponse(ctx *gin.Context, res Response) {
	ctx.AbortWithStatusJSON(res.Code, res.Data)
}

func OK(ctx *gin.Context, data ResponseData) {
	SendResponse(ctx, Response{
		Code: http.StatusOK,
		Data: data,
	})
}
func Fail(ctx *gin.Context, data ResponseData) {
	SendResponse(ctx, Response{
		Code: http.StatusBadRequest,
		Data: data,
	})
}
func ServiceFail(ctx *gin.Context, data ResponseData) {
	SendResponse(ctx, Response{
		Code: http.StatusServiceUnavailable,
		Data: data,
	})
}
