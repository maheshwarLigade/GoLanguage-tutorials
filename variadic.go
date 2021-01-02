package main

import "fmt"

func main() {

    a1 := add(1, 2, 3)
    a2 := add(1, 2, 3, 4)
    a3 := add(1, 2, 3, 4, 5)

    fmt.Println(a1, a2, a3)
}
// accept variable number of input param
func add(nos ...int) int {

    result := 0

    for _, n := range nos {
        result += n
    }

    return result
}