package main

import (
	"fmt"
	"strconv"
)

func main() {

	var a int 
	a = 10

	var b int16

	b = int16(a)
	fmt.Println("b :",b)

	//int16 -> string
	var c string
	c = fmt.Sprintf("%d", b)
	fmt.Println("c :",c)

	//string -> int
	num := "124232"
	numint , err := strconv.Atoi(num)
	fmt.Println("numint :",numint, "err :",err)

	//Show error because kevin is string it can not be changed into a int
	num := "Kevin"
	numint , err := strconv.Atoi(num)
	fmt.Println("numint :",numint, "err :",err)

}