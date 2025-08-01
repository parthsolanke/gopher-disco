package student

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Student struct {
	Name  string
	Marks []float64
}

func ScanStudent(scanner *bufio.Scanner, subjectCount int) (Student, error) {
	var student Student

	fmt.Println("\nEnter student details")

	fmt.Print("Name: ")
	scanner.Scan()
	student.Name = scanner.Text()

	fmt.Printf("Marks (space-separated, %d subjects): ", subjectCount)
	scanner.Scan()
	markStrs := strings.Fields(scanner.Text())

	if len(markStrs) != subjectCount {
		return student, fmt.Errorf("expected %d marks, got %d", subjectCount, len(markStrs))
	}

	for _, str := range markStrs {
		mark, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return student, fmt.Errorf("invalid mark: %v", err)
		}
		student.Marks = append(student.Marks, mark)
	}

	return student, nil
}

func ComputeAverage(marks []float64) float64 {
	total := 0.0
	for _, m := range marks {
		total += m
	}
	return total / float64(len(marks))
}

func GradeFromAverage(avg float64) string {
	switch {
	case avg >= 90:
		return "A"
	case avg >= 75:
		return "B"
	case avg >= 60:
		return "C"
	case avg >= 40:
		return "D"
	default:
		return "F"
	}
}

func PrintReportCard(students ...Student) {
	for _, s := range students {
		avg := ComputeAverage(s.Marks)
		grade := GradeFromAverage(avg)
		fmt.Printf("\nReport for %s:\n", s.Name)
		fmt.Printf("Marks: %v\n", s.Marks)
		fmt.Printf("Average: %.2f | Grade: %s\n", avg, grade)
	}
}
