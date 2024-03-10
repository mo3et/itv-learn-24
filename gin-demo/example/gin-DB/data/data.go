package data

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Tabler interface {
	TableName() string
}

type Student struct {
	ID   string `gorm:"primaryKey;column:id" json:"id"`
	Name string `gorm:"column:name" json:"name"`
	Age  int    `gorm:"column:age" json:"age"`
}

func (Student) TableName() string {
	return "student"
}

// 全局变量 才不会因为局部更改而消失
// 使用了局部变量 而不会对外部产生影响
var db *gorm.DB

func Init() {
	dsn := "root:kratosworld@tcp(127.0.0.1:3306)/myschool?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect.")
		// panic("failed to connect.")
	}

	if err := db.AutoMigrate(&Student{}); err != nil {
		panic(err)
	}
	log.Print("DB success run!")
}

func GetDB() *gorm.DB {
	return db
}
