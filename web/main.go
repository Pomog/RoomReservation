package main

import (
	"fmt"
	"net/http"
	"udemyCourse1/pkg/handlers"
)

const port = ":8080"

func main() {
	// // 1. Create a new instance of the router
	// router := NewRouter()

	// // 2. Launch the server, listening to port 8080
	// router.Run(":8080")

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Server starting on port %s\n", port)

	_ = http.ListenAndServe(port, nil)
}
