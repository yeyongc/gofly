package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// 中间层：处理跨域
func Cors() gin.HandlerFunc {
	cfg := cors.Config{
		AllowMethods:     []string{"POST", "GET", "PATCH", "HEAD", "OPTION", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "token", "Authorization", " Accept"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
	}
	return cors.New(cfg)
}
