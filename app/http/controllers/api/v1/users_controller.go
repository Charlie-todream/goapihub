package v1

import (
    "github.com/gin-gonic/gin"
    "goapihub/pkg/auth"
    "goapihub/pkg/response"
)

type UsersController struct {
    BaseAPIController
}

func (ctrl *UsersController) CurrentUser(c *gin.Context)  {
    userModel := auth.CurrentUser(c)
    response.Data(c,userModel)
}