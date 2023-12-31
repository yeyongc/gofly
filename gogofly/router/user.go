package router

import (
	"gogofly/api"
	"gogofly/middleware"

	"github.com/gin-gonic/gin"
)

func InitUserRoutes() {
	RegisterRoute(func(rgPublic, rgAuth *gin.RouterGroup) {
		rgPublicUser := rgPublic.Group("user")
		{
			rgPublicUser.POST("/login", api.LoginUser)
			rgPublicUser.POST("/register", api.AddUser)
		}

		rgAuthUser := rgAuth.Group("user")
		rgAuthUser.Use(middleware.Auth())
		{
			rgAuthUser.POST("/:id", api.GetUserById)
			rgAuthUser.POST("/get", api.GetUserById)

			// rgAuthUser.POST("/list", api.GetUserList)
			rgAuthUser.POST("/:id/update", api.UpdateUser)
			rgAuthUser.POST("/update", api.UpdateUser)
			rgAuthUser.DELETE("/:id/delete", api.DeleteUserById)
			rgAuthUser.DELETE("/delete", api.DeleteUserById)

			rgAuthUserAddress := rgAuthUser.Group("/address")
			{
				rgAuthUserAddress.GET("/get", api.GetAddressLsitByUserId)
				rgAuthUserAddress.POST("/update", api.UpdateAddress)
				rgAuthUserAddress.DELETE("/delete", api.DeleteAddress)
				rgAuthUserAddress.POST("/add", api.AddAddress)
			}

			rgAuthUserTodo := rgAuthUser.Group("/todo")
			{
				rgAuthUserTodo.POST("/add", api.AddTodo)
				rgAuthUserTodo.DELETE("/delete", api.DeleteTodo)
				rgAuthUserTodo.GET("/get", api.GetTodoList)
			}
		}
	})
}
