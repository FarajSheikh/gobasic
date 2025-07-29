package main

import "fmt"

func main() {
 
	num := 10

	if num == 10 {
		fmt.Println("ye to 10 hai ....")
	} else if num == 11 {
		fmt.Println("ye to 11 hai ....")
	} else {
		fmt.Println("na 10 na 11, kuch aur hi hai, ye hai", num)
	}
}