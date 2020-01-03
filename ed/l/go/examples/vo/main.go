package main

import (
	"fmt"

	"./CreateUserVO"
)

// GOPATH=$PWD/ed/go/examples/vo
func main() {
	//vo, err := CreateUserVO.New(map[string]string{})
	//vo, err := CreateUserVO.New(map[string]string{"name": "bond", "emailx": "bond@mi6.com"})
	vo, err := CreateUserVO.New(map[string]string{"name": "bond", "email": "bond@mi6.com"})
	if err != nil {
		fmt.Printf("got error: %s, \nvo errors: %+v\n", err, vo.GetErrors())
		return
	}

	createNewUser(vo)
}

func createNewUser(vo CreateUserVO.Instance) {
	saveNewUserIntoDB(vo)
	sendEmailToNewUser(vo)
	addNewUserIntoSearch(vo)
}

func saveNewUserIntoDB(vo CreateUserVO.Instance) {
}

func sendEmailToNewUser(vo CreateUserVO.Instance) {
}

func addNewUserIntoSearch(vo CreateUserVO.Instance) {
}
