package product

import (
	"encoding/json"
	"main/repo"
	"main/utils"
	"net/http"
	"strconv"
)

type RewUpdateProduct struct {
	Title       string
	Description string
	Price       float64
	ImgURl      string
}

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {

	productID := r.PathValue("id")
	productIDInt, err := strconv.Atoi(productID)
	if err != nil {
		utils.SendError(w, "please give a valid id", http.StatusBadRequest)
		return
	}

	var req RewUpdateProduct
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&req)
	if err != nil {
		utils.SendError(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	updatedProduct, err := h.productRepo.UpdateProduct(productIDInt, repo.Product{
		Title:       req.Title,
		Description: req.Description,
		Price:       req.Price,
		ImgURl:      req.ImgURl,
	})
	if err != nil {
		utils.SendError(w, "product not found", http.StatusNotFound)
		return
	}
	utils.SendData(w, updatedProduct, http.StatusOK)
}
