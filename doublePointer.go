package main

import "fmt"

func main() {

    var n = 10
    var p = &n // assign address of a to p
    var dp = &p // and then assign address of p to dp

    fmt.Println(n)
    fmt.Println(&n)

    fmt.Println("###################")

    fmt.Println(p)
    fmt.Println(&p)

    fmt.Println("###################")

    fmt.Println(dp)
    fmt.Println(&dp)

    fmt.Println("###################")

    fmt.Println(*dp)
    fmt.Println(**dp) // deference the pointer of pointer value use **
}