package db

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

// init Initialize DB
func init() {
	db, err := gorm.Open("postgres", "host=localhost dbname=noredd-app user=noredduser password=noredduser sslmode=disable port=5432")
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	DB = db
}

func GetDB() *gorm.DB{
	return DB
}