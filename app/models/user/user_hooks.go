package user

import (
	"goapihub/pkg/hash"
	"gorm.io/gorm"
)

// BeforeSave Gorm 模型的在创建和更新模型前调用
func (userModel *User) BeforeSave(tx *gorm.DB) (err error)  {

	if !hash.BcryptIsHashed(userModel.Password) {
		userModel.Password = hash.BcryptHash(userModel.Password)
	}

	return
}