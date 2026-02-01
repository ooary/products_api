package repository

import (
	"database/sql"
	"errors"
	"products_api/internal/models"
)

type ProductRepository struct {
	DB *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}
func (r *ProductRepository) GetAll() ([]models.Products, error) {
	query := "SELECT id,name,description,price,stock from products"
	rows, err := r.DB.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Products

	for rows.Next() {
		var p models.Products
		err := rows.Scan(&p.ID, &p.NAME, &p.DESCRIPTION, &p.PRICE, &p.STOCK)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func (r *ProductRepository) CreateProduct(product *models.Products) error {
	query := "INSERT INTO PRODUCTS (name,description,price,stock) VALUES (?,?,?,?) RETURNING id"
	err := r.DB.QueryRow(query, product.NAME, product.DESCRIPTION, product.PRICE, product.STOCK).Scan(&product.ID)
	return err

}

func (r *ProductRepository) GetProductDetail(id int) (*models.Products, error) {
	query := "SELECT id,name,description,price,stock from products where id = ?"
	var p models.Products
	err := r.DB.QueryRow(query, id).Scan(&p.ID, &p.NAME, &p.DESCRIPTION, &p.PRICE, &p.STOCK)
	errors.Is(err, sql.ErrNoRows)
	return &p, nil
}

func (r *ProductRepository) UpdateProduct(p *models.Products) error {
	query := "UPDATE products SET name = ? ,description =?, price = ? , stock =? where id =?"

	_, err := r.DB.Exec(query, p.NAME, p.DESCRIPTION, p.STOCK, p.PRICE, p.ID)
	return err

}

func (r *ProductRepository) DeleteProduct(id int) bool {
	query := "DELETE FROM PRODUCTS WHERE id = ?"
	res, err := r.DB.Exec(query, id)
	if err != nil {
		return false
	}
	count, _ := res.RowsAffected()
	if count == 0 {
		return false
	}

	return true
}
