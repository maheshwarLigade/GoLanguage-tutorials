package main

import "fmt"

type Address struct {
    city    string
	state string
	pincode string
}

type Employee struct {
    name    string
    age     int
    address Address
}

func main() {

    e := Employee{
        name: "Mahesh",
        age:  24,
        address: Address{
            city:    "Pune",
			state: "India",
			pincode: "413310",
        },
    }

    fmt.Println("Name:", e.name)
    fmt.Println("Age:", e.age)
	fmt.Println("City:", e.address.city)
	fmt.Println("PinCode:", e.address.pincode)
    fmt.Println("State:", e.address.state)
}