package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	core "github.com/capslock-inc/gprc-demo/Servers/postgresclient/Core"
	_ "github.com/lib/pq"
)

// althaf

const (
	host     = "localhost"
	port     = 5434
	user     = "admin"
	password = "mypassword"
	dbname   = "ecom"
)

var connstring string = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

func PostgresINIT() *sql.DB {
	psql, err := sql.Open("postgres", connstring)
	if err != nil {
		log.Fatalf("error connecting postgres: %v", err)
	} else {
		log.Printf("connection established")
	}
	return psql
}
func main() {
	dbconn := PostgresINIT()

	// create table if not exist
	createtablequery := `CREATE TABLE IF NOT EXISTS product(product_id int primary key, product_name text,  product_price int,description text);`
	res, err := dbconn.ExecContext(context.TODO(), createtablequery)
	if err != nil {
		log.Fatalf("could not able to create table: %v", err)
	} else {
		log.Printf("table created sucessfully :%v", res)
	}
	core.InsertProduct(dbconn)
	defer dbconn.Close()

}
