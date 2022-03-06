package v1

import (
    "goapihub/app/models/link"
    "goapihub/pkg/response"

    "github.com/gin-gonic/gin"
)

type LinksController struct {
    BaseAPIController
}

func (ctrl *LinksController) Index(c *gin.Context) {
    links := link.All()
    response.Data(c, links)
}




