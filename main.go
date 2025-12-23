package main

import (
	"main/cmd"
)

// func handleCors(w http.ResponseWriter) {
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS")
// 	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
// 	w.Header().Set("Content-Type", "application/json")
// }
// func handlePreflightReq(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "OPTIONS" {
// 		w.WriteHeader(200)
// 		return
// 	}
// }

func main() {
	cmd.Serve()
}

// func corsMiddleware(next http.Handler) http.HandlerFunc {
// 	handleCors := func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Access-Control-Allow-Origin", "*")
// 		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS")
// 		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
// 		w.Header().Set("Content-Type", "application/json")

//			next.ServeHTTP(w, r)
//		}
//		handler := http.HandlerFunc(handleCors)
//		return handler
//	}
