package main

import "fmt"

func main() {
	k, l := 24, 8086

	p := &k          // point to k
	fmt.Println(*p) // read k through the pointer
	*p = 29         // set k through the pointer
	fmt.Println(k)  // see the new value of k

	p = &l         // point to l
	*p = *p / 17   // divide l through the pointer
	fmt.Println(l) // see the n ew value of l
}