package main

import "fmt"

// read more than one value using Scanf
func read2value(){
	var name string
    var age int

    fmt.Print("Enter your name & age: ")
    fmt.Scanf("%s %d", &name, &age)
    fmt.Printf("%s is %d years old\n", name, age)
}

func main() {

    var name string

    fmt.Print("Enter your good name: ")
    fmt.Scanf("%s", &name)
	fmt.Println("Hello ", name)
	
	read2value()
}