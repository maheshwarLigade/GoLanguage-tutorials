package main

import "fmt"

func nextEven() func() int {

    i := 0
    return func() int {
        i+=2
        return i
    }
}

func main() {

    nextInt := nextEven()

    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())
	
	// this is the new calling new state
    nextInt2 := nextEven()
	fmt.Println("new state :: ",nextInt2())
	
	fmt.Println(nextInt())
}