package main

import "fmt"

var whatToSay string = "Hello World"

func main() {
	// // 1. Create a new instance of the router
	// router := NewRouter()

	// // 2. Launch the server, listening to port 8080
	// router.Run(":8080")

	fmt.Println(whatToSay)
	fmt.Println(saySomething())
}

func saySomething() string {
	return "something"
}
