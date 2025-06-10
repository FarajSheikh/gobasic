package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["age"] = 30
	m["price"] = 100

	fmt.Println("Your age is ",m["age"],"and Price",m["price"])
}