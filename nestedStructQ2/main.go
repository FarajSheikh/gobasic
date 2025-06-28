package main

import "fmt"

type Student struct {
	name, branch string
	year         int
}

type Teacher struct {
	name, subject string
	exp           int
	details       Student
}

func main() {
	result := Teacher{
		name: "Suman", subject: "java", exp: 5, details: Student{"Bongo", "CSE", 2},
	}

	fmt.Println("\nDatails of the teacher")
	fmt.Println("Teacher's name: ",result.name)
	fmt.Println("Subject: ", result.subject)
    fmt.Println("Experience: ", result.exp)

    fmt.Println("\nDetails of Student")
    fmt.Println("Student's name: ", result.details.name)
    fmt.Println("Student's branch name: ", result.details.branch)
    fmt.Println("Year: ", result.details.year)

}
