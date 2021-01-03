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

func main() {
    r1 := apply(3,2, add)
    r2 := apply(3,2, sub)
    fmt.Println(r1)
    fmt.Println(r2)
}