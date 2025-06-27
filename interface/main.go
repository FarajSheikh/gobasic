package main

import "fmt"

func main(){

	val := 122
	val2 := "122231"

	// interface can be give any type of value of datatype 
	var interfaceExample interface{}

	interfaceExample = val
	fmt.Println("Interface value =",interfaceExample)

	interfaceExample = val2
	fmt.Println("Interface value =",interfaceExample)

	interfaceExample = false
	fmt.Println("Interface value =",interfaceExample)


}