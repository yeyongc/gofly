package api

import (
	"gogofly/global"
	"gogofly/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// 基础qpi结构
type BaseApi struct {
	Ctx    *gin.Context       // 用于获取请求参数
	Errors error              // 处理错误，如请求参数错误等
	Logger *zap.SugaredLogger // 用于生成日志
}

// 请求配置结构
type BuildRequestOption struct {
	DTO               interface{} // gin.Context中的请求参数映射到DTO中
	BindParamsFromUri bool        // 参数是否来自uri，如动态路由的params
}

func (api *BaseApi) AddError(err error) {
	api.Errors = utils.AppendError(api.Errors, err)
}

func (api *BaseApi) GetErrors() error {
	return api.Errors
}

// 构建基本请求
func (api *BaseApi) BuildRequest(option *BuildRequestOption) *BaseApi {
	var errResult error

	// 绑定请求数据
	if option.BindParamsFromUri {
		api.Ctx.ShouldBindUri(option.DTO)
	}

	// 映射解构：将ctx内与DTO相对应的值存储到DTO实例中
	errResult = api.Ctx.ShouldBind(option.DTO)

	// 处理绑定数据出错
	if errResult != nil {
		// 解析错误
		errResult = utils.ParseValidateErrors(errResult.(validator.ValidationErrors), option.DTO)
		// 向追加错误
		api.AddError(errResult)

		// 发送错误响应
		api.Fail(ResponseJson{
			Msg: api.GetErrors().Error(),
		})
	}
	return api
}

func (api *BaseApi) Ok(resp ResponseJson) {
	Ok(api.Ctx, resp)
}

func (api *BaseApi) Fail(resp ResponseJson) {
	Fail(api.Ctx, resp)
}
func (api *BaseApi) ServerFail(resp ResponseJson) {
	ServerFail(api.Ctx, resp)
}

func NewBaseApi(ctx *gin.Context) BaseApi {
	return BaseApi{
		Logger: global.Logger,
		Ctx:    ctx,
	}
}
