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

// 通过迁移文件的名称来获取 MigrationFile 对象
func getMigrationFile(name string) MigrationFile  {
	for _,mfile := range migrationFiles {
		if name == mfile.FileName {
			return mfile
		}
	}
	return MigrationFile{}
}

// isNotMigrated 判断迁移是否执行
func (mfile MigrationFile) isNotMigrated(migrations []Migration) bool {
	for _,migration := range migrations {
		if migration.Migration == mfile.FileName {
			return false
		}
	}
	return true
}
