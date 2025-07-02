package main

import "fmt"

func main() {
	rvar := []string{"GFG", "Geeks", "GeeksForGeeks", "kevin"}

	rvar1 := []int{103,104,105,106}

	// i and j stores the value of rvariable
	// i store index number of individual string and
	// j store individual string of the given array

	for i, j := range rvar {
		fmt.Println(i,j)
	}

	for i,j := range rvar1 {
		fmt.Println(i,j)
	} 
}