package main

import "fmt"

func getFullName(fullName string) string {
	return "Woy " + fullName
}

func main() {

	sayFullName := getFullName

	result := sayFullName("Bayu fajar sidik")

	fmt.Println(result)

}
