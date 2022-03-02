package infra

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
)

type DB struct {
	Host       string
	UserName   string
	Password   string
	DBName     string
	Connection *gorm.DB
}

func DBConnect() *gorm.DB {
	var (
		host     = os.Getenv("DB_HOST")
		port     = os.Getenv("PORT")
		dbName   = os.Getenv("DB_USERNAME")
		user     = os.Getenv("DB_USERNAME")
		password = os.Getenv("DB_PASSWORD")
	)
	driverName := "mysql"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbName)

	// var dsnVer1 string = fmt.Sprintf("%s:%s:%s",)

	db, err := gorm.Open(driverName, dsn)
	if err != nil {
		panic(err.Error())
	}
	return db
}
