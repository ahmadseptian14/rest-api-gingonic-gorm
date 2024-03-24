package database

import (
	"fmt"
	dbconfig "gin-gonic-gorm/configs/db_config"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase()  {
	var errConnection error

	if dbconfig.DB_DRIVER == "mysql" {
		dsnMysql := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbconfig.DB_USER, dbconfig.DB_PASSWORD, dbconfig.DB_HOST, dbconfig.DB_PORT, dbconfig.DB_NAME)
		// dsnMysql := "root:root@tcp(127.0.0.1:3306)/go_gin_gonic?charset=utf8mb4&parseTime=True&loc=Local"
		DB, errConnection = gorm.Open(mysql.Open(dsnMysql), &gorm.Config{})
	}

	if dbconfig.DB_DRIVER == "pgsql" {
		dsnPgsql := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", dbconfig.DB_HOST, dbconfig.DB_USER, dbconfig.DB_PASSWORD, dbconfig.DB_NAME, dbconfig.DB_PORT)
		DB, errConnection = gorm.Open(postgres.Open(dsnPgsql), &gorm.Config{})
	}

	if errConnection != nil {
		panic("Cant connect to database")
	}

	log.Println("Connected")
}
