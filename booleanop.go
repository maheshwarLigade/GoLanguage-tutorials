package main

import "fmt"

func logicalAnd(){
	var a = true && true
    var b = true && false
    var c = false && true
    var d = false && false

    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)
    fmt.Println(d)
}

func logicalOR(){

	var a = true || true
    var b = true || false
    var c = false || true
    var d = false || false

    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)
    fmt.Println(d)
}

func logicalNot(){
	fmt.Println(!true)
    fmt.Println(!false)
    fmt.Println(!(4 < 3))
}

func gtlt(){
	var x = 3
    var y = 8
    // equality check 
	fmt.Println(x == y)
	// greater than operator 
    fmt.Println(y > x)

    if y > x {

        fmt.Println("y is greater than x")
    }
}

func main(){

	gtlt()
	fmt.Println()
	fmt.Println("Logical AND")
	logicalAnd()
	fmt.Println()
	fmt.Println("Logical OR ")
	logicalOR()
	fmt.Println()
	fmt.Println("Logical NOT")
	logicalNot()
	fmt.Println()
}