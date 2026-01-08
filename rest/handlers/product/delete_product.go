package product

import (
	"main/utils"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {

	productID := r.PathValue("id")
	productIDInt, err := strconv.Atoi(productID)
	if err != nil {
		utils.SendError(w, "please give a valid id", http.StatusBadRequest)
		return
	}
	h.svc.DeleteProduct(productIDInt)
	utils.SendData(w, "product deleted", http.StatusOK)

}
