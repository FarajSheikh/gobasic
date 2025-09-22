package main

import "fmt"

func newId() func() int {
	i := 0
	return func() int  {
		i++
		return i
	}
}

func main() {
	result := newId()

	fmt.Println("result   ",result())
	fmt.Println("result   ",result())
}