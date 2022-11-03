package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/eiannone/keyboard"
)

func main() {
	// reader := bufio.NewReader(os.Stdin)

	// for {
	// fmt.Print("-> ")

	// userInput, _ := reader.ReadString('\n')
	// userInput = strings.Replace(userInput, "\n", "", -1)

	// if userInput == "quit" {
	// break
	// } else {
	// fmt.Println(userInput)
	// }
	// }

	err := keyboard.Open()

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close
	}()

	coffees := make(map[int]string)

	coffees[1] = "Cappucino"
	coffees[2] = "Latte"
	coffees[3] = "Americano"
	coffees[4] = "Mocha"
	coffees[5] = "Macchiato"
	coffees[6] = "Espresso"

	fmt.Println("MENU")
	fmt.Println("----")
	fmt.Println("1 - Cappucino")
	fmt.Println("2 - Latte")
	fmt.Println("3 - Americano")
	fmt.Println("4 - Mocha")
	fmt.Println("5 - Macchiato")
	fmt.Println("6 - Espresso")
	fmt.Println("Q - Quit the program")

	for {
		char, _, err := keyboard.GetSingleKey()

		if err != nil {
			log.Fatal("Keyboard ERROR", err)
		}

		i, _ := strconv.Atoi(string(char))
		fmt.Println(fmt.Sprintf("You choose %s", coffees[i]))

		if char == 'q' || char == 'Q' {
			break
		}

	}

	fmt.Println("Program exiting...")

}
