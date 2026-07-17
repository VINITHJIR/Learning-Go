package database

import (
	"fmt"
	"log"

	"user-management-api/internal/config"
	"user-management-api/internal/domain"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.AppConfig.DBUser,
		config.AppConfig.DBPassword,
		config.AppConfig.DBHost,
		config.AppConfig.DBPort,
		config.AppConfig.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database Connection Failed : ", err)
	}

	DB = db
	err = DB.AutoMigrate(&domain.User{})
	if err != nil {
		log.Fatal("Migration Failed : ", err)
	}

	log.Println("✅ User Table Migrated Successfully")
	log.Println("✅ MySQL Connected Successfully")
}
