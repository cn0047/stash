package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

type Token struct {
	UserID int `json:"user_id"`
}

func (t *Token) Valid() error {
	fmt.Printf("valid: %#v\n", t)
	return nil
}

func main() {
	t := `token here`

	two(t)
}

func three(t string) {
	tk, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		hmacSampleSecret := ""
		return hmacSampleSecret, nil
	})

	if claims, ok := tk.Claims.(jwt.MapClaims); ok && tk.Valid {
		fmt.Printf("%v", claims)
	} else {
		fmt.Printf("err: %#v", err)
	}
}

func two(t string) {
	var claims Token
	tk, err := jwt.ParseWithClaims(t, &claims, func(token *jwt.Token) (interface{}, error) {
		fmt.Printf("🎾 alg: %#v\n", token.Header["alg"])
		fmt.Printf("🎾 t: %#v\n", token)
		return []byte(""), nil
	})
	fmt.Printf("🍄 %v\n", err)
	fmt.Printf("✅ %v\n", claims.UserID)
	fmt.Printf("%#v\n", tk)
}

func one(t string) {
	tk, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		fmt.Printf("🎾 t: %#v\n", token)
		return nil, nil
	})
	claims := tk.Claims.(jwt.MapClaims)
	fmt.Printf("🍄 %v\n", err)
	fmt.Printf("✅ %#v\n", claims)
}
