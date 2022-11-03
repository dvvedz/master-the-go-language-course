package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader *bufio.Reader

func main() {
	reader = bufio.NewReader(os.Stdin)

	userName := readString("What is yourname?")
	userAge := readInt("How old are you?")

	// this one ismmore efficient than string concatination
	fmt.Printf("Your name is %s. You are %d years old\n", userName, userAge)
}

func prompt() {
	fmt.Print("-> ")
}

func readString(s string) string {
	for {
		fmt.Println(s)
		prompt()
		userInp, _ := reader.ReadString('\n')
		userInp = strings.Replace(userInp, "\r\n", "", -1)
		userInp = strings.Replace(userInp, "\n", "", -1)

		if userInp == "" {
			fmt.Println("Please enter a name...")
		} else {
			return userInp
		}
	}
}

func readInt(s string) int {
	for {

		fmt.Println(s)
		prompt()

		userInp, _ := reader.ReadString('\n')
		userInp = strings.Replace(userInp, "\r\n", "", -1)
		userInp = strings.Replace(userInp, "\n", "", -1)

		num, err := strconv.Atoi(userInp)
		if err != nil {
			fmt.Println("Please enter a whole number...")
		} else {
			return num
		}

	}
}
