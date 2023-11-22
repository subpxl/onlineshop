package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() error { // conenction string
	dsn := "host=localhost user=onlineshop password=onlineshop dbname=onlineshop port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("error", err)
	}
	fmt.Println("db loaded")
	DB = db

	return nil
}

func SyncDB() {
	DB.AutoMigrate(Category{})
	DB.AutoMigrate(Product{})
	DB.AutoMigrate(User{})
	DB.AutoMigrate(Order{})
	DB.AutoMigrate(OrderItem{})
	DB.AutoMigrate(Role{})
	DB.AutoMigrate(Address{})
	DB.AutoMigrate(Payment{})
	fmt.Println("sync db")
}
