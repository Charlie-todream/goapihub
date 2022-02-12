package v1

import (
	"github.com/gin-gonic/gin"
	"goapihub/app/models/user"
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

func (vc *VerifyCodeController) SendUsingEmail(c *gin.Context) {

	// 1. 验证表单
	request := requests.VeifyCodeEmailRequest{}
	if ok := requests.Validate(c, &request, requests.VeifyCodeEmail); !ok {
		return
	}

	// 2. 发送 SMS
	err := verifycode.NewVerifyCode().SendEmail(request.Email)
	if err != nil {
		response.Abort500(c, "发送 Email 验证码失败~")
	} else {
		response.Success(c)
	}
}

func (sc *SignupController) SignupUsingPhone(c *gin.Context){
	// 1. 表单验证
	request := requests.SignupUsingPhoneRequest{}

	if ok := requests.Validate(c,&request,requests.SignupUsingPhone);!ok {
		return
	}

	// 2. 验证成功，创建数据
	_user := user.User{
		Name: request.Name,
		Phone: request.Phone,
		Password: request.Password,
	}
	_user.Create()

	if _user.ID >0 {
		response.CreatedJSON(c,gin.H{
			"data":_user,
		})
	}else {
		response.Abort500(c,"创建用户失败，请稍后尝试~")
	}
}
