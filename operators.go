package main

import "fmt"

// sing operatos in go language
func singOperaor(){
	fmt.Println(4)
    fmt.Println(+4)
    fmt.Println(-4)

    var i = 1

    fmt.Println(-i)
    fmt.Println(-(-i))
}
// assignment operator with increment decrement
func assignIncDcr(){
  
    i := 4
    //increment by one i=i+1
    i++
    i++

    fmt.Println(i)
    // decrement by one i=i-1
    i--
    fmt.Println(i)
}
// compound operator
func compoundOperator(){

    var y int = 2
    // addition
    y = y + 2

    fmt.Println(y)
    // addition using comound
    y += 3
    fmt.Println(y)
    // multiplication using compound
    y *= 7
    fmt.Println(y)
}

func main() {

     singOperaor()
     assignIncDcr()
     compoundOperator()
}