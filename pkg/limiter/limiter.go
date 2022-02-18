package limiter

import (
	"github.com/gin-gonic/gin"
	limiterlib "github.com/ulule/limiter/v3"
	sredis "github.com/ulule/limiter/v3/drivers/store/redis"
	"goapihub/pkg/config"
	"goapihub/pkg/logger"
	"goapihub/pkg/redis"
	"strings"
)

// 获取IP
func GetKeyIP(c *gin.Context)  string {
	return c.ClientIP()
}
// 检测CheckRate 请求是否超额
func CheckRate(c *gin.Context,key string,formatted string) (limiterlib.Context,error){
	 var context limiterlib.Context
	 rate,err := limiterlib.NewRateFromFormatted(formatted);
	 if err != nil {
	 	logger.LogIf(err)
	 	return  context,err
	 }
	// 初始化存储，使用我们程序里共用的 redis.Redis 对象
	store,err := sredis.NewStoreWithOptions(redis.Redis.Client,limiterlib.StoreOptions{
		Prefix: config.GetString("app.name") + ":limiter",
	})
	if err != nil {
		logger.LogIf(err)
		return context, err
	}

	// 使用上面的初始化的 limiter.Rate 对象和存储对象
	limiterObj := limiterlib.New(store, rate)

	// 获取限流的结果
	return limiterObj.Get(c, key)

}
// imitor 的 Key，路由+IP，针对单个路由做限流
func GetKeyRouteWithIP(c *gin.Context) string {
	return routeToKeyString(c.FullPath()) + c.ClientIP()
}
func routeToKeyString(routeName string)  string {
	routeName = strings.ReplaceAll(routeName,"/","-")
	routeName = strings.ReplaceAll(routeName,":","_")
	return routeName
}