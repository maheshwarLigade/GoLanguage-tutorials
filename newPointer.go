package main

import (
    "fmt"
    "reflect"
)

type Employee struct {
    name       string
    occupation string
}

func main() {

    var pu *Employee = new(Employee)
    fmt.Println(pu)
    fmt.Println(reflect.TypeOf(pu))

    pu.name = "Mahesh Ligade"
    pu.occupation = "teacher"
    fmt.Println(pu)
}
