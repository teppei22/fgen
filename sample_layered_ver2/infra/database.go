package infra

import (
	"fmt"
	"os"
	"sample_layered/domain/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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
		dbPort   = os.Getenv("DB_PORT")
		dbName   = os.Getenv("DB_NAME")
		user     = os.Getenv("DB_USERNAME")
		password = os.Getenv("DB_PASSWORD")
	)
	driverName := "mysql"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, dbPort, dbName)
	fmt.Println(host, port, dbName, user, password)
	fmt.Println(dsn)

	db, err := gorm.Open(driverName, dsn)
	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(model.Task{})

	return db
}

func DBClose(db *gorm.DB) error {
	if err := db.Close(); err != nil {
		return err
	}
	return nil
}
