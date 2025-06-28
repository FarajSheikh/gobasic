package main

import "fmt"

type Author struct {
	Name, branch string
	year int
}

//Creating nested structure
type HR struct {
	datails Author
}

func main() {
	result := HR {
		datails: Author{"Sona", "ECE", 2013},
	}	
	fmt.Println("\nDetails of Author")
	fmt.Println(result)
}