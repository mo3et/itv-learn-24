package util

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
}

func NewMysqlDB() *gorm.DB {
	dsn := "root:kratosworld@tcp(127.0.0.1:3306)/myschool?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect.")
		// panic("failed to connect.")
	}
	if err := db.AutoMigrate(&Student); err != nil {
		panic(err)
	}

	return db
}
