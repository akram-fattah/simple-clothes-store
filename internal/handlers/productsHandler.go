package handlers

import (
  "net/http"
  "encoding/json"
)

func (h* Handler) ProductHandler() http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    response := map[string]string {
      "message": "Ok",
    }
    json.NewEncoder(w).Encode(response)
  }
}