package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func passwordVerification() {

	reader := bufio.NewReader(os.Stdin)

	for {

		fmt.Print("\nEnter Password (type 'exit' to quit): ")

		password, _ := reader.ReadString('\n')
		password = strings.TrimSpace(password)

		if password == "password" {

			fmt.Println("\nLogin Successful!\n")
			mainMenu()
			return

		} else if password == "exit" {

			fmt.Println("\nExiting Program...")
			os.Exit(0)

		} else {

			fmt.Println("Wrong Password!\n")

		}
	}
}