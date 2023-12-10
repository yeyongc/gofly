package cmd

import (
	"fmt"
	"gogofly/conf"
	"gogofly/global"
	"gogofly/router"
	"gogofly/utils"
)

func Start() {
	var initError error

	// 初始化配置文件
	conf.InitConfig()

	// 初始化日志配置
	global.Logger = conf.InitLogger()

	// 初始化数据库
	db, err := conf.InitDB()
	global.DB = db
	if err != nil {
		initError = utils.AppendError(initError, err)
	}
	if initError != nil {
		if global.Logger != nil {
			global.Logger.Error(initError.Error())
		}
		panic(initError.Error())
	}

	// 初始化redis
	rdClient, err := conf.InitRedis()
	global.RedisClient = rdClient
	if err != nil {
		initError = utils.AppendError(initError, err)
	}
	// global.RedisClient.Set("kwe", "asdfdg")
	if initError != nil {
		if global.Logger != nil {
			global.Logger.Error(initError.Error())
		}
		panic(initError.Error())
	}

	// 初始化路由器
	router.InitRouter()
}

func Clean() {
	fmt.Println("clean all things...")

}
