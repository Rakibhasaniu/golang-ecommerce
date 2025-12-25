package product

import (
	"encoding/json"
	"net/http"

	"main/database"
	"main/utils"
)

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	var product database.Product

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&product)
	if err != nil {
		utils.SendError(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}
	product = database.Store(product)

	utils.SendData(w, product, http.StatusCreated)

}
