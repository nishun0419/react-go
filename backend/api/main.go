package main

import (
    "time"
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
    c "github.com/nishun0419/react-go/controllers"
)

func main() {
        r := gin.Default()
	    // r.Use() で使用したいミドルウェアなどの指定ができる。
	    // https://godoc.org/github.com/gin-gonic/gin#RouterGroup.Use
	    r.Use(cors.New(cors.Config{
	        // 許可したいHTTPメソッドの一覧
	        AllowMethods: []string{
	            "POST",
	            "GET",
	            "OPTIONS",
	            "PUT",
	            "DELETE",
	        },
	        // 許可したいHTTPリクエストヘッダの一覧
	        AllowHeaders: []string{
	            "Access-Control-Allow-Headers",
	            "Content-Type",
	            "Content-Length",
	            "Accept-Encoding",
	            "X-CSRF-Token",
	            "Authorization",
	        },
	        // 許可したいアクセス元の一覧
	        AllowOrigins: []string{
	            "http://localhost:3000",
	        },
	        // 自分で許可するしないの処理を書きたい場合は、以下のように書くこともできる
	        // AllowOriginFunc: func(origin string) bool {
	        //  return origin == "https://www.example.com:8080"
	        // },
	        // preflight requestで許可した後の接続可能時間
	        // https://godoc.org/github.com/gin-contrib/cors#Config の中のコメントに詳細あり
	        MaxAge: 24 * time.Hour,
	    }))
        r.GET("/", c.GetTodos)
        r.POST("/todo/post", c.PostTodo)
        r.GET("/todo/delete/:id", c.DeleteTodo)
        r.Run(":8888") 
}