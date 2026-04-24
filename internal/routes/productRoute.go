package routes

import (
  "net/http"
  "simple-clothes-store/internal/handlers"
)

func productRoute(mux *http.ServeMux, handler *handlers.Handler) {
  mux.HandleFunc("/product", handler.ProductHandler())
}