package main

import "fmt"

type Employee struct {
    name       string
    occupation string
}

func change(pu *Employee) {

    pu.name = "Mahesh"
    pu.occupation = "engineer"
}

func main() {

    u := Employee{"Mahesh", "teacher"}
    fmt.Println(u)

    change(&u)

    fmt.Println(u)
}