package main

import "github.com/gin-gonic/gin"


func main() {
	r := gin.Default()
	r.Static("./acc", "./statics")
	r.LoadHTMLGlob("./templates/*")

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Title": "Hello world!",
		})
	})
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
	r.Run(":8086")
}
