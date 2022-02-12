package user

import (
	"goapihub/app/models"
	"goapihub/pkg/database"
)

// User 用户模型
type User struct {
	models.BaseModel

	Name string `json:"name,omitempty"`
	Email string `json:"-"`
	Phone string `json:"-"`
	Password string `json:"-"` // 指示json解析器忽略字段 接口返回数据 三个字段会被隐藏
}

func (userModel *User) Create()  {
	database.DB.Create(&userModel)
}