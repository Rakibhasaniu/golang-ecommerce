package product

import (
	"net/http"

	"main/utils"
)

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {

	products, err := h.productRepo.GetProducts()
	if err != nil {
		utils.SendError(w, "Failed to get products", http.StatusInternalServerError)
		return
	}
	utils.SendData(w, products, http.StatusOK)

}
