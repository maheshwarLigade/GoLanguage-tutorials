package main

import "fmt"

func main() {
	k, l := 24, 8086

	p := &k          // point to k
	fmt.Println("Value of k from pointer = ",*p) // read k through the pointer
	fmt.Println("Value of P ",p) // value of p
	*p = 29         // set k through the pointer
	fmt.Println("Value of k = ",k)  // see the new value of k

	p = &l         // point to l
	*p = *p / 17   // divide l through the pointer
	fmt.Println("Value of l = ",l) // see the n ew value of l
}