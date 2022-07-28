package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Student struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/go-crud-demo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			// 解决查表的时候会自动添加复数的问题 , 例如 user 变成了 users
			SingularTable: true,
		},
	})
	fmt.Println(db)
	fmt.Println(err)
	r := gin.Default()

	r.GET("/student/list", func(c *gin.Context) {
		var data []Student
		result := db.Find(&data)
		if result != nil {
			db.Create(data)
			c.JSON(200, gin.H{
				"code": 200,
				"msg":  "新增成功",
				"data": data,
			})
		}
	})

	r.POST("/student", func(c *gin.Context) {
		var data Student
		err := c.ShouldBindJSON(&data)
		if err == nil {
			db.Select("Name", "Age").Create(&data)
			c.JSON(200, gin.H{
				"code": 200,
				"msg":  "新增成功",
			})
		}
	})

	r.PUT("/student", func(c *gin.Context) {
		var data Student
		err := c.ShouldBindJSON(&data)
		if err == nil {
			db.Where("id=?", data.Id).Updates(&data)
			c.JSON(200, gin.H{
				"code": 200,
				"msg":  "更新成功",
			})
		}
	})

	r.DELETE("/student/:id", func(c *gin.Context) {
		id := c.Param("id")
		db.Delete(Student{}, id)
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "删除成功",
		})
	})

	port := ":8080"
	r.Run(port)
}
