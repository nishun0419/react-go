package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	m "github.com/nishun0419/react-go/models"
	"log"
)

type todoData struct {
    NAME   string `json:"name"`
}
func GetTodos(c *gin.Context){
   todo := m.Todo{}
   todos := todo.GetTodos()
   c.JSON(http.StatusOK, todos)
}

func PostTodo(c *gin.Context){
	todo := m.Todo{}
	var name todoData
	c.BindJSON(&name)
	log.Printf(name.NAME)
	restodo := todo.PostTodo(name.NAME)
	c.JSON(http.StatusOK, restodo)
}

func DeleteTodo(c *gin.Context){
	todo := m.Todo{}

	id,_ := strconv.Atoi(c.Param("id"))

	todo.DeleteTodo(id)

	log.Printf("%s","test")

}