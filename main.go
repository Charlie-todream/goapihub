package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"goapihub/bootstrap"
	btsConfig "goapihub/config"
	"goapihub/pkg/config"
)

func init()  {
	// 加载config 目录下的配置信息
	btsConfig.Initialize()
}

func main() {
	// 配置初始化，依赖命令行 --env 参数
	btsConfig.Initialize()
	var env string
	flag.StringVar(&env,"env","","加载 .env 文件，如--env=testing 加载的是 .env.testing 文件")
	flag.Parse()
	config.InitConfig(env)

	// new 一个Gin Engine 实例
	router := gin.New()

	// 初始化路由绑定
	bootstrap.SetupRoute(router)

	// 运行服务
	err := router.Run(":" + config.Get("app.port"))

	if err != nil {
		fmt.Println(err.Error())
	}
}
