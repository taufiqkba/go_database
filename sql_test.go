package godatabase

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
)

func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "INSERT INTO customer(idcustomer, name) VALUES(3, 'bayu')"
	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	script := "SELECT idcustomer, name FROM customer"
	rows, err := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var idcustomer int
		var name string
		rows.Scan(&idcustomer, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("Id: ", idcustomer)
		fmt.Println("Name: ", name)
	}
}

func TestQuerySQLComplex(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT idcustomer, name, email, balance, rating, created_at, birth_date, married FROM customer"
	rows, err := db.QueryContext(ctx, script)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	for rows.Next() {
		var idcustomer int32
		var name string
		var email sql.NullString
		var balance int32
		var rating float64
		var createdAt time.Time
		var birthDate sql.NullTime
		var married bool

		err := rows.Scan(&idcustomer, &name, &email, &balance, &rating, &createdAt, &birthDate, &married)
		if err != nil {
			panic(err)
		}
		fmt.Println("Id:", idcustomer)
		fmt.Println("Name:", name)
		// to hide NULL value
		if email.Valid {
			fmt.Println("Email: ", email.String)
		}
		fmt.Println("Balance:", balance)
		fmt.Println("Rating:", rating)
		fmt.Println("Date Created:", createdAt)
		// to hide NULL value
		if birthDate.Valid {
			fmt.Println("Birth Date: ", birthDate.Time)
		}
		fmt.Println("Married Status:", married)
		// fmt.Println("\nId: ", idcustomer, "\nName: ", name, "\nEmail: ", email, "\nBalance: ", balance, "\nRating: ", rating, "\nDate Created: ", createdAt, "\nBirth Date: ", birthDate, "\nMarried Status: ", married)
	}

}

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	// case sql injection dibobol
	username := "admin'; #"
	password := "salah"

	script := "SELECT username FROM user WHERE username = '" + username + "' AND password = '" + password + "' LIMIT 1"
	rows, err := db.QueryContext(ctx, script)
	fmt.Println(script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Success login,", username)
	} else {
		fmt.Println("Failed Login!")
	}
}

func TestSqlInjectionSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin"
	password := "admin"

	script := "SELECT username FROM user WHERE username = ? AND password = ?  LIMIT 1" // sql query with parameter
	rows, err := db.QueryContext(ctx, script, username, password)
	// fmt.Println(script)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Success login,", username)
	} else {
		fmt.Println("Failed Login!")
	}
}
