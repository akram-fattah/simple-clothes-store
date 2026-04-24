package handlers

import (
	"simple-clothes-store/internal/repository"
)


type Handler struct {
	ProductRepo repository.ProductRepository
	AdminRepo   repository.AdminRepository
	JWTSecret   string
}

func NewHandler(productRepo repository.ProductRepository, adminRepo repository.AdminRepository, jwtSecret string) *Handler {
	return &Handler{
		ProductRepo: productRepo,
		AdminRepo:   adminRepo,
		JWTSecret:   jwtSecret,
	}
}
