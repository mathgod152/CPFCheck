package dbinit

import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
)

var (
	DB *sql.DB
	err error
)

func init() {
	server := "localhost"
	port := 5430
	user := "root"
	password := "root"
	dbname := "validate"

	ConnectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", server, port, user, password, dbname)
	fmt.Println("Connecting with:", ConnectionString)

	DB, err = sql.Open("postgres", ConnectionString)
	if err != nil {
		panic("Error creating connection pool: " + err.Error())
	}

	err = DB.Ping()
	if err != nil {
		panic("Error connecting to database: " + err.Error())
	}

	fmt.Println("Database connected successfully!")
}
