package migrate

import (
	"database/sql"
	"gorm.io/gorm"
)

type migrationFunc func(migrator gorm.Migrator,db *sql.DB)

// 所有迁移文件数组
var migrationFiles []MigrationFile

// 代表单个迁移文件
type MigrationFile struct {
	Up migrationFunc
	Down migrationFunc
	FileName string
}

// 增加一个迁移文件，所有的迁移文件都需要调用此方法来注册
func Add(name string,up migrationFunc,down migrationFunc)  {
	migrationFiles = append(migrationFiles,MigrationFile{
		FileName: name,
		Up: up,
		Down: down,
	})
}