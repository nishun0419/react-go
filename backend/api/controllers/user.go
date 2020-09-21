package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	m "github.com/nishun0419/react-go/models"
	"log"
)
func GetUsers(c *gin.Context){
   u := m.User{}
   users := u.GetUsers()
   c.HTML(http.StatusOK, "index.html", users)
}

func PostUser(c *gin.Context){
	u := m.User{}

	name,_ := c.GetPostForm("name")
	email,_ := c.GetPostForm("email")

	u.PostUser(name, email)

	c.Redirect(http.StatusMovedPermanently, "/")

}

func DeleteUser(c *gin.Context){
	u := m.User{}

	id,_ := strconv.Atoi(c.Param("id"))

	u.DeleteUser(id)

	log.Printf("%s","test")

	c.Redirect(http.StatusMovedPermanently, "/")
}