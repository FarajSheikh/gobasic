package main

import "fmt"

//  func add(a int, b int)int{
// 	return a+b
//  }

//go can return multiple value
// func getlanguages() (string, string, string){
// 	return "golang", "javascript", "c"
//}


// passes function inside a processIt function
//  func processIt(fn func(a int) int){
// 	fn(1)
//  }


// go can return a function inside a function
 func processIt() func(a int) int {
	return func(a int) int {
		return 4
	}
 }
 

 func main(){
	// fn := func(a int) int {
	// 	return 2
	// }
    // processIt(fn)

	fn := processIt()
	fn(6)
	
	// fmt.Println(add(3,5))
	//lang1, lang2, lang3 := getlanguages()
	//fmt.Println(lang1,lang2,lang3)
	// lang1, lang2, _ := getlanguages()
	// fmt.Println(lang1, lang2)
    // fmt.Println(getlanguages())
 }