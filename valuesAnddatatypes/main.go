package main

import "fmt"

func main(){

	const number2 = 786767

	var number int 
	number = 12
	var addressnumber *int
	addressnumber = &number

	var decnum float32
	decnum = 12232.2321
	var addfloat *float32
	addfloat = &decnum

	var flag bool
	flag = true
	var addflag *bool
	addflag = &flag

	var name string
	name = "Kevin"
	var addname *string
	addname = &name

	number3 := 343
	fmt.Println(number3)

	fmt.Println(number)
	fmt.Println(decnum)
	fmt.Println(flag)
	fmt.Println(name)

	fmt.Printf("number value = %d, decimal value = %f, flag value = %v, name value = %s\n",number,decnum,flag,name)
    fmt.Println(number2)
	fmt.Printf("address of number = %v, address of decimal = %v, address of flag = %v,address of name = %v\n",addressnumber,addfloat,addflag,addname)
}