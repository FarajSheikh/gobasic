package main 

import "fmt"

//non primitive data type

type Student struct{
	Name string
	Class int
	RollNumber int
	//Nested Struct -->
	StudentAddress Address
}

//Nested Structure -->
type Address struct{
		lane1 string
		lane2 string
		post string
		Dist string
		Village string

	}

type Phone struct{
    BasicInfo
	IMEI string
	Config
}

type Laptop struct {
	BasicInfo
	SerialNo string
	Configuration Config
}

type BasicInfo struct {
	Brand string
	Model string
}

type Config struct {
	Ram int
	ROM int 
	processor string
	OS string
}


func main(){
	var kevin Student

	kevin.Name = "kevinElev"
	kevin.Class = 12
	kevin.RollNumber = 103
	kevin.StudentAddress.Dist = "varansi"
	kevin.StudentAddress.Village = "jashohar"

	ben := Student{
		Name: "Benten",
		Class: 11,
		RollNumber: 100,

		//Nested Structure
		StudentAddress: Address{
			Dist: "sdsd",
			Village: "dsas",
			lane1: "sdad",
		},
	}


	fmt.Println(kevin)
	fmt.Println(ben)
}