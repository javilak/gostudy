package main

import (
	"fmt"
	"ginbubble/controller"
	"ginbubble/dao"
	"ginbubble/models"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := dao.InitMySQL()

	//dsn := "*:*@(localhost:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//if err != nil {
	//	panic(err)
	//}
	err = dao.DB.AutoMigrate(&models.Todo{})
	if err != nil {
		fmt.Printf("AutoMigrate failed %v/n", err)
	}

	r := gin.Default()
	r.Static("/static", "static")
	r.LoadHTMLGlob("template/*")
	r.GET("/", controller.IndexHandler)

	v1Group := r.Group("v1")
	{
		v1Group.POST("/todo", controller.V1Post)

		v1Group.GET("/todo", controller.V1Get)

		v1Group.PUT("/todo/:id", controller.V1Put)

		v1Group.DELETE("/todo/:id", controller.V1Deleted)
	}
	r.Run()
}
