package main

import "fmt"

func addition(x int, y int) int {

    return x + y
}
// go will omit the type of input parameter
func multiplication(x, y int) int {

    return x * y
}

func main() {

    fmt.Println(addition(5, 4))
    fmt.Println(multiplication(5, 4))
}