package db_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/joaocsv/hexagonal-example/adapters/db"
	"github.com/joaocsv/hexagonal-example/app"
	"github.com/stretchr/testify/require"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")

	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	table := `CREATE TABLE products (
		"id" string,
		"name" string,
		"price" float,
		"status" string
	);`

	stmt, err := db.Prepare(table)

	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func createProduct(db *sql.DB) {
	insert := `INSERT INTO products VALUES ("abc", "Product test", 2.00, "disabled")`

	stmt, err := db.Prepare(insert)

	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func TestProductDb_Get(t *testing.T) {
	setUp()
	defer Db.Close()

	productDb := db.NewProductDb(Db)

	product, err := productDb.Get("abc")

	require.Nil(t, err)

	require.Equal(t, "Product test", product.GetName())
	require.Equal(t, 2.00, product.GetPrice())
	require.Equal(t, "disabled", product.GetStatus())
}

func TestProductDb_Save(t *testing.T) {
	setUp()

	defer Db.Close()

	productDb := db.NewProductDb(Db)

	product := app.NewProduct()
	product.Name = "Product Test"
	product.Price = 2.0
	product.Status = app.ENABLED

	productResult, err := productDb.Save(product)

	require.Nil(t, err)
	require.Equal(t, productResult.GetName(), product.GetName())
	require.Equal(t, productResult.GetPrice(), product.GetPrice())
	require.Equal(t, productResult.GetStatus(), product.GetStatus())

	product.Status = app.DISABLED

	productResult, err = productDb.Save(product)

	require.Nil(t, err)
	require.Equal(t, productResult.GetName(), product.GetName())
	require.Equal(t, productResult.GetPrice(), product.GetPrice())
	require.Equal(t, productResult.GetStatus(), product.GetStatus())
}
