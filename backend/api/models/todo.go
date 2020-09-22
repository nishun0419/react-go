package models

import "time"

type Todo struct {
	ID 		int
	Name 		string          `gorm:"not null"`
	CreatedAt	time.Time
	UpdatedAt	time.Time
	DeletedAt	*time.Time
}
func (t Todo) GetTodos() *[]Todo {
   db := GormDB()
  	if !db.HasTable(&Todo{}){
		db.CreateTable(&Todo{})
	}
   todos := &[]Todo{}
   db.Find(todos)
   return todos
}

func (t Todo) PostTodo(name string) Todo{
	db := GormDB()
	todo := Todo{Name:name}
	db.Create(&todo)
	return todo;
}

func(t Todo) DeleteTodo(id int){
	db := GormDB()
	todo := Todo{}

	db.First(&todo,id)
	db.Delete(&todo)
	return
}