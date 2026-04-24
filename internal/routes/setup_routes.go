package routes

import (
  "net/http"
  "simple-clothes-store/internal/handlers"
)

func SetupRoutes(mux *http.ServeMux, handler *handlers.Handler) {
  productRoute(mux, handler)
}