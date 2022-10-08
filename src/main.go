package main

import "github.com/bobbyrc/flowmeter/src/routes"

func main() {
	router := routes.Initialize()

	router.Run() // listen and serve on localhost:8080
}
