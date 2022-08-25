package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Person struct {
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
}

func main() {
	db, _ := sqlx.Open("mysql", "root:king@tcp(127.0.0.1:3306)/learn")
	defer db.Close()

	db.Query("CREATE TABLE PERSON (first_name varchar(50), last_name varchar(50))")

	db.Query("INSERT INTO PERSON(first_name, last_name) values('Afzal', 'Ansari'),('Praveen','Yadav')")

	table_rec := []Person{}
	db.Select(&table_rec, "SELECT * FROM PERSON")

	fmt.Println(table_rec)

}
