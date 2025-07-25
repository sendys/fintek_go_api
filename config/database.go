package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DatabaseConfig holds the configuration for the database connection.

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:bismillah@tcp(127.0.0.1:3306)/fintek_shared?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal konek database: ", err)
	}

	DB = database
}
