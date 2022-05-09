package main

import "fmt"

func getName() (string, string) {
	return "Bayu", "fajar sidik"
}

func main() {
	firstName, lastName := getName()
	fmt.Println(firstName)
	fmt.Println(lastName)
}
