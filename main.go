package main

import (
	"fmt"
	"log"
	"net/http"

	"simple-clothes-store/configserver"
	"simple-clothes-store/internal/handlers"
	"simple-clothes-store/internal/models"
	"simple-clothes-store/internal/repository"
	"simple-clothes-store/internal/routes"
	"simple-clothes-store/internal/utils"
)

func main() {
	config, err := configserver.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configs: %v", err)
	}

	db, err := repository.NewDB(config.DSN())
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()
	fmt.Println("Connected to PostgreSQL successfully")

	productRepo := repository.NewProductRepository(db)
	adminRepo := repository.NewAdminRepository(db)

	if err := seedDefaultAdmin(adminRepo); err != nil {
		log.Printf("Warning: failed to seed default admin: %v", err)
	}

	handler := handlers.NewHandler(productRepo, adminRepo, config.JWTSecret)

	mux := http.NewServeMux()
	routes.SetupRoutes(mux, handler, config.JWTSecret)

	serverAddress := fmt.Sprintf(":%s", config.ServerPort)
	server := &http.Server{
		Addr:    serverAddress,
		Handler: mux,
	}
	fmt.Printf("Server starting on port %s\n", serverAddress)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func seedDefaultAdmin(repo repository.AdminRepository) error {
	existing, err := repo.GetByUsername("admin")
	if err != nil {
		return err
	}
	if existing != nil {
		return nil 
	}

	hashedPassword, err := utils.HashPassword("admin123")
	if err != nil {
		return err
	}

	return repo.Create(&models.Admin{
		Username: "admin",
		Password: hashedPassword,
	})
}
