package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

// func main() {
// 	jon := person{
// 		firstName: "Jon",
// 		lastName:  "Wick",
// 		contactInfo: contactInfo{
// 			email:   "jon@email.com",
// 			zipCode: 90101,
// 		},
// 	}
//
// 	// jonPointer := &jon
// 	jon.updateName("jonny")
// 	jon.print()
// }

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
