package handler

import (
	"net/http"

	"main/database"
	"main/utils"
)

func GetProduct(w http.ResponseWriter, r *http.Request) {

	utils.SendData(w, database.List(), http.StatusOK)

}
