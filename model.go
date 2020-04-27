package main

import (
	"database/sql"
)

type product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// get requests
func (prod *product) getProduct(db *sql.DB) error {
	return db.QueryRow("SELECT name, price FROM products WHERE id=$1", prod.ID).
		Scan(&prod.Name, &prod.Price)
}

func getProducts(db *sql.DB, start, count int) ([]product, error) {
	rows, err := db.Query(
		"SELECT id, name, price FROM products LIMIT $1 OFFSET $2", count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	products := []product{}

	for rows.Next() {
		var prod product
		if err := rows.Scan(&prod.ID, &prod.Name, &prod.Price); err != nil {
			return nil, err
		}
		products = append(products, prod)
	}

	return products, nil
}

// post requests
func (prod *product) createProduct(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO products(name, price) VALUES($1, $2) RETURNING id", prod.Name, prod.Price).Scan(&prod.ID)

	if err != nil {
		return err
	}

	return nil
}

// put requests
func (prod *product) updateProduct(db *sql.DB) error {
	_, err := db.Exec("UPDATE products SET name=$1, price=$2 WHERE id=$3", prod.Name, prod.Price, prod.ID)
	return err
}

// delete requests
func (prod *product) deleteProduct(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM products WHERE id=$1", prod.ID)
	return err
}
