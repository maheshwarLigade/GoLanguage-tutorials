package main

import "fmt"

// speified named return variable
func decrement(x, y, z int) (a, b, c int) {

    a = x - 1
    b = y - 1
    c = z - 1

    return
}

func main() {

    x, y, z := decrement(10, 100, 1000)

    fmt.Println(x, y, z)
}
