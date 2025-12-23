package handler

import (
	"net/http"
	"strconv"

	"main/database"
	"main/utils"
)

func GetProductById(w http.ResponseWriter, r *http.Request) {
	productzid := r.PathValue("id")
	pid, err := strconv.Atoi(productzid)
	if err != nil {
		utils.SendError(w, "please give a valid id", http.StatusBadRequest)
		return
	}
	product := database.ListById(pid)
	if product == nil {
		utils.SendError(w, "product not found", http.StatusNotFound)
		return
	}
	utils.SendData(w, product, http.StatusOK)

}
