package router

import (
	"github.com/gin-gonic/gin"
)

func InitHostRoutes() {
	RegisterRoute(func(rgPublic, rgAuth *gin.RouterGroup) {
		// reAuthHost := rgAuth.Group("host")
		// {
		// 	// reAuthHost.POST("/shutdown", api.ShutdownHost)
		// }

	})
}
