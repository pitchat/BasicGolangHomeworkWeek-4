package todo

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/pitchat/BasicGolangHomeworkWeek-4/database"
	"net/http"
	"strconv"
)

//Delete todo
func (todo Todo) Delete(conn *sql.DB) error {

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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	t.ID = id

	err = database.Delete(t)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
