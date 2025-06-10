package main

import "fmt"

func age(n int) int {
	if n < 18 {
		fmt.Println("you are under 18")
	}else {
		fmt.Println("you are over 18")
	}
	
return n

}
func main() {
    age(20)
}