package main

import "fmt"

/* A closure is a function value that references variables from outside its body.
the function may access and assign to the referenced variables; in this sense the 
function is "bound" to the variables.  */

 func company() func() int {
	a := 0
    return func() int {
		a++
		return a
	}
 }
  
func main() {
	val := company()
	fmt.Println(val())
	fmt.Println(val())
	fmt.Println(val())

	v := company()
	fmt.Println(v())
}