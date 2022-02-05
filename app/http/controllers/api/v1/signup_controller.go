package v1

import (
	"github.com/gin-gonic/gin"
	"goapihub/app/models/user"
	"goapihub/app/requests"
	"net/http"
)

type  SignupController struct {
	BaseAPIController
}

// 检验手机是否被注册
func (sc *SignupController) IsPhoneExist(c *gin.Context)  {

	// 获取参数，并做表单验证
	requst := requests.SignupPhoneExistRequest{}
	if ok := requests.Validate(c,&requst,requests.SignupPhoneExist); !ok {
		return
	}

	// 检查数据库并返回值

	c.JSON(http.StatusOK,gin.H{
		"exist":user.IsPhoneExist(requst.Phone),
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
	c.JSON(http.StatusOK,gin.H{
		"exist":user.IsEmailExist(request.Email),
	})
}