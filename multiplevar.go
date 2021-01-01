package main

import "fmt"

func main() {

    var x, y, z = 24, 25, 26

    var (
        name       = "Mahesh Ligade"
        occupation = "Engineer"
    )

    fmt.Println(x, y, z)
    fmt.Printf("%s is a %s\n", name, occupation)
}