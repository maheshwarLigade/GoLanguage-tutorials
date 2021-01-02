package main

import (
    "fmt"
    "math/rand" // math random to get random value
    "time"   // time API
)
// function which will return the four random values
func fourrandom() (int, int, int, int) {

    rand.Seed(time.Now().UnixNano())
    a := rand.Intn(20)
    b := rand.Intn(20)
	c := rand.Intn(20)
	d := rand.Intn(20)

    return a, b, c, d
}

func main() {

    a, b, c, d := fourrandom() // result asigning to four variables

    fmt.Println(a, b, c, d)
}