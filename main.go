package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/pitchat/BasicGolangHomeworkWeek-4/database"
	"github.com/pitchat/BasicGolangHomeworkWeek-4/todo"
)

func main() {
	database.InitDB()
	r := setupRouter()
	r.Run(":1234") //listen and serve on 0.0.0.0:1234
}

func setupRouter() *gin.Engine{
	r := gin.Default()

	r.Use(authMiddleware)

	api := r.Group("/api")
	api.GET("/todos", todo.GetHandler)
	api.GET("/todos/:id", todo.GetByIDHandler)
	api.POST("/todos", todo.CreateHandler)
	api.PUT("/todos/:id", todo.UpdateHandler)
	api.DELETE("/todos/:id", todo.DeleteByIDHandler)	
	return r
}

func authMiddleware(c *gin.Context){
	token := c.GetHeader("Authorization")
	fmt.Println("token:",token)
   if token != "Bearer token123"{
	   c.JSON(http.StatusUnauthorized, gin.H{"error": http.StatusText(http.StatusUnauthorized)})
	   c.Abort()
	   return
   }
   c.Next()
}
