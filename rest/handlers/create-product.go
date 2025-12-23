package handler

import (
	"encoding/json"
	"net/http"

	"main/database"
	"main/utils"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {

	var product database.Product

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&product)
	if err != nil {
		utils.SendError(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}
	product.ID = len(database.List()) + 1
	database.Store(product)
	utils.SendData(w, product, http.StatusCreated)

}
