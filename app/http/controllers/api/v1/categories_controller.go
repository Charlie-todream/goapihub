package v1

import (
    "goapihub/app/models/category"
    "goapihub/app/requests"
    "goapihub/pkg/response"

    "github.com/gin-gonic/gin"
)

type CategoriesController struct {
    BaseAPIController
}

func (ctrl *CategoriesController) Index(c *gin.Context) {
    categories := category.All()
    response.Data(c, categories)
}


func (ctrl *CategoriesController) Store(c *gin.Context) {

    request := requests.CategoryRequest{}
    if ok := requests.Validate(c, &request, requests.CategorySave); !ok {
        return
    }

    categoryModel := category.Category{
        Name: request.Name,
        Description: request.Description,
    }

    categoryModel.Create()
    if categoryModel.ID > 0 {
        response.Created(c, categoryModel)
    } else {
        response.Abort500(c, "创建失败，请稍后尝试~")
    }
}
