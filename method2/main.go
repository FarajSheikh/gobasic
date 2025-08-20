package main

import "fmt"

type UserA struct {
	name      string
	address   string
	contactNo int
}

type UserB struct {
	name      string
	address   string
	contactNo int
}

type UserOperations interface {
	addUser()
}

func main() {
	var kevin UserA
	kevin.name = "kevinEleven"
	kevin.address = "Mirzapur"
	kevin.contactNo = 42342345

	elev := UserB{
		name:      "eleven",
		address:   "Varansi",
		contactNo: 432424267,
	}
	// // kevin.addUser()
	// elev.addUser()

	var UserOperations UserOperations
	UserOperations = kevin // UserA
	UserOperations.addUser()

	UserOperations = elev
	UserOperations.addUser()
}

func (user UserA) addUser() {
	user.name = "34"
	fmt.Println("dfsfsdf")
}

func (user UserB) addUser() {
	user.name = "3478"
	fmt.Println("dfsfjjjhjhjhjsdf")
}
