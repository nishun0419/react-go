package main

import (
    "github.com/gin-gonic/gin"
    c "github.com/nishun0419/react-go/controllers"
)

func main() {
        r := gin.Default()
        r.LoadHTMLGlob("views/*")
        r.GET("/", c.GetUsers)
        r.POST("/user/post", c.PostUser)
        r.GET("/user/delete/:id", c.DeleteUser)
        r.Run(":8888") 
}