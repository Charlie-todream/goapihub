package user

import "goapihub/pkg/database"

// 模型相关的数据库操作

// 判断 Email 是否已经被注册
func IsEmailExist(email string) bool  {
	var count int64
	database.DB.Model(User{}).Where("email = ?",email).Count(&count)
	return count > 0
}

// IsPhoneExist 判断手机号是否已经注册
func IsPhoneExist(phone string) bool {
	var count int64
	database.DB.Model(User{}).Where("phone=?",phone).Count(&count)
	return count > 0
}

// GetByPhone 通过手机号来获取用户
func GetByPhone(phone string) (userModel User)  {
	database.DB.Where("phone = ?",phone).First(&userModel)
	return
}

// 通过手机号 /Email/用户名来获取用户
func GetByMulti(loginID string) (userModel User)  {
	database.DB.Where("phone = ?",loginID).Or("email=?",loginID).Or("name=?",loginID).First(&userModel)
	return
}