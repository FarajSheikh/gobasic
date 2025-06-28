package main

import "fmt"

type Employee struct {
	firstName, lastName string
	age, salary int
}
func main() {
	
	//passing the address of struct variable
	//emp8 is a pointer to the Employee struct
	emp8 := &Employee{"Kevin", "Eleven", 21, 6000}

	// (*emp8). firstName is the syntax to access
	// the firstName field of the emp8 struct
	fmt.Println("first Name :",(*emp8).firstName)
	fmt.Println("Age :",(*emp8).age)

	fmt.Println("first Name :",emp8.firstName)
	fmt.Println("Age :",emp8.age)
	
}