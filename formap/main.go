package main

import "fmt"

func main() {
	mmap := map[int]string{
		22: "Geeks",
		33: "GFG",
		44: "GeeksforGeeks",
	}
	for key, value := range mmap {
		fmt.Println(key,value)
	}
}
