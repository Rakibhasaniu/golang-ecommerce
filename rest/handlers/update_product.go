package handler

import (
	"encoding/json"
	"main/database"
	"main/utils"
	"net/http"
	"strconv"
)

func UpdateProduct(w http.ResponseWriter, r *http.Request) {

	productID := r.PathValue("id")
	productIDInt, err := strconv.Atoi(productID)
	if err != nil {
		utils.SendError(w, "please give a valid id", http.StatusBadRequest)
		return
	}

	var product database.Product
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&product)
	if err != nil {
		utils.SendError(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	product.ID = productIDInt
	updatedProduct := database.Update(productIDInt, product)
	if updatedProduct == nil {
		utils.SendError(w, "product not found", http.StatusNotFound)
		return
	}
	utils.SendData(w, updatedProduct, http.StatusOK)
}
