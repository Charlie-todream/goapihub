package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"goapihub/app/cmd"
	make2 "goapihub/app/cmd/make"
	"goapihub/bootstrap"
	btsConfig "goapihub/config"
	"goapihub/pkg/config"
	"goapihub/pkg/console"
	"os"
)

func init()  {
	// 加载config 目录下的配置信息
	btsConfig.Initialize()
}

//func main() {
//	// 配置初始化，依赖命令行 --env 参数
//	btsConfig.Initialize()
//	var env string
//	flag.StringVar(&env,"env","","加载 .env 文件，如--env=testing 加载的是 .env.testing 文件")
//	flag.Parse()
//	config.InitConfig(env)
//
//	// new 一个Gin Engine 实例
//	router := gin.New()
//    // 初始化DB
//	bootstrap.SetupDB()
//	bootstrap.SetupRedis()
//	bootstrap.SetupLogger()
//	gin.SetMode(gin.ReleaseMode)
//	// 初始化路由绑定
//	bootstrap.SetupRoute(router)
//	// 运行服务
//	err := router.Run(":" + config.Get("app.port"))
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//}

func main() {

	// 应用的主入口，默认调用 cmd.CmdServe 命令
	var rootCmd = &cobra.Command{
		Use:   config.Get("app.name"),
		Short: "A simple forum project",
		Long:  `Default will run "serve" command, you can use "-h" flag to see all subcommands`,

		// rootCmd 的所有子命令都会执行以下代码
		PersistentPreRun: func(command *cobra.Command, args []string) {

			// 配置初始化，依赖命令行 --env 参数
			config.InitConfig(cmd.Env)

			// 初始化 Logger
			bootstrap.SetupLogger()

			// 初始化数据库
			bootstrap.SetupDB()

			// 初始化 Redis
			bootstrap.SetupRedis()

			// 初始化缓存
			// 初始化缓存
			bootstrap.SetupCache()
		},
	}

	// 注册子命令
	rootCmd.AddCommand(
		cmd.CmdServe,
		cmd.CmdKey,
		cmd.CmdPlay,
		make2.CmdMake,
		cmd.CmdMigrate,
		cmd.CmdDBSeed,
	)

	// 配置默认运行 Web 服务
	cmd.RegisterDefaultCmd(rootCmd, cmd.CmdServe)

	// 注册全局参数，--env
	cmd.RegisterGlobalFlags(rootCmd)

	// 执行主命令
	if err := rootCmd.Execute(); err != nil {
		console.Exit(fmt.Sprintf("Failed to run app with %v: %s", os.Args, err.Error()))
	}
}
