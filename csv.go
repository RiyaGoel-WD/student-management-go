package main

import (
	"encoding/csv"
	"os"
)

const fileName = "students.csv"

func saveStudents() {

	file, _ := os.Create(fileName)

	defer file.Close()

	writer := csv.NewWriter(file)

	defer writer.Flush()

	for _, s := range students {

		writer.Write([]string{

			s.RollNumber,
			s.Name,
			s.College,
			s.Branch,
			s.Attendance,
			s.Physics,
			s.Chemistry,
			s.Maths,
		})

	}

}

func loadStudents() {

	students = nil

	file, err := os.Open(fileName)

	if err != nil {

		return

	}

	defer file.Close()

	reader := csv.NewReader(file)

	records, _ := reader.ReadAll()

	for _, r := range records {

		if len(r) < 8 {
			continue
		}

		s := Student{

			RollNumber: r[0],
			Name: r[1],
			College: r[2],
			Branch: r[3],
			Attendance: r[4],
			Physics: r[5],
			Chemistry: r[6],
			Maths: r[7],
		}

		students = append(students, s)

	}

}