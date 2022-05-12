package core

import (
	"context"
	"database/sql"
	"log"
)

func InsertProduct(db *sql.DB) {
	insertquery := `INSERT INTO "product"("product_id","product_name","product_price","description") values(101,'soap',10,'hello')`
	_, err := db.ExecContext(context.TODO(), insertquery)
	if err != nil {
		log.Fatalf("error inserting product: %v", err)
	}

}
