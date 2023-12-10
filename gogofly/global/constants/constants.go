package constants

import "time"

const (
	// redis keys' prefix
	REDIS_USER_PREFIX = "user:"

	// 用户登录信息
	LOGIN_USER = "LOGIN_USER"

	// redis存储key单位时间-- 1分钟
	REDIS_DURATION_UNIT = time.Second * 60
)

// 自定义错误响应状态码
const (
	// 用户服务出错状态码
	ERR_CODE_SERVICE_USER_LOGIN    = 10501 // 登录出错
	ERR_CODE_SERVICE_USER_ADD      = 10502 // 注册或添加出错
	ERR_CODE_SERVICE_USER_UPDATE   = 10503 // 更新出错
	ERR_CODE_SERVICE_USER_DELETE   = 10504 // 用户删除出错
	ERR_CODE_SERVICE_USER_GET      = 10505 // 用户获取出错
	ERR_CODE_SERVICE_USER_LIST_GET = 10506 // 用户列表获取出错

	// token相关出错
	ERR_CODE_TOKEN_INVALID     = 10201 // 无效token
	ERR_CODE_TOKEN_RENEW       = 10202 // 重新生成token出错
	ERR_CODE_TOKEN_NOT_MATCHED = 10203 // token不匹配
)
