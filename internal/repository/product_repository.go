package repository

import (
	"database/sql"
	"fmt"

	"simple-clothes-store/internal/models"
)


type ProductRepository interface {
	GetAll() ([]models.Product, error)
	GetByID(id int) (*models.Product, error)
	Create(p *models.Product) error
	Update(p *models.Product) error
	Delete(id int) error
}

type productRepo struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return &productRepo{db: db}
}

func (r *productRepo) GetAll() ([]models.Product, error) {
	rows, err := r.db.Query(`SELECT id, name, description, price, image FROM products ORDER BY id`)
	if err != nil {
		return nil, fmt.Errorf("failed to query products: %w", err)
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Image); err != nil {
			return nil, fmt.Errorf("failed to scan product: %w", err)
		}
		products = append(products, p)
	}
	return products, rows.Err()
}

func (r *productRepo) GetByID(id int) (*models.Product, error) {
	var p models.Product
	err := r.db.QueryRow(
		`SELECT id, name, description, price, image FROM products WHERE id = $1`,
		id,
	).Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Image)

	if err == sql.ErrNoRows {
		return nil, nil 
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get product: %w", err)
	}
	return &p, nil
}

func (r *productRepo) Create(p *models.Product) error {
	return r.db.QueryRow(
		`INSERT INTO products (name, description, price, image) VALUES ($1, $2, $3, $4) RETURNING id`,
		p.Name, p.Description, p.Price, p.Image,
	).Scan(&p.ID)
}

func (r *productRepo) Update(p *models.Product) error {
	res, err := r.db.Exec(
		`UPDATE products SET name = $1, description = $2, price = $3, image = $4 WHERE id = $5`,
		p.Name, p.Description, p.Price, p.Image, p.ID,
	)
	if err != nil {
		return fmt.Errorf("failed to update product: %w", err)
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check rows affected: %w", err)
	}
	if rows == 0 {
		return fmt.Errorf("product not found")
	}
	return nil
}

func (r *productRepo) Delete(id int) error {
	res, err := r.db.Exec(`DELETE FROM products WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("failed to delete product: %w", err)
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check rows affected: %w", err)
	}
	if rows == 0 {
		return fmt.Errorf("product not found")
	}
	return nil
}
