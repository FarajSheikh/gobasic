package main

import "fmt"

	type Student struct{

		Name string
		Class int
		RollNumber string
		StudentAddress Address

	}

	type Address struct{
		lame1 string
		lame2 string
		Post string
		Dist string
		Village string
	}

	type Phone struct{
		BasicInfo
		IMEI string
		Config  
	}

	type Laptop struct{
		BasicInfo
		SerialNo string
		Configuratin Config
	}

	type BasicInfo struct{
		Brand string
		Model string
	}

	type Config struct{
		RAM int
		ROM int
		Processor string
	}



func main() {
	
	 
	var kevin Student
	kevin.Name = "kevinelev"
	kevin.Class = 12
	kevin.RollNumber = "103"
	kevin.StudentAddress.Dist = "Mirzapur"  

	faraj := Student{
		Name: "Faraj Sheikh",
		Class: 12,
		RollNumber: "104",
		StudentAddress: Address{
			Dist: "dsasfsdf",
			Village: "sdsd",
		},

	}

	fmt.Println("hello..... team... this is Kevin Struct", kevin)
	
	fmt.Println("hello..... team... this is Kevin Struct", faraj)
}