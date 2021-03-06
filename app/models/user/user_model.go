package user

import (
	"goapihub/app/models"
	"goapihub/pkg/database"
	"goapihub/pkg/hash"
)

// User 用户模型
type User struct {
	models.BaseModel

	City          string `json:"city,omitempty"`
	Introduction string `json:"introduction,omitempty"`
	Avatar        string `json:"avatar,omitempty"`

	Name string `json:"name,omitempty"`
	Email string `json:"-"`
	Phone string `json:"-"`
	Password string `json:"-"` // 指示json解析器忽略字段 接口返回数据 三个字段会被隐藏
	models.CommonTimestampsField
}

func (userModel *User) Create()  {
	database.DB.Create(&userModel)
}
// ComparePassword 密码是否正确
func (userModel *User) ComparePassword(_password string) bool {
	return hash.BcryptCheck(_password, userModel.Password)
}

// Get 通过 ID 获取用户
func Get(idstr string) (userModel User) {
	database.DB.Where("id", idstr).First(&userModel)
	return
}

func (userModel *User) Save() (rowsAffected int64) {
	result := database.DB.Save(&userModel)
	return result.RowsAffected
}

