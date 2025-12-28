package product

import (
	"encoding/json"
	"net/http"

	"main/repo"
	"main/utils"
)

type RewCreateProduct struct {
	Title       string
	Description string
	Price       float64
	ImgURl      string
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	var req RewCreateProduct

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		utils.SendError(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}
	createdProduct, err := h.productRepo.CreateProduct(repo.Product{
		Title:       req.Title,
		Description: req.Description,
		Price:       req.Price,
		ImgURl:      req.ImgURl,
	})
	if err != nil {
		utils.SendError(w, "Failed to create product", http.StatusInternalServerError)
		return
	}
	utils.SendData(w, createdProduct, http.StatusCreated)

}
