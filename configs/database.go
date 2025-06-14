package configs

import (
	"log"
	"myapp/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(){
	dsn := "root:@tcp(127.0.0.1:3306)/db_fiber?charset=utf8mb4&parseTime=True&loc=Local"
	database,err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	
	database.AutoMigrate(&models.User{})
	DB = database
}