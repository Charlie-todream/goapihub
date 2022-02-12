package v1

import (
	"github.com/gin-gonic/gin"
	"goapihub/app/requests"
	"goapihub/pkg/captcha"
	"goapihub/pkg/logger"
	"goapihub/pkg/response"
	"goapihub/pkg/verifycode"
)

// 用户控制器
type VerifyCodeController struct {
	BaseAPIController
}

// 生成显示图片验证码
func (vc *VerifyCodeController) ShowCaptcha(c *gin.Context) {
	// 生产验证码
	id, b64s, err := captcha.NewCaptcha().GenerateCaptcha()
	logger.LogIf(err)
	// 返回给用户
	response.JSON(c, gin.H{
		"captcha_id":    id,
		"captcha_image": b64s,
	})

}

// 发送手机验证码
func (vc *VerifyCodeController) SendUsingPhone(c *gin.Context) {

	// 验证表单
	request := requests.VerifyCodePhoneRequest{}
	if ok := requests.Validate(c, &request, requests.VerifyCodePhone); !ok {
		return
	}

	// 2 发送SMS
	if ok := verifycode.NewVerifyCode().SendSMS(request.Phone); !ok {
		response.Abort500(c, "发送短信失败")
	} else {
		response.Success(c)
	}

}
