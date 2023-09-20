package main

import (
	"fmt"
	"foods-spider-server/database/mysql"
	"foods-spider-server/pkg"
	"gorm.io/gen"
	"gorm.io/gorm"
	"log"
)

func Gen(db *gorm.DB) {
	g := gen.NewGenerator(gen.Config{
		OutPath:      "internal/model/dal",
		ModelPkgPath: "model/models",
		OutFile:      "",
		Mode:         gen.WithoutContext | gen.WithDefaultQuery,
	})
	g.UseDB(db)
	g.ApplyBasic(g.GenerateAllTable()...)
	g.Execute()
}
func main() {
	// 初始化配置
	pkg.InitConfig()

	// 链接数据库拿到DB
	db, err := mysql.InitSqlxDB()
	fmt.Print(err)
	if err != nil {
		log.Fatal(err)
	}
	Gen(db)
}
