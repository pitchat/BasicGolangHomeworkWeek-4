package todo

import (
	"strconv"
	"net/http"
	"github.com/gin-gonic/gin"
	"database/sql"
	"github.com/pitchat/test4/database"
)

//GetAll get all todos
func (todo Todo)GetAll(conn *sql.DB) ([]database.DataLayer,error){
	
	tt := []database.DataLayer{}
	
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
		tt = append(tt,database.IConv(t))
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

//GetByKey get todo by key
func (todo Todo)GetByKey(conn *sql.DB) (database.DataLayer,error){

	row := conn.QueryRow("SELECT id, title, status FROM todos where id = $1", todo.ID)
	err := row.Scan(&todo.ID, &todo.Title, &todo.Status)
	if err != nil {
		return todo, err
	}
	return database.IConv(todo), err

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

	t2, err := database.GetByKey(t)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}	
	c.JSON(http.StatusOK, t2)
}