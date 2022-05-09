package main

import "fmt"

type HasName interface {
	GetName() string
}

func sayHello(hasName HasName) string {
	fmt.Println("Hello", hasName.GetName())
}

func main() {

}
