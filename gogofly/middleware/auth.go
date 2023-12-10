package middleware

import (
	"gogofly/api"
	"gogofly/global/constants"
	"gogofly/service"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	// header key
	TOKEN_NAME = "Authorization"

	// token 前缀
	TOEKN_PREFIX = "Bearer: "

	// token即将过期时间（5 mins）
	TOKEN_WILL_EXPIRE = 5 * 60 * time.Second
)

// 响应未授权
func tokenErr(ctx *gin.Context, code int) {
	api.Fail(ctx, api.ResponseJson{
		Status: http.StatusUnauthorized,
		Code:   code,
		Msg:    "Invalid token",
	})
}

// auth：生成鉴权函数
func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取请求头中的token
		token := ctx.GetHeader(TOKEN_NAME)

		// token不存在，直接执行未授权响应并返回
		if token == "" || !strings.HasPrefix(token, TOEKN_PREFIX) {
			tokenErr(ctx, constants.ERR_CODE_TOKEN_NOT_MATCHED)
			return
		}

		// 处理request token的前缀
		token = strings.Replace(token, TOEKN_PREFIX, "", 1)

		// 验证token --- 通过redis中token对比
		user, err := service.GetUserByTokenFromRedis(token)
		if err != nil {
			// redis中不存在token，说明未授权或token过期，直接返回
			tokenErr(ctx, constants.ERR_CODE_TOKEN_INVALID)
			return
		} else {

			// token过期，直接返回
			duration, err := service.GetCachedTokenDuration(token)
			if err != nil || duration.Seconds() <= 0 {
				tokenErr(ctx, constants.ERR_CODE_TOKEN_INVALID)
				return
			}

			// token即将过期，进行token续期
			if duration.Seconds() < TOKEN_WILL_EXPIRE.Seconds() {

				newToken, err := service.GenerateAndCacheTokenToRedis(user.ID, user.Name)
				if err != nil {
					tokenErr(ctx, constants.ERR_CODE_TOKEN_RENEW)
					return
				}

				// 将新的token返回客户端
				ctx.Header("token", newToken)
			}
		}
		//设置用户信息提供给后续操作
		ctx.Set(constants.LOGIN_USER, user)

		// 执行后续操作
		ctx.Next()
	}
}
