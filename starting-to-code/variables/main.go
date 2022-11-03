package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	// # Declaring variables
	// can't start with a number
	// should use a lowercase first letter
	// should use camelCase

	// Declare then assign value
	// var firstNumber int
	// firstNumber = 1

	// Another way
	// var secondeNumber = 2

	// Declare name, assign value and let go figure put type
	// thirdNumber := 3

	// fmt.Println(firstNumber, secondeNumber, thirdNumber)

	// Number guessing game
	rand.Seed(time.Now().UnixNano())

	var firstNumber int = rand.Intn(8) + 2
	var secondNumber int = 5 + 2
	var substraction int = 7 + 2
	var answer int = firstNumber*secondNumber - substraction

	guessTheNumber(firstNumber, secondNumber, substraction, answer)

}

func guessTheNumber(fn int, sn int, sub int, answer int) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Guess the number game")
	fmt.Println("---------------------")
	fmt.Println("Write down a number and press ENTER")
	// fmt.Print("-> ")

	// usrInp, _ := reader.ReadString('\n')

	fmt.Println("Multiply your number by", fn, "and press ENTER when ready.")
	reader.ReadString('\n')

	fmt.Println("Now multiply the result by", sn, "and press ENTER when ready.")
	reader.ReadString('\n')

	fmt.Println("Divide the result by the number you origanlly thought of and press ENTER when ready.")
	reader.ReadString('\n')

	fmt.Println("Now subtract", sub, "and press ENTER when ready.")
	reader.ReadString('\n')

	fmt.Println("The answer is:", answer)
}
