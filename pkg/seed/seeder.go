package seed

import (
	"goapihub/pkg/console"
	"goapihub/pkg/database"
	"gorm.io/gorm"
)

type SeederFunc func(db *gorm.DB)

type Seeder struct {
	Func SeederFunc
	Name string
}

var seeders []Seeder
var orderedSeederNames []string

func Add(name string,fn SeederFunc)  {
	seeders = append(seeders,Seeder{
		Name:name,
		Func: fn,
	})
}

func SetRunOrder(names []string)  {
	orderedSeederNames = names
}

// 通过名称来获取seeder对象
func GetSeeder(name string) Seeder  {
	for _,sdr := range seeders {
		if name == sdr.Name {
			return sdr
		}
	}
	return Seeder{}
}

// RunAll 运行所有的Seeder

func RunAll()  {
	// 先运行ordered
	executed := make(map[string]string)
    // 先运行 ordered
	for _,name := range orderedSeederNames {
		sdr := GetSeeder(name)
		if len(sdr.Name) > 0 {
			console.Warning("runnning ordered Seeder:" + sdr.Name)
			sdr.Func(database.DB)
			executed[name]= name
		}
	}

	// 在再运行剩下的
	for _,sdr := range seeders {
		if _,ok := executed[sdr.Name];!ok {
			console.Warning("running Seeder: "+ sdr.Name)
			sdr.Func(database.DB)
		}
	}
}

// 运行单个 Seeder
func RunSeeder(name string)  {
	for _,sdr := range seeders {
		if name == sdr.Name {
			sdr.Func(database.DB)
			break
		}
	}
}
