package main

import "fmt"

type Triangle struct {
    h int
	w int
	l int
}

func main() {

    t1 := Triangle{2, 3,4}
    t2 := Triangle{2, 3,4}

    if t1 == t2 {

        fmt.Println("The triangel structs are equal")
    } else {

        fmt.Println("The triangel structs are not equal")
    }
}