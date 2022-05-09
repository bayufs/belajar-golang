package main

import "fmt"

func newMap(data string) map[string]string {
	if data == "" {
		return nil
	} else {
		return map[string]string{
			"name": data,
		}
	}
}

func main() {
	var person map[string]string = newMap("Nazilla")
	if person == nil {
		fmt.Println("Data kogong")
	} else {
		fmt.Println(person)
	}
}
