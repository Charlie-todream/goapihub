package v1

import (
	"github.com/gin-gonic/gin"
	"goapihub/pkg/captcha"
	"goapihub/pkg/logger"
	"net/http"
)

// 用户控制器
type VerifyCodeController struct {
	BaseAPIController
}

// 生成显示图片验证码
func (vc *VerifyCodeController) ShowCaptcha(c *gin.Context)  {
	// 生产验证码
	id,b64s,err := captcha.NewCaptcha().GenerateCaptcha()
	logger.LogIf(err)
	// 返回给用户
	c.JSON(http.StatusOK,gin.H{
		"captcha_id":id,
		"captcha_image":b64s,
	})
}