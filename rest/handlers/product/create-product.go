package product

import (
	"encoding/json"
	"net/http"

	"main/domain"
	"main/utils"
)

type RewCreateProduct struct {
	Title       string
	Description string
	Price       float64
	ImgUrl      string
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {

	var req RewCreateProduct

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		utils.SendError(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}
	createdProduct, err := h.svc.CreateProduct(domain.Product{
		Title:       req.Title,
		Description: req.Description,
		Price:       req.Price,
		ImgUrl:      req.ImgUrl,
	})
	if err != nil {
		utils.SendError(w, "Failed to create product", http.StatusInternalServerError)
		return
	}
	utils.SendData(w, createdProduct, http.StatusCreated)

}
