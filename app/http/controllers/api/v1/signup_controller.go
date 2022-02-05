package v1

import (
	"fmt"
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
	// 请求对象
	//type PhoneExistRequest struct {
	//	Phone string `json:"phone"`
	//}
	requst := requests.SignupPhoneExistRequest{}

	// 解析Json请求
	if err := c.ShouldBindJSON(&requst); err != nil {
		// 解析失败,返回422 状态码和错误信息
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity,gin.H{
			"error":err.Error(),
		})
		// 打印错误信息
		fmt.Println(err.Error())
		// 出错了，中断请求
		return
	}

	// 表单验证
	errs := requests.ValidateSignupPhoneExist(&requst,c)
	// errs 返回长度等于0 即通过 大于0有错误发生
	if len(errs) > 0{
		// 验证失败，返回422 状态码和错误信息
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity,gin.H{
			"errors":errs,
		})
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"exist":user.IsPhoneExist(requst.Phone),
	})
}

// 验证邮箱是否已经注册
func (sc *SignupController) IsEmailExist(c *gin.Context) {
	// 初始化请求对象
	request := requests.SignupEmailExistRequest{}

	// 解析Json请求对象
	if err := c.ShouldBindJSON(&request);err != nil {
		// 解析失败 返回422错误
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity,gin.H{
			"error":err.Error(),
		})
		// 打印错误信息
		fmt.Println(err.Error())
		return
	}
	// 表单验证
	errs := requests.ValidateSignupEmailExist(&request,c)
	if len(errs) > 0 {
		// 验证失败 返回422
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity,gin.H{
			"errors":errs,
		})
		return
	}

	// 检查数据库并返回响应
	c.JSON(http.StatusOK,gin.H{
		"exist":user.IsEmailExist(request.Email),
	})
}