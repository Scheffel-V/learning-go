package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	var alex person
	alex.firstName = "Alex"
	alex.lastName = "Anderson"
	alex.contact = contactInfo{"alex@anderson.com", 97400}
	alex.print()

	alexPointer := &alex
	alexPointer.updateName("Alexxe")
	alex.print()

}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Println(p)
}
