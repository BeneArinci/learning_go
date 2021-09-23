package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}
type person struct {
	name    string
	surname string
	age     int
	contactInfo
}

func main() {

	// initializing person variables
	bene := person{"Bene", "Arinci", 32, contactInfo{"Flat 21", 11111}}

	mySelf := person{
		name:    "Bene",
		surname: "Arinci",
		age:     32,
		contactInfo: contactInfo{
			email:   "aaa@gmail.com",
			zipCode: 10101,
		},
	}

	var gianlu person

	// updating a struct values
	gianlu.name = "Gianlu"
	gianlu.age = 32

	// this creates a pointer to the address bene is stored
	benePointer := &bene

	benePointer.updateName("Beneeeeee")

	bene.print()

	mySelf.print()
	gianlu.print()

}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).name = newFirstName
}
