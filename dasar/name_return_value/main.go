package main

import "fmt"

func getName() (firstName string, lastName string) {

	firstName = "Bayu"
	lastName = "fajar sidik"

	return
}

func main() {
	firstName, lastName := getName()

	fmt.Println(firstName)
	fmt.Println(lastName)
}
