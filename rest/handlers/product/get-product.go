package product

import (
	"fmt"
	"net/http"

	"main/utils"
)

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {

	products, err := h.productRepo.GetProducts()
	if err != nil {
		fmt.Println(err)
		utils.SendError(w, "Failed to get products", http.StatusInternalServerError)
		return
	}
	utils.SendData(w, products, http.StatusOK)

}
