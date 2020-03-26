package store

import (
	"context"
	"lab/productLab/internal/entity"
)

// Create store a new product
func (s *Store) Create(ctx context.Context, product *entity.Product) (*entity.Product, error) {
	err := s.client.QueryRow(
		"INSERT INTO products(name, price, comments, timestamp) VALUES($1, $2, $3, $4) RETURNING id;",
		product.Name, product.Price, product.Comments, product.Timestamp).Scan(&product.ID)
	if err != nil {
		return product, err
	}

	return product, nil
}

// ByID get by id
func (s *Store) ByID(ctx context.Context, id int) (*entity.Product, error) {
	var product entity.Product
	err := s.client.QueryRow("SELECT id, name, price, comments, timestamp FROM products WHERE id=$1",
		id).Scan(&product.ID, &product.Name, &product.Price, &product.Comments, &product.Timestamp)
	if err != nil {
		return &product, err
	}
	return &product, err
}

// Update updates product atributes
func (s *Store) Update(ctx context.Context, id int, newProduct *entity.Product) error {
	_, err := s.client.Exec("UPDATE products SET name=$1, price=$2, comments=$3, timestamp=$4 WHERE id=$5", newProduct.Name, newProduct.Price, newProduct.Comments, newProduct.Timestamp, id)
	return err
}

// Delete deletes by id
func (s *Store) Delete(ctx context.Context, id int) error {
	_, err := s.client.Exec("DELETE FROM products WHERE id=$1", id)
	return err
}

// List gets all products
func (s *Store) List(ctx context.Context, start, count int) ([]*entity.Product, error) {
	rows, err := s.client.Query(
		"SELECT id, name,  price, comments FROM products LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	products := []*entity.Product{}

	for rows.Next() {
		var p entity.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Timestamp); err != nil {
			return nil, err
		}
		products = append(products, &p)
	}
	return products, nil
}
