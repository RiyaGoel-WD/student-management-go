package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var students []Student

func addStudent() {

	reader := bufio.NewReader(os.Stdin)

	var s Student

	fmt.Print("Roll Number : ")
	s.RollNumber, _ = reader.ReadString('\n')
	s.RollNumber = strings.TrimSpace(s.RollNumber)

	fmt.Print("Name : ")
	s.Name, _ = reader.ReadString('\n')
	s.Name = strings.TrimSpace(s.Name)

	fmt.Print("College : ")
	s.College, _ = reader.ReadString('\n')
	s.College = strings.TrimSpace(s.College)

	fmt.Print("Branch : ")
	s.Branch, _ = reader.ReadString('\n')
	s.Branch = strings.TrimSpace(s.Branch)

	fmt.Print("Attendance : ")
	s.Attendance, _ = reader.ReadString('\n')
	s.Attendance = strings.TrimSpace(s.Attendance)

	fmt.Print("Physics : ")
	s.Physics, _ = reader.ReadString('\n')
	s.Physics = strings.TrimSpace(s.Physics)

	fmt.Print("Chemistry : ")
	s.Chemistry, _ = reader.ReadString('\n')
	s.Chemistry = strings.TrimSpace(s.Chemistry)

	fmt.Print("Maths : ")
	s.Maths, _ = reader.ReadString('\n')
	s.Maths = strings.TrimSpace(s.Maths)

	students = append(students, s)

	saveStudents()

	fmt.Println("\nStudent Added Successfully\n")

}

func viewStudents() {

	loadStudents()

	fmt.Println()

	for _, s := range students {

		fmt.Println("------------------------------------")

		fmt.Println("Roll :", s.RollNumber)
		fmt.Println("Name :", s.Name)
		fmt.Println("College :", s.College)
		fmt.Println("Branch :", s.Branch)
		fmt.Println("Attendance :", s.Attendance)
		fmt.Println("Physics :", s.Physics)
		fmt.Println("Chemistry :", s.Chemistry)
		fmt.Println("Maths :", s.Maths)

	}

	fmt.Println("------------------------------------")

}

func deleteStudent() {

	reader := bufio.NewReader(os.Stdin)

	loadStudents()

	fmt.Print("Enter Roll Number : ")

	roll, _ := reader.ReadString('\n')
	roll = strings.TrimSpace(roll)

	var temp []Student

	found := false

	for _, s := range students {

		if s.RollNumber != roll {

			temp = append(temp, s)

		} else {

			found = true

		}

	}

	students = temp

	saveStudents()

	if found {

		fmt.Println("Deleted Successfully")

	} else {

		fmt.Println("Student Not Found")

	}

}