package migrate

import (
	"goapihub/pkg/console"
	"goapihub/pkg/database"
	"goapihub/pkg/file"
	"gorm.io/gorm"
	"io/ioutil"
)

// 数据迁移操作类
type Migrator struct {
	Folder string
	DB *gorm.DB
	Migrator gorm.Migrator
}

// 对应数据的 migrations 表里的一条数据
type Migration struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement;"`
	Migration string `gorm:"type:varchar(255);not null;unique;"`
	Batch     int
}

// NewMigrator 创建 Migrator 实例，用以执行迁移操作
func NewMigrator() *Migrator {

	// 初始化必要属性
	migrator := &Migrator{
		Folder:   "database/migrations/",
		DB:       database.DB,
		Migrator: database.DB.Migrator(),
	}
	// migrations 不存在的话就创建它
	migrator.createMigrationsTable()

	return migrator
}

// 创建 migrations 表
func (migrator *Migrator) createMigrationsTable() {

	migration := Migration{}

	// 不存在才创建
	if !migrator.Migrator.HasTable(&migration) {
		migrator.Migrator.CreateTable(&migration)
	}
}

func (migrator *Migrator) Up()  {

	// 读取所有迁移文件，确保按照时间排序
	migrateFiles := migrator.readAllMigrationFiles()

	// 获取当前批次的值
	batch := migrator.getBatch()

	// 获取所有迁移数据
	migrations := []Migration{}
	migrator.DB.Find(&migrations)

	// 可以通过此值来判断数据库是否已是最新
	runed := false

	// 对迁移文件进行遍历，如果没有执行过，就执行 up 回调
	for _, mfile := range migrateFiles {

		// 对比文件名称，看是否已经运行过
		if mfile.isNotMigrated(migrations) {
			migrator.runUpMigration(mfile, batch)
			runed = true
		}
	}

	if !runed {
		console.Success("database is up to date.")
	}

}

// 获取当前这个批次的值
func (migrator *Migrator) getBatch() int {

	// 默认为 1
	batch := 1

	// 取最后执行的一条迁移数据
	lastMigration := Migration{}
	migrator.DB.Order("id DESC").First(&lastMigration)

	// 如果有值的话，加一
	if lastMigration.ID > 0 {
		batch = lastMigration.Batch + 1
	}
	return batch
}

// 从文件目录读取文件，保证正确的时间排序
func (migrator *Migrator) readAllMigrationFiles() []MigrationFile {
	// 读取 database/migrations/ 目录下的所有文件
	// 默认是会按照文件名称进行排序
	files,err := ioutil.ReadDir(migrator.Folder)
	console.ExitIf(err)

	var migrateFiles []MigrationFile

	for _, f := range files {
		// 去除文件后缀
		fileName := file.FileNameWithoutExtension(f.Name())
		// 通过迁移文件获取 MIgrationFile对象
		mfile := getMigrationFile(fileName)
		if len(mfile.FileName) > 0 {
			migrationFiles = append(migrateFiles,mfile)
		}
	}
	return migrateFiles
}

func (migrator *Migrator) runUpMigration(mfile MigrationFile,batch int)  {
	if mfile.Up != nil {

		// 友好提示
		console.Warning("migration" + mfile.FileName)
		// 执行UP方法
		mfile.Up(database.DB.Migrator(),database.SQLDB)

		// 提示已迁移了哪个文件
		console.Success("migrated " + mfile.FileName)
	}
	// 入库
	err := migrator.DB.Create(&Migration{
		Migration: mfile.FileName,
		Batch: batch,
	}).Error
	console.ExitIf(err)
}