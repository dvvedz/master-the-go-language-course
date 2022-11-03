package main

import (
	"fmt"
	"myapp/mypackage"
)

func main() {
	// create a variable and specify type
	var hello string = "Hello"

	// assign variable without specifying type
	world := "World"

	fmt.Println(hello, world)

	// Use dot notation to call mypackage Hello funciton
	fmt.Println(mypackage.Hello("Oskar"))
}
