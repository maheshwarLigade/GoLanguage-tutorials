package main

import "fmt"

var nasaword string = "PSLV"

func main() {

    n := 10

    fmt.Println(nasaword)
    fmt.Println(n)

    newScope()
}

func newScope() {

    fmt.Println(nasaword)
}