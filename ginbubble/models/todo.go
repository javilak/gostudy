package models

import "ginbubble/dao"

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func Createtodo(todo *Todo) (err error) {
	if err = dao.DB.Create(&todo).Error; err != nil {
		return err
	}
	return
}

func GetTodoList() (todolist []*Todo, err error) {
	if err = dao.DB.Find(&todolist).Error; err != nil {
		return nil, err
	}
	return
}

func GetTodoById(id int) {
	todo := &Todo{}
	if err := dao.DB.Where("id=?", id).First(&todo).Error; err != nil {

	}
}
