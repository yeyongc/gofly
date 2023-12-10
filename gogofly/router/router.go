package router

import (
	"context"
	"fmt"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "gogofly/docs"
	"gogofly/global"
	"gogofly/middleware"
)

type IFnRegisterRouter = func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup)

var (
	gfnRoutes []IFnRegisterRouter
)

// 注册路由
func RegisterRoute(fn IFnRegisterRouter) {
	if fn == nil {
		return
	}
	gfnRoutes = append(gfnRoutes, fn)
}

// 初始化路由器
func InitRouter() {
	// 处理程序关闭相关上下文
	ctx, cancelCtx := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancelCtx()

	// 默认路由
	r := gin.Default()

	// 注册middleware
	r.Use(middleware.Cors())

	// 处理路由group
	rgPublic := r.Group("/api/v1/public")
	rgAuth := r.Group("/api/v1")

	initBasePlatformRoutes()

	for _, registeredFn := range gfnRoutes {
		registeredFn(rgPublic, rgAuth)
	}

	// 集成swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 读取配置文件中的端口
	stPort := viper.GetString("server.Port")
	if stPort == "" {
		stPort = "8999"
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", stPort),
		Handler: r,
	}

	// 开启一个新的goroutine运行server
	go func() {
		global.Logger.Info("Start listen to port:", stPort)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.Logger.Error(fmt.Sprintf("Start server error: %s", err.Error()))
			return
		}
	}()

	<-ctx.Done()

	ctx, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutdown()

	if err := srv.Shutdown(ctx); err != nil {
		global.Logger.Error(fmt.Sprintf("Stop server error: %s", err.Error()))
		return
	}

	global.Logger.Info("Stop server success")
}

// 初始化路由
func initBasePlatformRoutes() {
	InitUserRoutes()
	InitHostRoutes()
}
