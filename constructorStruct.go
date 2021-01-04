package main

import "fmt"

type Employee struct {
    name       string
    department string
    age        int
}

func newEmployee(name string, department string, age int) *Employee {

    ptr := Employee{name, department, age}
    return &ptr
}

func main() {

    e := newEmployee("Mahesh", "IT", 29)

    fmt.Printf("%s is %d years old and he is  part of a %s department\n", e.name, e.age, e.department)
}