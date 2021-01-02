package main

import "fmt"

func main() {

	 // defined function which take two input and retunr addition of these two
    addition := func(a, b int) int {
        return a + b
    }(9, 1)

    fmt.Println("9+1 =", addition)
}