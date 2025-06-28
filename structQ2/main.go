package main

import "fmt"

//defining the struct
type Car struct {
	Name, Model, Color string
	WeightInKg float64
}

func main() {
	c := Car{Name: "Nexon", Model: "XM", Color: "Grey", WeightInKg: 1920}

	//Accessing struct feilds
	//using the dot operator
	fmt.Println("Car Name :",c.Name)
	fmt.Println("Car Color :",c.Color)

	//Assiging a new value
	//to a struct field
    c.Color = "Black"

	//Displaying the result
	fmt.Println("car :",c)
}