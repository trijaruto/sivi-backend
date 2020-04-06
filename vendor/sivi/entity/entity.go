package entity

import "github.com/dgrijalva/jwt-go"

type ResponseHttp struct {
	Code    int8        `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseToken struct {
	Code    int8        `json:"code"`
	Message string      `json:"message"`
	Expired string      `json:"expired"`
	Data    interface{} `json:"data"`
	jwt.StandardClaims
}
