package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
)

func main() {

	engine := gin.Default()

	v1 := engine.Group("/v1")
	{
		v1.GET("/login", WebRoot)
		v1.GET("/submit", WebRoot)
		v1.GET("/read", WebRoot)
	}

	v2 := engine.Group("/v2", middleware1, middleware2)
	{
		v2.GET("/login", WebRoot)
		v2.GET("/submit", WebRoot)
		v2.GET("/read", WebRoot)
	}

	engine.Run(":122")
}


func WebRoot(context *gin.Context) {
	urls := context.Request.URL.Path
	id  := context.Query("id")

	messges := urls + "/" + id

	context.String(http.StatusOK, messges)
	log.Println("exec handler")
}

func middleware1(context *gin.Context)  {
	log.Print("执行中间逻辑")
	context.Next()
}

func middleware2(context *gin.Context)  {
	log.Print("执行中间逻辑---------2")
	context.Next()

	log.Println("exec middleware2")
}


