package main

import "fmt"

// define function
type display func(string) string

// create function which will display message on console
func greetMessage(name string) string {

    return fmt.Sprintf("Hello %s", name)
}

func goodMorningGreet(name string) string {

    return fmt.Sprintf("Good morning %s", name)
}

func goodEveningGreet(name string) string {

    return fmt.Sprintf("Good evening %s", name)
}


func main() {
    var f display // define variable of type display

    f = greetMessage
	fmt.Println(f("Maheshwar"))
	
	f = goodMorningGreet
	fmt.Println(f("Maheshwar"))

	f = goodEveningGreet
	fmt.Println(f("Maheshwar"))
}