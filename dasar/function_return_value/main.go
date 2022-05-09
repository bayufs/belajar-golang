package main

import "fmt"

func sayHello(name string) string {

	var data string

	if name == "" {
		data = "Bree"
	} else {
		data = name
	}

	return data

}

func main() {
	result := sayHello("")

	fmt.Println(result)
}
