// package common

package main

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const (
	tokenSignKey  = "www,Akaki,com"
	tokenLifeHour = 72
)

func main() {
	token, err := GenToken([]byte(tokenSignKey))
	if err != nil {
		fmt.Println("Creating token failed")
	}
	ParseToken(token, tokenSignKey)
}

/// 生成token
func GenToken(signKey []byte) (string, error) {
	// Create the token
	token := jwt.New(jwt.SigningMethodHS256)
	// Set some claims
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * tokenLifeHour).Unix()
	token.Claims = claims
	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString(signKey)
	return tokenString, err
}

/// 解析验证token
func ParseToken(tokenStr, signKey string) (token *jwt.Token, err error) {
	jwtToken, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(signKey), nil
	})
	fmt.Println("token value is = ", jwtToken)
	if err == nil && jwtToken.Valid {
		fmt.Println("Your token is valid.  I like your style.")
		return jwtToken, nil
	}
	fmt.Println("This token is terrible!  I cannot accept this.")
	return nil, err
}
