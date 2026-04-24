package routes

import (
	"net/http"

	"simple-clothes-store/internal/handlers"
	"simple-clothes-store/internal/middleware"
)

func productRoutes(mux *http.ServeMux, handler *handlers.Handler, jwtSecret string) {
	mux.HandleFunc("GET /products", handler.GetAllProductsHandler())
	mux.HandleFunc("GET /products/", handler.GetProductHandler())

	adminAuth := middleware.AdminAuth(jwtSecret)

	mux.Handle("POST /products", adminAuth(handler.CreateProductHandler()))
	mux.Handle("PUT /products/", adminAuth(handler.UpdateProductHandler()))
	mux.Handle("DELETE /products/", adminAuth(handler.DeleteProductHandler()))
}
