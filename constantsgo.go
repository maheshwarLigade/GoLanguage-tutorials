package main

import "fmt"

func main() {

    var age int = 34
    const WEIGHT = 72

    age = 35
    age = 36

	// uncomment below line and check the second output
    // WEIGHT = 75

    fmt.Println(age, WEIGHT)
}