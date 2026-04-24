package handlers

import (
	"encoding/json"
	"net/http"

	"simple-clothes-store/internal/utils"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *Handler) LoginHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(map[string]string{"error": "method not allowed"})
			return
		}

		var req LoginRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "invalid json"})
			return
		}

		admin, err := h.AdminRepo.GetByUsername(req.Username)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "server error"})
			return
		}
		if admin == nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"error": "invalid credentials"})
			return
		}

		if err := utils.CheckPassword(req.Password, admin.Password); err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"error": "invalid credentials"})
			return
		}

		token, err := utils.GenerateToken(admin.Username, h.JWTSecret)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "failed to generate token"})
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "login successful",
			"token":   token,
		})
	}
}
