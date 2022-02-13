package v1

import (
	"github.com/gin-gonic/gin"
	"goapihub/app/models/user"
	"goapihub/app/requests"
	"goapihub/pkg/response"
)

type  SignupController struct {
	BaseAPIController
}

// 检验手机是否被注册
func (sc *SignupController) IsPhoneExist(c *gin.Context)  {

	// 获取参数，并做表单验证
	request := requests.SignupPhoneExistRequest{}
	if ok := requests.Validate(c,&request,requests.SignupPhoneExist); !ok {
		return
	}

	// 检查数据库并返回值

	response.JSON(c, gin.H{
		"exist": user.IsPhoneExist(request.Phone),
	})
}

// 验证邮箱是否已经注册
func (sc *SignupController) IsEmailExist(c *gin.Context) {

	// 初始化请求对象
	request := requests.SignupEmailExistRequest{}
	if ok := requests.Validate(c,&request,requests.SignupEmailExist);!ok {
		return
	}
	// 检查数据库并返回响应
	response.JSON(c, gin.H{
		"exist": user.IsEmailExist(request.Email),
	})
}

// SignupUsingEmail 使用 Email + 验证码进行注册
func (sc *SignupController) SignupUsingEmail(c *gin.Context) {

	// 1. 验证表单
	request := requests.SignupUsingEmailRequest{}
	if ok := requests.Validate(c, &request, requests.SignupUsingEmail); !ok {
		return
	}

	// 2. 验证成功，创建数据
	userModel := user.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}
	userModel.Create()

	if userModel.ID > 0 {
		response.CreatedJSON(c, gin.H{
			"data": userModel,
		})
	} else {
		response.Abort500(c, "创建用户失败，请稍后尝试~")
	}
}