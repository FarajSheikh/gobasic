// jab kisi function ke sath structure laga de usse method kahte hai

package main

import "fmt"

type UserA struct {
	Name      string
	Address   string
	ContactNo int
}

type UserB struct {
	Name      string
	Address   string
	ContactNo int
}

func main() {
	var kevin UserA
	kevin.Name = "Kevin eleven"
	kevin.Address = "mirzapur"
	kevin.ContactNo = 34324423
	elev := UserB{
		Name:      "elevelev",
		Address:   "Varansi",
		ContactNo: 123213123,
	}
	kevin.AddUser(5)
	kevin.addUserC(4) 
	elev.AddUser(6)
	add()
}

func add() {
	fmt.Println("add function run ho raha hai ")
}

func (user UserA) AddUser(a int) {
	fmt.Println("User B wala method chal raha hai, user =", user)
}

func (u UserB) AddUser(a int) {
	fmt.Println("User B wala method chal raha hai, user =", u)
}

func (user UserA) addUserC(a int) {
	fmt.Println("User C wala method chal raha hai, user =", user)
}
