package repository

import (
	"context"
	"database/sql"
	"errors"
	"exam-inventory/model"
)

type ProductRepository interface {
	FindAll(ctx context.Context) ([]model.Product, error)
	Update(ctx context.Context, tx *sql.Tx, contact model.Product) error
	GetProductByProductName(ctx context.Context, productName string) (model.Product, error)
}

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *productRepository {
	return &productRepository{db}
}

func (repo *productRepository) FindAll(ctx context.Context) ([]model.Product, error) {
	var query string = "SELECT id, name, price, stock, created_at FROM products"
	var products []model.Product

	rows, err := repo.db.QueryContext(ctx, query)
	if err != nil {
		return products, err
	}
	for rows.Next() {
		var product model.Product
		rows.Scan(product.GetId(), product.GetName(), product.GetPrice(), product.GetStock(), product.GetCreatedAt())
		products = append(products, product)
	}
	return products, nil
}

func (repo *productRepository) GetProductByProductName(ctx context.Context, productName string) (model.Product, error) {
	var product model.Product
	var query string = "SELECT id, name, price, stock, created_at FROM products WHERE name = ? "

	rows := repo.db.QueryRowContext(ctx, query, productName)
	err := rows.Scan(product.GetId(), product.GetName(), product.GetPrice(), product.GetStock(), product.GetCreatedAt())
	if err != nil {
		return product, errors.New("nama product tidak ada dalam tabel product")
	}
	return product, nil
}

func (repo *productRepository) Update(ctx context.Context, tx *sql.Tx, product model.Product) error {
	query := "UPDATE products SET stock=? WHERE id=?"

	_, err := tx.ExecContext(ctx, query, product.GetStock(), product.GetId())
	if err != nil {
		return err
	}

	return nil
}
