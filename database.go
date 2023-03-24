package godatabase

import (
	"database/sql"
	"time"
)

func GetConnection() *sql.DB {
	// parseTime = true untuk konversi untuk time data time.Time dari []uint8 menjadi time.Time
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/go_database?parseTime=true")
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
