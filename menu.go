package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func mainMenu() {

	reader := bufio.NewReader(os.Stdin)

	for {

		fmt.Println("========== MAIN MENU ==========")
		fmt.Println("1. Admin")
		fmt.Println("2. Student")
		fmt.Println("3. Exit")

		fmt.Print("Enter Choice : ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {

		case "1":
			adminMenu()

		case "2":
			viewStudents()

		case "3":
			fmt.Println("Good Bye!")
			os.Exit(0)

		default:
			fmt.Println("Invalid Choice")

		}

	}

}

func adminMenu() {

	reader := bufio.NewReader(os.Stdin)

	for {

		fmt.Println("\n========== ADMIN ==========")

		fmt.Println("1. Add Student")
		fmt.Println("2. View Students")
		fmt.Println("3. Delete Student")
		fmt.Println("4. Back")

		fmt.Print("Enter Choice : ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {

		case "1":
			addStudent()

		case "2":
			viewStudents()

		case "3":
			deleteStudent()

		case "4":
			return

		default:
			fmt.Println("Invalid Choice")

		}

	}

}