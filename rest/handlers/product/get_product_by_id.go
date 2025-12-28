package product

import (
	"net/http"
	"strconv"

	"main/utils"
)

func (h *Handler) GetProductById(w http.ResponseWriter, r *http.Request) {
	productzid := r.PathValue("id")
	pid, err := strconv.Atoi(productzid)
	if err != nil {
		utils.SendError(w, "please give a valid id", http.StatusBadRequest)
		return
	}
	product, err := h.productRepo.GetProductById(pid)
	if err != nil {
		utils.SendError(w, "product not found", http.StatusNotFound)
		return
	}
	utils.SendData(w, product, http.StatusOK)

}
