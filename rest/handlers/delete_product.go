package handler

import (
	"main/database"
	"main/utils"
	"net/http"
	"strconv"
)

func DeleteProduct(w http.ResponseWriter, r *http.Request) {

	productID := r.PathValue("id")
	productIDInt, err := strconv.Atoi(productID)
	if err != nil {
		utils.SendError(w, "please give a valid id", http.StatusBadRequest)
		return
	}
	database.Delete(productIDInt)
	utils.SendData(w, "product deleted", http.StatusOK)

}
