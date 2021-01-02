package main

import "fmt"

func change(pv *int) {

    *pv = 21
}

func main() {

    var count int = 44
    fmt.Println(count)

    change(&count)
    fmt.Println(count)
}