package main

import "fmt"

func main() {

    x := 7
    y := 3

    z := addition(x, y)

    fmt.Printf("addition: %d\n", z)
}

func addition(a int, b int) int {

    return a + b
}