package main

import "fmt"

type Address struct {
	Name string
	city string
	Pincode int
}

func main(){
    
	//Declaring a variable of a struct type
	//All the struct fields are initailized
	//with their zero value
	var a Address
	fmt.Println(a)

	//Declaring and initializing a 
	//struct using a struct literal
	a1 := Address{"Kevin","varansi",324242}
	fmt.Println("Address1 :",a1)

	//Naming fields while
	//initializing a struct
	a2 := Address{Name: "Ben",city: "ballia",Pincode: 323213}
	fmt.Println("Address2 :",a2)

	//Ininitialized fields are set to
	//their coresponding zero-value
	a3 := Address{Name: "Delhi"}
	fmt.Println("Address3 :",a3)

}
