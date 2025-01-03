package dbinit

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

func init() {
	server := "postgres"  
	port := 5432          
	user := "root"
	password := "root"
	dbname := "validate"

	ConnectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", server, port, user, password, dbname)
	fmt.Println("Connecting with:", ConnectionString)

	for retries := 0; retries < 10; retries++ {
		DB, err = sql.Open("postgres", ConnectionString)
		if err != nil {
			log.Printf("Attempt %d: Error creating connection pool: %v", retries+1, err)
		} else {
			err = DB.Ping()
			if err == nil {
				log.Println("Database connected successfully!")
				return 
			}
			log.Printf("Attempt %d: Error connecting to database: %v", retries+1, err)
		}

		time.Sleep(2 * time.Second)
	}

	panic("Failed to connect to database after 10 attempts")
}
