package controller

import (
	"ginbubble/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func V1Post(c *gin.Context) { //增
	var todo models.Todo
	c.BindJSON(&todo)
	if err := models.Createtodo(&todo); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, todo)
	}

}

func V1Get(c *gin.Context) {

	if todolist, err := models.GetTodoList(); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	} else {
		c.JSON(http.StatusOK, todolist)
	}
}

func V1Put(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"error": "id不存在",
		})
		return
	}
	var todo Todo
	{
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.BindJSON(&todo)
	if err = db.Save(&todo).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func V1Deleted(c *gin.Context) {
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
}
