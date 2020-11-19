package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName   string
	lastName    string
	contactInfo
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateFirstName(newName string) {
	(*p).firstName = newName
}

func main() {
	var alex person

	alex.firstName = "alex"
	alex.lastName = "anderson"
	jim := person{
		firstName:   "Jim",
		lastName:    "Party",
		contactInfo: contactInfo{
			email: "test@gmail.com",
			zipCode: 1234567,
		},
	}

	jim.updateFirstName("Jimmy")
	jim.print()
}
