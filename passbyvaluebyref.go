package main

import "fmt"

package main

import "fmt"

type Employee struct {
    name       string
    occupation string
}

func main() {
    // variables scope is within the main()
	var x int = 100
	var y int = 200
 
	fmt.Printf("Before swaping, value of x : %d\n", x )
	fmt.Printf("Before swaping, value of y : %d\n", y )
 
	// call the function with x and y values
	swapbyvalue(x, y)
 
	fmt.Printf("After swaping, value of x : %d\n", x )
	fmt.Printf("After swaping, value of y : %d\n", y )

	e := Employee{"Mahesh Ligade", "teacher"}
    fmt.Printf("inside main %v\n", e)

    modify(e)
    fmt.Printf("inside main %v\n", e)
 }

func swapbyvalue(int v, int w) int {

	var temp int
    
	temp = v // save the value of v
	v = w    // put w into v 
	w = temp // put temp into w 
 
	return temp;
 }

 func modify(e Employee) {

    e.occupation = "driver"
    fmt.Printf("inside modify %v\n", e)
}
