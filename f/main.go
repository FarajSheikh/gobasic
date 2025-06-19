package main

import "fmt"

func getlanguages() (string, string, string){
	return "golang", "javascript", "c"
}

 func main(){
	
	lang1, lang2, lang3 := getlanguages()
	fmt.Println(lang1,lang2,lang3)
	//fmt.Println(getlanguages())
 }