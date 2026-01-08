package product

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"main/utils"
)

var cnt int

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	reqQuery := r.URL.Query()
	pageAsString := reqQuery.Get("page")
	limitAsString := reqQuery.Get("limit")
	page, _ := strconv.ParseInt(pageAsString, 10, 32)
	limit, _ := strconv.ParseInt(limitAsString, 10, 32)
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	products, err := h.svc.GetProducts(int(page), int(limit))
	if err != nil {
		fmt.Println(err)
		utils.SendError(w, "Failed to get products", http.StatusInternalServerError)
		return
	}
	// total, err := h.svc.CountProducts()
	// if err != nil {
	// 	fmt.Println(err)
	// 	utils.SendError(w, "Failed to get products", http.StatusInternalServerError)
	// 	return
	// }
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		total, err := h.svc.CountProducts()
		if err != nil {
			fmt.Println(err)
			utils.SendError(w, "Failed to get products", http.StatusInternalServerError)
			return
		}
		cnt = total

	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		total, err := h.svc.CountProducts()
		if err != nil {
			fmt.Println(err)
			utils.SendError(w, "Failed to get products", http.StatusInternalServerError)
			return
		}
		cnt = total
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		total, err := h.svc.CountProducts()
		if err != nil {
			fmt.Println(err)
			utils.SendError(w, "Failed to get products", http.StatusInternalServerError)
			return
		}
		cnt = total

	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		total, err := h.svc.CountProducts()
		if err != nil {
			fmt.Println(err)
			utils.SendError(w, "Failed to get products", http.StatusInternalServerError)
			return
		}
		cnt = total

	}()
	wg.Wait()
	utils.SendPagination(w, products, int(page), int(limit), cnt)

}
