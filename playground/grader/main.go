package main

import (
	"bufio"
	"fmt"
	"grader/student"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter number of students: ")
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	var students []student.Student

	for i := 0; i < n; i++ {
		s, err := student.ScanStudent(scanner, 3)
		if err != nil {
			fmt.Println("Invalid input:", err)
			i--
			continue
		}
		students = append(students, s)
	}

	fmt.Println("\nReport cards: ")
	student.PrintReportCard(students...)
}
