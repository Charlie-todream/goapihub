package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goapihub/app/models/user"
	"goapihub/pkg/config"
	"goapihub/pkg/jwt"
	"goapihub/pkg/response"
)

func AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims,err := jwt.NewJWT().ParserToken(c)

		if err != nil {
			response.Unauthorized(c,fmt.Sprintf("请查看%v 相关的接口认证文档",config.GetString("app.name")))
			return
		}

		userModel := user.Get(claims.UserID)

		if userModel.ID == 0 {
			response.Unauthorized(c,"找不到对应用户,用户可能删除")
			return
		}

		// 将用户 gin.context 里 后续 auth 包将从这里拿到当前用户数据

		c.Set("current_user_id", userModel.GetStringID())
		c.Set("current_user_name", userModel.Name)
		c.Set("current_user", userModel)
		c.Next()

	}
}