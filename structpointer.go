package main

import "fmt"

type Circle struct {
    redius int
    circum int
}

func main() {

    c := Circle{3, 4}

    c_p := &c

    (*c_p).redius = 1
    c_p.circum = 2

    fmt.Println(c)
}