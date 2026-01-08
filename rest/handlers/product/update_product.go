package product

import (
	"encoding/json"
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
	pid, err := strconv.Atoi(productID)
	if err != nil {
		utils.SendError(w, "please give a valid id", http.StatusBadRequest)
		return
	}

	// Get existing product
	existingProduct, err := h.svc.GetProductById(pid)
	if err != nil {
		utils.SendError(w, "internal server error", http.StatusInternalServerError)
		return
	}
	if existingProduct == nil {
		utils.SendError(w, "product not found", http.StatusNotFound)
		return
	}

	var req RewUpdateProduct
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&req)
	if err != nil {
		utils.SendError(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}

	// Merge changes with existing product
	if req.Title != "" {
		existingProduct.Title = req.Title
	}
	if req.Description != "" {
		existingProduct.Description = req.Description
	}
	if req.Price != 0 {
		existingProduct.Price = req.Price
	}
	if req.ImgURl != "" {
		existingProduct.ImgUrl = req.ImgURl
	}

	updatedProduct, err := h.svc.UpdateProduct(*existingProduct)
	if err != nil {
		utils.SendError(w, "internal server error", http.StatusInternalServerError)
		return
	}
	if updatedProduct == nil {
		utils.SendError(w, "product not found", http.StatusNotFound)
		return
	}
	utils.SendData(w, updatedProduct, http.StatusOK)
}
