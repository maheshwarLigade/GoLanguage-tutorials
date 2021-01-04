package main

import "fmt"

type Employee struct {
    name       string
    department string
    age        int
}

func main() {

	e := Employee{"Ram", "Operation", 42}
	
	//empty assignment
	e1 := Employee{}
	e1.name="Rocky"
	e1.department="IT"
	e1.age=28


	fmt.Printf("%s is %d years old and he is  part of a %s department\n", e.name, e.age, e.department)
	fmt.Println("\n Dot Operator assigment result\n")
	fmt.Printf("%s is %d years old and he is  part of a %s department\n", e1.name, e1.age, e1.department)
}