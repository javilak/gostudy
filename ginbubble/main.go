package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func initMySQL() (db *gorm.DB, err error) {
	dsn := "*:*@(localhost:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return db, err
}

func main() {
	db, err := initMySQL()

	//dsn := "*:*@(localhost:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//if err != nil {
	//	panic(err)
	//}
	err = db.AutoMigrate(&Todo{})
	if err != nil {
		fmt.Printf("AutoMigrate failed %v/n", err)
	}

	r := gin.Default()
	r.Static("/static", "static")
	r.LoadHTMLGlob("template/*")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	v1Group := r.Group("v1")
	{
		v1Group.POST("/todo", func(c *gin.Context) { //增
			var todo Todo
			c.BindJSON(&todo)
			if err = db.Create(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, todo)
			}

		})

		v1Group.GET("/todo", func(c *gin.Context) {
			var todolist []Todo
			if err = db.Find(&todolist).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{
					"error": err,
				})
			} else {
				c.JSON(http.StatusOK, todolist)
			}
		})

		v1Group.GET("/todo/:id", func(c *gin.Context) {

		})

		v1Group.PUT("/todo/:id", func(c *gin.Context) {
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{
					"error": "id不存在",
				})
				return
			}
			var todo Todo
			if err := db.Where("id=?", id).First(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
				return
			}
			c.BindJSON(&todo)
			if err = db.Save(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, todo)
			}
		})

		v1Group.DELETE("/todo/:id", func(c *gin.Context) {
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{
					"error": "id不存在",
				})
				return
			}
			if err := db.Where("id=?", id).Delete(Todo{}).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
				return
			} else {
				c.JSON(http.StatusOK, gin.H{id: "deleted"})
			}
		})
	}
	r.Run()
}
