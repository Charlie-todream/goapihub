package v1

import (
	"github.com/gin-gonic/gin"
	"goapihub/app/models/user"
	"goapihub/app/requests"
	"goapihub/pkg/response"
)

type PasswordController struct {
	BaseAPIController
}

func (pc *PasswordController) ResetByPhone(c *gin.Context)  {
	requst := requests.ResetByPhoneRequest{}

	if ok := requests.Validate(c,&requst,requests.ResetByPhone); !ok {
		return
	}

	// 更新密码
	userModel := user.GetByPhone(requst.Phone)

	if userModel.ID == 0 {
		response.Abort404(c)
	}else {
		userModel.Password = requst.Password
		userModel.Save()
		response.Success(c)
	}

}

