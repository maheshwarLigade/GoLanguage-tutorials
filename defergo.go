package main

import "fmt"

func main() {

    fmt.Println("starting")

    defer greet()

    fmt.Println("ending")
}

func greet() {

    fmt.Println("Good morning")
}