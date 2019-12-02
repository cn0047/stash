package main

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

type Token struct {
	UserID int `json:"user_id"`
}

func (t *Token) Valid() error {
	return nil
}

func main() {
	secret := []byte("204")
	t := newToken(secret)
	parseToken(t, secret)
}

func newToken(secret interface{}) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": 100,
	})

	t, err := token.SignedString(secret)
	if err != nil {
		panic(fmt.Errorf("error1: %w", err))
	}

	return t
}

func parseToken(t string, secret interface{}) {
	var claims Token
	_, err := jwt.ParseWithClaims(t, &claims, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		panic(fmt.Errorf("error2: %w", err))
	}

	fmt.Printf("User ID: %#v\n", claims.UserID)
}

func three(t string) {
	tk, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		hmacSampleSecret := ""
		return hmacSampleSecret, nil
	})
	if err != nil {
		panic(fmt.Errorf("error-three: %w", err))
	}

	if claims, ok := tk.Claims.(jwt.MapClaims); ok && tk.Valid {
		fmt.Printf("%v", claims)
	} else {
		fmt.Printf("err: %#v", err)
	}
}

func two(t string) {
	var claims Token
	tk, err := jwt.ParseWithClaims(t, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// return nil, nil // -> key is of invalid type
		return []byte(""), nil // -> signature is invalid
	})

	fmt.Printf("Error: %v\n", err)
	fmt.Printf("Token: %#v\n", tk)
	fmt.Printf("Claims: %#v\n", claims)
}

func one(t string) {
	tk, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		// You can use token here.
		return nil, nil
	})
	claims := tk.Claims.(jwt.MapClaims)
	fmt.Printf("Error: %v\n", err)
	fmt.Printf("Claims: %#v\n", claims)
}
