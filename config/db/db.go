package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

// init Initialize DB
func init() {
	dbuser := os.Getenv("DBUSER")
	dbpassword := os.Getenv("DBPASSWD")

	db, err := gorm.Open("postgres", fmt.Sprintf("host=localhost dbname=noredd-app user=%s password=%s sslmode=disable port=5432", dbuser, dbpassword))
	if err != nil {
		panic(err)
	}
	db.LogMode(true)
	DB = db
}

func GetDB() *gorm.DB {
	return DB
}
