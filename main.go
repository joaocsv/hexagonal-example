package main

import (
	"database/sql"

	"github.com/joaocsv/hexagonal-example/adapters/db"
	"github.com/joaocsv/hexagonal-example/app"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	connection, _ := sql.Open("sqlite3", "db.sqlite")
	productDb := db.NewProductDb(connection)
	productService := app.NewProductService(productDb)

	product, _ := productService.Create("Coca-cola 2L", 8.30)

	productService.Enable(product)
}
