package main

import (
    "/exportingstr/model"
    "fmt"
)

func main() {

    e := model.Employee{Name: "Maheshwar"}

    fmt.Printf("Employee Name is %s", e.Name)
}
