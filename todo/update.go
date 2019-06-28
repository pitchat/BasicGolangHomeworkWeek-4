package todo

import (
	"strconv"
	"net/http"
	"github.com/gin-gonic/gin"
	"database/sql"
	"github.com/pitchat/test4/database"
)

//Update todo
func (todo Todo)Update(conn *sql.DB) error{
	
	stmt, err := conn.Prepare("UPDATE todos SET title=$2, status=$3 WHERE id=$1;")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(todo.ID, todo.Title, todo.Status)

	return err
}

//UpdateHandler gin api
func UpdateHandler(c *gin.Context) {
	t := Todo{}
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}
	t.ID = id

	err = database.Update(t)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusOK, t)
}