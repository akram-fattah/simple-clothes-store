package routes

import (
	"net/http"

	"simple-clothes-store/internal/handlers"
)

func SetupRoutes(mux *http.ServeMux, handler *handlers.Handler, jwtSecret string) {
	productRoutes(mux, handler, jwtSecret)
	authRoutes(mux, handler)
}

func authRoutes(mux *http.ServeMux, handler *handlers.Handler) {
	mux.HandleFunc("POST /admin/login", handler.LoginHandler())
}
