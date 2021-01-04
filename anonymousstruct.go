package main

import "fmt"

func main() {

    e := struct {
		name       string
		department string
		age        int
    }{
        name: "Ram",
        department: "Operation",
        age: 42,
    }

    fmt.Printf("%s is %d years old and he is  part of a %s department\n", e.name, e.age, e.department)
}