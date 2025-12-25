package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"main/utils"
	"net/http"
	"strings"
)

type contextKey string

const UserContextKey contextKey = "user"

func (m *Middlewares) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			utils.SendError(w, "Unauthorized: No token provided", http.StatusUnauthorized)
			return
		}
		headerArr := strings.Split(authHeader, " ")
		if len(headerArr) != 2 {
			utils.SendError(w, "Unauthorized: Invalid token format", http.StatusUnauthorized)
			return
		}
		if headerArr[0] != "Bearer" {
			utils.SendError(w, "Unauthorized: Invalid token format", http.StatusUnauthorized)
			return
		}
		token := headerArr[1]
		tokenParts := strings.Split(token, ".")
		if len(tokenParts) != 3 {
			utils.SendError(w, "Unauthorized: Invalid token format", http.StatusUnauthorized)
			return
		}
		jwtHeader := tokenParts[0]
		jwtPayload := tokenParts[1]
		jwtSignature := tokenParts[2]
		message := jwtHeader + "." + jwtPayload
		byteSecret := []byte(m.cnf.JwtSecret)
		byteMessage := []byte(message)
		h := hmac.New(sha256.New, byteSecret)
		h.Write(byteMessage)
		hash := h.Sum(nil)
		newSignature := base64UrlEncode(hash)
		if newSignature != jwtSignature {
			utils.SendError(w, "Unauthorized: Invalid token", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
func base64UrlEncode(input []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(input)
}
