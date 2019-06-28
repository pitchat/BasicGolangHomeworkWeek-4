package main

import (
	"github.com/pitchat/test4/todo"
	"github.com/pitchat/test4/database"
	"github.com/gin-gonic/gin"
)

func main(){
	
	database.InitDB()

	router := gin.Default()
	api := router.Group("/api")
	api.GET("/todos", todo.GetHandler)
	api.GET("/todos/:id", todo.GetByIDHandler)
	api.POST("/todos", todo.CreateHandler)
	api.PUT("/todos/:id", todo.UpdateHandler)
	api.DELETE("/todos/:id", todo.DeleteByIDHandler)
	router.Run(":1234") //listen and serve on 0.0.0.0:1234

}