package main

import (
	"fmt"
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/pitchat/BasicGolangHomeworkWeek-4/database"
	"github.com/pitchat/BasicGolangHomeworkWeek-4/todo"
)

func TestHandler(t *testing.T) {
	database.InitDB()
	router := setupRouter()
	fmt.Println("Begin test")
	//list := []todo.Todo{}
	todo := todo.Todo{}

	
	
	//Case 2 Get All
	c2 := httptest.NewRecorder()
	req2, _ := http.NewRequest("GET", "/api/todos/1", nil)
	router.ServeHTTP(c2, req2)
	if c2.Code !=  200 {
		t.Errorf("http response code is incorrect(200). result value: %d.", c2.Code)
	}

	json.Unmarshal([]byte(c2.Body.String()), &todo)
	if todo.ID != int64(1) {
		t.Errorf("todo ID is incorrect. expect value: %d, result value: %d.", 1, todo.ID)
	}
}

func subTestCreate(t *testing.T, router *gin.Engine){
	todo := todo.Todo{}
	//Case 1 Create
	c1 := httptest.NewRecorder()
	req1, _ := http.NewRequest("POST", "/api/todos", bytes.NewBufferString(`{"title":"walk dog","status": "active"}`))
	router.ServeHTTP(c1, req1)
	if !(c1.Code ==  201 ||  c1.Code ==  202) {
		t.Errorf("create todo http response code is incorrect(201,202). result value: %d.", c1.Code)
	}
	json.Unmarshal([]byte(c1.Body.String()), &todo)
	if todo.ID ==  0 {
		t.Errorf("create todo ID is zero.")
	}
	if todo.Title !=  "walk dog" {
		t.Errorf("todo Title is incorrect. expect value: %q, result value: %q.", "walk dog", todo.Title)
	}
	if todo.Status !=  "active" {
		t.Errorf("todo Status is incorrect. expect value: %q, result value: %q.", "active", todo.Status)
	}
}

