package v1

import (
    "github.com/gin-gonic/gin"
    "goapihub/app/models/user"
    "goapihub/app/requests"
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

    request := requests.PaginationRequest{}
    if ok := requests.Validate(c, &request, requests.Pagination); !ok {
        return
    }

    data, pager := user.Paginate(c, 10)
    response.JSON(c, gin.H{
        "data":  data,
        "pager": pager,
    })
}

