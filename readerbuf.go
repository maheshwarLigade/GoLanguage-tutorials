package main

import (
	"fmt"
	"bufio"
	"os"
)

func main(){
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your friend name: ")

	name, _ := reader.ReadString('\n')
	fmt.Printf("Hello %s\n", name)
}