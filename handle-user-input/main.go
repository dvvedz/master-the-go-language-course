package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inp := bufio.NewReader(os.Stdin)

	fmt.Print("Type something: ")

	// Get user input and exit "type mode" when newline char is detected in the input
	userInp, err := inp.ReadString('\n')

	if err != nil {
		fmt.Println("User input error")
		os.Exit(1)
	}

	fmt.Println("You Wrote:", userInp)
}
