package main

import "fmt"

type person struct {
	name    string
	surname string
	age     int
}

func main() {

	// initializing person variables
	bene := person{"Bene", "Arinci", 32}

	mySelf := person{name: "Bene", surname: "Arinci", age: 32}

	var gianlu person

	// updating a struct values
	gianlu.name = "Gianlu"
	gianlu.age = 32

	fmt.Println(bene)
	fmt.Println(mySelf)
	fmt.Println(gianlu)
	fmt.Printf("%+v", gianlu)
}
