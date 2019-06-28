package todo

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"database/sql"
	"github.com/pitchat/test4/database"
)

//Insert todo
func (todo Todo)Insert(conn *sql.DB) (database.DataLayer,error){
	
	query := `
	INSERT INTO todos (title, status) 
	VALUES 	($1, $2) RETURNING id
	`;

	row := conn.QueryRow(query,todo.Title,todo.Status)
	err := row.Scan(&todo.ID)
	return database.IConv(todo), err
}

//CreateHandler gin api
func CreateHandler(c *gin.Context) {
	t := Todo{}
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	t2, err := database.Insert(t)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusCreated, t2)

}