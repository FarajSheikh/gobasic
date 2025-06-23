package main

import "fmt" 

//you can use any type in sum function then use interface like
//func sum(num ...interface{}) int{
//.......}
//but you can not use a plus (+) operator

func sum(num ...int) int {
	total := 0



	for _, num := range num {
		total = total + num
	}
	return total
}

func main(){

	num := []int{3,4,5,6} //--> you can use slice also
	result := sum(num...)

	//result := sum(3,4,5,6)
	fmt.Println(result)
}