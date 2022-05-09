package main

import "fmt"

func random() interface{} {
	return "OK"
}

func main() {
	var result interface{} = random()

	switch typeValue := result.(type) {
	case string:
		fmt.Println(result, "Ini adalah string", typeValue)
	case int:
		fmt.Println(result, "Ini adalah int", typeValue)
	default:
		fmt.Println("Unknown value")
	}
}
