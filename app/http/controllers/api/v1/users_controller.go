package v1

import (
    "github.com/gin-gonic/gin"
    "goapihub/app/models/user"
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

func (ctrl *UsersController) Index(c *gin.Context)  {
    data := user.All()
    response.Data(c, data)
}