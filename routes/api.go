package routes

import (
	"github.com/gin-gonic/gin"
	auth "goapihub/app/http/controllers/api/v1"
)

func RegisterAPIRoutes(r *gin.Engine) {
	//
	v1 := r.Group("/v1")
	{

		// 授权相关
		authGroup := v1.Group("/auth")
		{
			suc := new(auth.SignupController)
			// 判断手机是否注册
			authGroup.POST("/signup/phone/exist", suc.IsPhoneExist)
			// 判断 Email 是否已注册
			authGroup.POST("/signup/email/exist", suc.IsEmailExist)
			// 发送验证码
			vcc := new(auth.VerifyCodeController)
			authGroup.POST("/verify-codes/captcha", vcc.ShowCaptcha)
			authGroup.POST("/verify-codes/phone", vcc.SendUsingPhone)
		}
	}
}
