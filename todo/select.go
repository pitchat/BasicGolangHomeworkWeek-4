package todo

import (
	"strconv"
	"net/http"
	"github.com/gin-gonic/gin"
	"database/sql"
	"github.com/pitchat/test4/database"
)

//GetAll get all todos
func (todo Todo)GetAll(conn *sql.DB) ([]Todo,error){
	
	tt := []Todo{}
	
	query := "SELECT id, title, status FROM todos"
	rows, err := conn.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var t Todo
		if err := rows.Scan(&t.ID, &t.Title, &t.Status); err != nil {
			return nil, err
		}
		tt = append(tt,t)
	}

	return tt, err
}

//GetHandler gin api
func GetHandler(c *gin.Context) {
	t := Todo{}
	conn, err := database.Connect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}
	defer conn.Close()

	tt, err := t.GetAll(conn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusOK, tt)

}

//GetByID get todo by id
func (todo Todo)GetByID(conn *sql.DB) (Todo,error){

	row := conn.QueryRow("SELECT id, title, status FROM todos where id = $1", todo.ID)
	err := row.Scan(&todo.ID, &todo.Title, &todo.Status)
	if err != nil {
		return todo, err
	}
	return todo, err

}

//GetByIDHandler for retrive Todo by ID
func GetByIDHandler(c *gin.Context) {
	t := Todo{}
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}
	t.ID = id

	conn, err := database.Connect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}
	defer conn.Close()
	
	t, err = t.GetByID(conn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, t)
}