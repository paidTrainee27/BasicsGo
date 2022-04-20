package main

import (
	"errors"
	"fmt"
	"go-basics/logs"
)

type student struct {
	id      int
	name    string
	address address
}

type address string

func (a address) validate() (err error) {
	if a == "" {
		err = errors.New("student address cant be empty")
	}
	return
}
func (s student) String() string {
	return fmt.Sprintf("%s with id %d and address %s", s.name, s.id, s.address)
}

func (s *student) updateAddress(address address) {
	s.address = address
}

func main() {
	// createStruct()
	// anonymousStruct()
	updateStruct()
}

func updateStruct() {
	user1 := &student{}
	user1.id = 2
	user1.name = "aayan"
	user1.address = "street back"
	fmt.Println("student created as", *user1)
	user1.updateAddress("New street front")
	fmt.Println("student updates as", *user1)
}

func createStruct() {
	user1 := &student{}
	user1.id = 1
	user1.name = "ryaaz"
	user1.address = "street back"

	if err := user1.address.validate(); err != nil {
		logs.PrintError(fmt.Errorf("Error %s", err))
	} else {
		fmt.Println("student created as", *user1)
	}
}

func anonymousStruct() {
	user2 := struct {
		id      int
		name    string
		address address
	}{1, "Mr.X", " Wall street"}

	logs.PrintJson(user2)
}
