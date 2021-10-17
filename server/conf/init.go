package conf

import (
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
	db, err := gorm.Open(sqlite.Open("./resources/app.db"), &gorm.Config{
		NamingStrategy: nameStrategy,
	})
	if err != nil {
		panic(err)
	}
	Db = db
}
