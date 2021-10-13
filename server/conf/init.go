package conf

import (
	"app/server/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var Db *gorm.DB

func init() {
	initDb()
}

func initDb() {
	nameStrategy := schema.NamingStrategy{
		SingularTable: true,
	}
	db, err := gorm.Open(sqlite.Open("./app.db"), &gorm.Config{
		NamingStrategy: nameStrategy,
	})
	_ = db.AutoMigrate(&model.Source{}, &model.BookRule{})
	if err != nil {
		panic(err)
	}
	Db = db
}
