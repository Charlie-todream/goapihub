package topic

import (
    "github.com/gin-gonic/gin"
    "goapihub/pkg/app"
    "goapihub/pkg/database"
    "goapihub/pkg/paginator"
)

func Get(idstr string) (topic Topic) {
    database.DB.Where("id", idstr).First(&topic)
    return
}

func GetBy(field, value string) (topic Topic) {
    database.DB.Where("? = ?", field, value).First(&topic)
    return
}

func All() (topics []Topic) {
    database.DB.Find(&topics)
    return
}

func IsExist(field, value string) bool {
    var count int64
    database.DB.Model(Topic{}).Where(" = ?", field, value).Count(&count)
    return count > 0
}
// Paginate 分页
// Paginate 分页内容
func Paginate(c *gin.Context, perPage int) (users []Topic, paging paginator.Paging) {
    paging = paginator.Paginate(
        c,
        database.DB.Model(Topic{}),
        &users,
        app.V1URL(database.TableName(&Topic{})),
        perPage,
    )
    return
}
