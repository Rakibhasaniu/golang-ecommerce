package handler

import (
	"encoding/json"
	"main/database"
	"main/utils"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user database.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		utils.SendError(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}
	user = user.Store()
	utils.SendData(w, user, http.StatusCreated)
}
