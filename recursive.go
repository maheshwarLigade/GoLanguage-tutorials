package main

import "fmt"

func main() {

    fmt.Println(fact(5))
    fmt.Println(fact(10))
    fmt.Println(fact(3))
}

// take input as number and return the result
func fact(n int) int {

    if n == 0 || n == 1 {
        return 1
    }

    return n * fact(n-1)
}

