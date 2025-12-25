package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
)

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type Payload struct {
	Sub         int    `json:"sub"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	IsShopPwner bool   `json:"is_shop_pwner"`
}

func CreateJWT(secret string, payload Payload) (string, error) {
	header := Header{
		Alg: "HS256",
		Typ: "JWT",
	}
	jsonHeader, err := json.Marshal(header)
	if err != nil {
		return "", err
	}
	base64Header := base64UrlEncode(jsonHeader)
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}
	base64Payload := base64UrlEncode(jsonPayload)

	message := base64Header + "." + base64Payload
	secretHeader := []byte(secret)

	byteArrMessage := []byte(message)
	h := hmac.New(sha256.New, secretHeader)
	h.Write(byteArrMessage)

	signature := h.Sum(nil)
	base64Signature := base64UrlEncode(signature)

	jwt := message + "." + base64Signature

	return jwt, nil
}

func base64UrlEncode(input []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(input)
}
