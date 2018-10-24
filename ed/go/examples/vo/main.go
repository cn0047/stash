package main

import (
	"fmt"

	"./NewUserVO"
)

// GOPATH=$PWD/ed/go/examples/vo
func main() {
	vo, err := NewUserVO.New(map[string]string{})
	//vo, err := NewUserVO.New(map[string]string{"name": "bond", "email": "bond@mi6.com"})
	if err != nil {
		fmt.Printf("got error: %s, \nvo errors: %+v\n", err, vo.GetErrors())
		return
	}

	createNewUser(vo)
}

func createNewUser(vo NewUserVO.Instance) {
	saveNewUserIntoDB(vo)
	sendEmailToNewUser(vo)
	addNewUserIntoSearch(vo)
}

func saveNewUserIntoDB(vo NewUserVO.Instance) {
}

func sendEmailToNewUser(vo NewUserVO.Instance) {
}

func addNewUserIntoSearch(vo NewUserVO.Instance) {
}
