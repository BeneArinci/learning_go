package main

import "fmt"

type person struct {
	name    string
	surname string
	age     int
}

func main() {
	bene := person{"Bene", "Arinci", 32}
	mySelf := person{name: "Bene", surname: "Arinci", age: 32}
	fmt.Println(bene)
	fmt.Println(mySelf)
}
