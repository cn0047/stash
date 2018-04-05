package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	m1 := map[string]string{"code": "200", "msg": "OK"}
	j, _ := json.Marshal(m1)
	fmt.Printf("\n%s", j)

	m2 := map[string]interface{}{"code": 200, "msg": "OK"}
	j2, _ := json.Marshal(m2)
	fmt.Printf("\n%s", j2)

	a1 := [2]string{"200", "OK"}
	j3, _ := json.Marshal(a1)
	fmt.Printf("\n%s", j3)
}

/*
{"code":"200","msg":"OK"}
{"code":200,"msg":"OK"}
["200","OK"]
*/
