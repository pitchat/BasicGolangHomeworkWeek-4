package todo

import (
	"strconv"
	"net/http"
	"github.com/gin-gonic/gin"
	"database/sql"
	"github.com/pitchat/test4/database"
)

//Delete todo
func (todo Todo)Delete(conn *sql.DB) error{
	
	stmt, err := conn.Prepare("DELETE FROM todos WHERE id=$1;")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(todo.ID)

	return err
}

//DeleteByIDHandler gin api
func DeleteByIDHandler(c *gin.Context) {
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

	err = t.Delete(conn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}