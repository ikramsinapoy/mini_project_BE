package configs

import (
	"foodcal/models/foods"
	"foodcal/models/transactions"
	"foodcal/models/users"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "root:Baruga2018@tcp(127.0.0.1:3306)/foodcal?charset=utf8mb4&parseTime=True&loc=Local"
	var err error

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Database tidak connect")
	}
	Migration()
}

func Migration() {
	DB.AutoMigrate(users.User{}, transactions.Transaction{}, foods.Food{})
}
