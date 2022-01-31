package models

import (
	jwt "github.com/dgrijalva/jwt-go"
)

/*Claim es la estructura usada para procesar el JWT*/
type Claim struct {
	Username string             `json:"username"`
	jwt.StandardClaims
}
