package main 

import "fmt"

func main(){

    text := make(chan string)

    go func() { text <- "channel received" }()

    msg := <-text
    fmt.Println(msg)

}