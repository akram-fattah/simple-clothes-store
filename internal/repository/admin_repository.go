package repository

import (
	"database/sql"
	"fmt"

	"simple-clothes-store/internal/models"
)

type AdminRepository interface {
	GetByUsername(username string) (*models.Admin, error)
	Create(admin *models.Admin) error
}

type adminRepo struct {
	db *sql.DB
}

func NewAdminRepository(db *sql.DB) AdminRepository {
	return &adminRepo{db: db}
}

func (r *adminRepo) GetByUsername(username string) (*models.Admin, error) {
	var a models.Admin
	err := r.db.QueryRow(
		`SELECT id, username, password FROM admins WHERE username = $1`,
		username,
	).Scan(&a.ID, &a.Username, &a.Password)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get admin: %w", err)
	}
	return &a, nil
}

func (r *adminRepo) Create(admin *models.Admin) error {
	return r.db.QueryRow(
		`INSERT INTO admins (username, password) VALUES ($1, $2) RETURNING id`,
		admin.Username, admin.Password,
	).Scan(&admin.ID)
}
