package godatabase

import (
	"database/sql"
	"time"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/go_database")
	if err != nil {
		panic(err)
	}

	// using DB
	db.SetMaxIdleConns(10)  //minimum connection to database
	db.SetMaxOpenConns(100) //maximum connection to database
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)
	return db

}
