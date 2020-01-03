package main

import (
	"encoding/base64"
	"fmt"
)

type Container struct {
	Key   string
	Value interface{}
}

func main() {
	f1()
}

func f1() {
	v := `{"msg": "ok"}`
	encoded := base64.StdEncoding.EncodeToString([]byte(v))
	fmt.Printf("\n%s", encoded)

	decoded, _ := base64.StdEncoding.DecodeString(encoded)
	fmt.Printf("\n%s", decoded)
}

/*
eyJtc2ciOiAib2sifQ==
{"msg": "ok"}
*/
