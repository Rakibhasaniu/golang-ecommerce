package user

import (
	"encoding/json"
	"main/utils"
	"net/http"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest LoginRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&loginRequest)
	if err != nil {
		utils.SendError(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}
	user, err := h.svc.GetUserByEmail(loginRequest.Email, loginRequest.Password)
	if err != nil {
		utils.SendError(w, "User not found", http.StatusNotFound)
		return
	}
	cnf := h.config
	accessToken, err := utils.CreateJWT(cnf.JwtSecret, utils.Payload{
		Sub:         user.ID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		IsShopPwner: user.IsShopOwner,
	})
	if err != nil {
		utils.SendError(w, "Failed to create JWT", http.StatusInternalServerError)
		return
	}
	utils.SendData(w, map[string]interface{}{
		"jwt": accessToken,
	}, http.StatusOK)
}
