package main

import "fmt"

func endApp() {
	fmt.Println("APliaksi selesai")
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("Ada error")
	}

	fmt.Println("Aplikasi berjalan")
}

func main() {
	runApp(true)
}
