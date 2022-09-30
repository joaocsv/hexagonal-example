package db

import (
	"database/sql"

	"github.com/joaocsv/hexagonal-example/app"
	_ "github.com/mattn/go-sqlite3"
)

type ProductDb struct {
	db *sql.DB
}

func NewProductDb(db *sql.DB) *ProductDb {
	return &ProductDb{db: db}
}

func (p *ProductDb) Get(id string) (app.ProductInterface, error) {
	var product app.Product

	stmt, err := p.db.Prepare("SELECT id, name, price, status FROM products where id=?")

	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price, &product.Status)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *ProductDb) Save(product app.ProductInterface) (app.ProductInterface, error) {
	var rows int

	stmt, err := p.db.Prepare("SELECT id FROM products WHERE id = ?")

	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(product.GetID()).Scan(&rows)

	if err != nil {
		return nil, err
	}

	if rows == 0 {
		_, err = p.create(product)

		if err != nil {
			return nil, err
		}
	} else {
		_, err = p.update(product)

		if err != nil {
			return nil, err
		}
	}

	return product, nil
}

func (p *ProductDb) create(product app.ProductInterface) (app.ProductInterface, error) {
	stmt, err := p.db.Prepare("INSERT INTO products (id, name, price, status) VALUES (?,?,?,?)")

	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProductDb) update(product app.ProductInterface) (app.ProductInterface, error) {
	stmt, err := p.db.Prepare("UPDATE products SET name = ?, price = ?, status = ? WHERE id = ?")

	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(product.GetName(), product.GetPrice(), product.GetStatus(), product.GetID())

	if err != nil {
		return nil, err
	}

	return product, nil
}
