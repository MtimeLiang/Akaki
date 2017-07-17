package token

import (
	jwt "github.com/dgrijalva/jwt-go"
)

const (
	SignKey = "www,Akaki,com"
)

// New token
// c := jwt.MapClaims{"username": "liang", "exp": time.Now().Add(time.Hour * 72).Unix(),}
// token, err := New([]byte(SignKey), c)
func New(signKey []byte, c jwt.MapClaims) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = c
	tokenString, err := token.SignedString(signKey)
	return tokenString, err
}

// Parse token
// t, err := Parse(token, SignKey)
func Parse(tokenStr, signKey string) (token *jwt.Token, err error) {
	jwtToken, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(signKey), nil
	})
	if err == nil && jwtToken.Valid {
		return jwtToken, nil
	}
	return nil, err
}

// Get MapClaims
// m, ok := GetMapClaims(token)
func GetMapClaims(t *jwt.Token) (m jwt.MapClaims, ok bool) {
	m, ok = interface{}(t.Claims).(jwt.MapClaims)
	return m, ok
}
