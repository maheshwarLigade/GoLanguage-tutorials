package main

import "fmt"

type Info func(string, string, int) string

type Employee struct {
    name       string
    department string
    age        int
    info       Info
}

func main() {

    e := Employee{
        name:       "Maheshwar",
        department: "IT",
        age:        41,
        info: func(name string, department string, age int) string {

            return fmt.Sprintf("%s is %d years old and he is  part of a %s department\n", name, age, department)
        },
    }

    fmt.Printf(e.info(e.name, e.department, e.age))
}