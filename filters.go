package main

import "fmt"

type Employee struct {
    name       string
    department string
    isPermanent    bool
}

func main() {

    e1 := Employee{"Maheshwar", "it", false}
    e2 := Employee{"Ram", "hr", true}
    e3 := Employee{"Sham", "operation", true}
    e4 := Employee{"Govind", "it", false}
    e5 := Employee{"Sri", "admin", true}

	 // array of employees
    emps := []Employee{e1, e2, e3, e4, e5}

	// filter out only permanent employee
    permanent := filter(emps, func(e Employee) bool {
        if e.isPermanent == true {
            return true
        }
        return false
    })

    itwale := filter(emps, func(e Employee) bool {

        if e.department == "it" {
            return true
        }
        return false
    })

    fmt.Println("permanent:")
    fmt.Printf("%v\n", permanent)

    fmt.Println("ITians:")
    fmt.Printf("%v\n", itwale)

}

func filter(s []Employee, f func(Employee) bool) []Employee {
    var res []Employee

    for _, v := range s {

        if f(v) == true {
            res = append(res, v)
        }
    }
    return res
}