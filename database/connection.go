package database

import (
	"test/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(){
	connection, err := gorm.Open(mysql.Open("root:@/technicaltest"), &gorm.Config{})

	if err != nil {
		panic("Database tidak Terhubung")
	}

	DB = connection

	connection.AutoMigrate(&model.User{})
}
