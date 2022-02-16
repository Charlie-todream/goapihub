package v1

import (
	"github.com/gin-gonic/gin"
	"goapihub/app/requests"
	"goapihub/pkg/auth"
	"goapihub/pkg/jwt"
	"goapihub/pkg/response"
)

type LoginController struct {
	BaseAPIController
}

func (lc *LoginController) LoginByPhone(c *gin.Context)  {
	// 1. 验证表单
	request := requests.LoginByPhoneRequest{}
	if ok := requests.Validate(c,&request,requests.LoginByPhone); !ok {
		return
	}
	// 尝试登陆
	user,err := auth.LoginByPhone(request.Phone)
	if err != nil {
		response.Error(c,err,"账号不存在或者密码错误")
	}else {
		token := jwt.NewJWT().IssueToken(user.GetStringID(),user.Name)
		response.JSON(c,gin.H{
			"token":token,
		})
	}
}

// LoginByPassword 多种方法登录，支持手机号、email 和用户名
func (lc *LoginController) LoginByPassword(c *gin.Context) {
	// 1. 验证表单
	request := requests.LoginByPasswordRequest{}
	if ok := requests.Validate(c, &request, requests.LoginByPassword); !ok {
		return
	}

	// 2. 尝试登录
	user, err := auth.Attempt(request.LoginID, request.Password)
	if err != nil {
		// 失败，显示错误提示
		response.Unauthorized(c, "登录失败")

	} else {
		token := jwt.NewJWT().IssueToken(user.GetStringID(), user.Name)
		response.JSON(c, gin.H{
			"token": token,
		})
	}
}