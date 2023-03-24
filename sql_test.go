package godatabase

import (
	"context"
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
		var name, email string
		var balance int32
		var rating float64
		var createdAt, birthDate time.Time
		var married bool

		err := rows.Scan(&idcustomer, &name, &email, &balance, &rating, &createdAt, &birthDate, &married)
		if err != nil {
			panic(err)
		}

		fmt.Println("\nId: ", idcustomer, "\nName: ", name, "\nEmail: ", email, "\nBalance: ", balance, "\nRating: ", rating, "\nDate Created: ", createdAt, "\nBirth Date: ", birthDate, "\nMarried Status: ", married)
	}

}
