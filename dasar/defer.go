package main

import "fmt"

func logging() {
	fmt.Println("Selesai logging")
}

func runApplication(value int) {
	defer logging()

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func main() {
	runApplication(0)
}
