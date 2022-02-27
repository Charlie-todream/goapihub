package seeders

import (
	"fmt"
	"goapihub/database/factories"
	"goapihub/pkg/console"
	"goapihub/pkg/logger"
	"goapihub/pkg/seed"
	"gorm.io/gorm"
)

func init()  {
	seed.Add("SeedUsersTable", func(db *gorm.DB) {
		// 创建10个用户对象
		users := factories.MakeUsers(10)
		// 批量创建用户
		result := db.Table("users").Create(&users)

		// 记录错误
		if err := result.Error;err != nil {
			logger.LogIf(err)
			return
		}

		// 打印运行情况
		console.Success(fmt.Sprintf("Table [%v] %v rows seeded", result.Statement.Table, result.RowsAffected))
	})
}