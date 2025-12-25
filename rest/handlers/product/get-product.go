package product

import (
	"net/http"

	"main/database"
	"main/utils"
)

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {

	utils.SendData(w, database.List(), http.StatusOK)

}
