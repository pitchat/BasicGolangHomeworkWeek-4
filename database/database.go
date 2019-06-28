package database


import (
	"database/sql"
	"log"
	"os"
	_"github.com/lib/pq"
)

//dbURL for global database connection string 
var dbURL string

//InitDB initial database table todo
func InitDB(){
	dbURL = os.Getenv("DATABASE_URL")
	if len(dbURL) == 0{
		log.Fatal("Environment variable DATABASE_URL is empty")
	}
	db, err := sql.Open("postgres",dbURL)
	if err != nil {
		log.Fatal("Can't connect db",err.Error())
	}
	defer db.Close()

	_, err = db.Exec("DROP TABLE todos")
	if err != nil {
		log.Fatal("Can't drop table fatal error",err.Error())
	}

	createTb := `
	CREATE TABLE IF NOT EXISTS todos(
			id SERIAL PRIMARY KEY,
			title TEXT,
			status TEXT
	);
	`;

	_, err = db.Exec(createTb)
	if err != nil {
		log.Fatal("Can't create table fatal error",err.Error())
	}
	
}

// Connect open connection to database
func Connect()(*sql.DB, error){
	db, err := sql.Open("postgres",dbURL)
	return db,err
}