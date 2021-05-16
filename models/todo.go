package models

import (
	"GinCode/dao"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"` // 完成状态
}


func CreateTodo(todo *Todo) error {
	return dao.DB.Debug().Create(&todo).Error
}

func GetTodoList() ([]Todo, error) {
	var todoList []Todo
	return todoList, dao.DB.Debug().Find(&todoList).Error
}

func GetTodoById(id string) (todo *Todo, err error) {
	todo = new(Todo)
	return todo, dao.DB.Debug().Where("id=?", id).First(&todo).Error
}

func UpdateTodo(todo *Todo) error {
	return dao.DB.Debug().Save(&todo).Error
}

func DeleteTodoById(id string) error {
	return dao.DB.Debug().Where("id=?", id).Delete(&Todo{}).Error
}
