package core

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ProcessIn() (int, int, error) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter 2 numbers: ")

	scanner.Scan()
	input := scanner.Text()
	parts := strings.Fields(input)

	if len(parts) != 2 {
		return 0, 0, errors.New("please enter exactly 2 numbers")
	}

	num1, err1 := strconv.Atoi(parts[0])
	num2, err2 := strconv.Atoi(parts[1])

	if err1 != nil || err2 != nil {
		return 0, 0, errors.New("invalid input, must be integers")
	}

	return num1, num2, nil
}
