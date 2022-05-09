package main

import (
	"fmt"
)

func main() {
	/* Initialize
	Manifest Typing
	Desc: Pendeklarasian variable dengan menggunakan type data nya sekalian
	*/
	// var firstName string = "Nazilla"
	// var lastName string = "Nazilla"

	/* Initialize
	Type inference
	Desc: Pendeklarasian variable tanpa menggunakan type data
	*/
	firstName2 := "Bayu"
	lastName2 := "Fajar"

	fmt.Printf("halo %s %s!\n", firstName2, lastName2)
}
