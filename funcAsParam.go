package main

import "fmt"

// addition of two no
func add(x int, y int) int {
    
    return x+y
}

// substraction of two no
func sub(x int, y int) int {
    
    return x-y
}

// pass above two function as argument
func apply(x int, y int, f func(int,int) int) int {

    r := f(x,y)
    return r
}
// higher order function to return function 
func getAddSub() (func(int, int) int, func(int, int) int) {

    add := func(x, y int) int {
        return x + y
    }

    sub := func(x, y int) int {
        return x - y
    }

    return add, sub
}

func applyReturn(x, y int, add func(int, int) int, sub func(int, int) int) (int, int) {

    r1 := add(x, y)
    r2 := sub(x, y)

    return r1, r2
}

func main() {

    x := 3
    y := 2

    r1 := apply(x,y, add)
    r2 := apply(x,y, sub)
    fmt.Println(r1)
    fmt.Println(r2)

    // return function
    add, sub := getAddSub()
    fmt.Println("function as a return type")
    re1, re2 := applyReturn(x, y, add, sub)

    fmt.Printf("%d + %d = %d\n", x, y, re1)
    fmt.Printf("%d - %d = %d\n", x, y, re2)
}