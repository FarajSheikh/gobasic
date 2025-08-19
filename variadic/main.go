package main

import "fmt"

func main() {
	fmt.Println("Sum = ",add(5,6,7))
}

func add(b ...int) int {   //variadic function
	sum := 0
	for _ , val := range b {
		sum = sum + val
	}
	return sum

}   