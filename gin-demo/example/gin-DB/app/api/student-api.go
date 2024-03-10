package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mo3et/itv-learn-24/gin-demo/example/gin-DB/data"
)

// 新增学生
/*
http://127.0.0.1:8080/student/save
{
	"id":"1",
	"name":"monet",
	"age":24
}
*/

// type Student struct {
// 	// gorm.Model
// 	ID   string `gorm:"primaryKey;column:id" json:"id"`
// 	Name string `gorm:"column:name" json:"name"`
// 	Age  int    `gorm:"column:age" json:"age"`
// }

// func (Student) TableName() string {
// 	return "student"
// }

// var db *gorm.DB = data.GetDB()

var stu data.Student

// var stu Student

func Save(c *gin.Context) {
	// var stu Student
	db := data.GetDB()
	if err := c.ShouldBindJSON(&stu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	log.Printf("Current db value: %+v", db)

	result := db.Create(&stu)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error.Error(),
		})
		log.Printf("Create failed. %v", result.Error.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": "添加成功",
	})
	log.Print("create success.")
}

func SelectById(c *gin.Context) {
	// var stu Student
	db := data.GetDB()
	id := c.Query("id")
	if err := db.First(&stu, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": stu,
	})
}

func Update(c *gin.Context) {
	// var stu Student
	db := data.GetDB()
	if err := c.ShouldBindJSON(&stu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	// db.Model(&stu).Where("id =?", stu.ID).Updates(map[string]interface{}{"Name": stu.Name, "Age": stu.Age})
	if err := db.Model(&stu).Where("id =?", stu.ID).Updates(data.Student{Name: stu.Name, Age: stu.Age}).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": "Update success."})
}

func Delete(c *gin.Context) {
	db := data.GetDB()
	id := c.Query("id")
	if err := db.Delete(&data.Student{}, id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": "Delete success."})
}
