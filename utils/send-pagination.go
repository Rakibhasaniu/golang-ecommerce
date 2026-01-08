package utils

import (
	"net/http"
)

type PaginatedData struct {
	Data       any        `json:"data"`
	Pagination Pagination `json:"pagination"`
}
type Pagination struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
}

func SendPagination(w http.ResponseWriter, data any, page, limit, total int) {
	paginatedData := PaginatedData{
		Data: data,
		Pagination: Pagination{
			Page:       page,
			Limit:      limit,
			Total:      total,
			TotalPages: total / int(limit),
		},
	}
	SendData(w, paginatedData, http.StatusOK)
}
