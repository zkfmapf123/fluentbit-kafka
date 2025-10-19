package config

import (
	"fmt"
	"os"

	"github.com/zkfmapf123/fpg/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBConn struct {
	DB *gorm.DB
}

func NewDBConn() *DBConn {

	user := os.Getenv("MYSQL_USER")
	host := os.Getenv("MYSQL_HOST")
	password := os.Getenv("MYSQL_PASSWORD")
	port := os.Getenv("MYSQL_PORT")
	dbName := os.Getenv("MYSQL_DATABASE")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return &DBConn{
		DB: db,
	}
}

func (d *DBConn) CreateTable() {

	d.DB.AutoMigrate(&models.User{})

}
