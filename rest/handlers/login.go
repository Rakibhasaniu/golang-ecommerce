package handler

import (
	"encoding/json"
	"main/database"
	"main/utils"
	"net/http"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest LoginRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&loginRequest)
	if err != nil {
		utils.SendError(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}
	user := database.FindByEmail(loginRequest.Email, loginRequest.Password)
	if user == nil {
		utils.SendError(w, "User not found", http.StatusNotFound)
		return
	}
	utils.SendData(w, user, http.StatusOK)
}
