package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, world")
	app := App{}
	port := "9090"
	app.initApp()
	app.Run(port)
}
