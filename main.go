package main

import (
	"database/sql"

	dbAdapter "github.com/CelsoTaliatelli/go-hexagonal/adapters/db"
	"github.com/CelsoTaliatelli/go-hexagonal/application"
)

func main() {
	db, _ := sql.Open("sqlite3", "db.sqlite")
	productDbAdapter := dbAdapter.NewProductDb(db)
	productService := application.NewProductService(productDbAdapter)
	productService.Create("Product example", 30)

}
