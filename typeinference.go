package main

import (
    "fmt"
    "reflect"
)

func main() {

    var name = "Mahesh Ligade"
    var age = 29

    fmt.Println("Type of name ",reflect.TypeOf(name))
    fmt.Println("Type of age ",reflect.TypeOf(age))

    fmt.Printf("%s is %d years old\n", name, age)
}