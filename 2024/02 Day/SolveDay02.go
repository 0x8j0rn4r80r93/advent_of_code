package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	green = "\033[32m"
	red   = "\033[31m"
	reset = "\033[0m"
)

func main() {
	fmt.Println("Starting the program...") // Debug print

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	fmt.Println("File opened successfully") // Debug print

	scanner := bufio.NewScanner(file)
	lineNumber := 1
	safeCount := 0
	unsafeCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		if isSafe(line) {
			fmt.Printf("%sLine %d is Safe: %s%s\n", green, lineNumber, line, reset)
			safeCount++
		} else {
			fmt.Printf("%sLine %d is Unsafe: %s%s\n", red, lineNumber, line, reset)
			unsafeCount++
		}
		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("%sTotal Safe reports: %d%s\n", green, safeCount, reset)
	fmt.Printf("%sTotal Unsafe reports: %d%s\n", red, unsafeCount, reset)

	fmt.Println("Finished processing the file") // Debug print
}

func isSafe(report string) bool {
	levels := strings.Fields(report)
	if len(levels) < 2 {
		fmt.Println("Not enough levels to determine safety") // Debug print
		return false
	}

	first, err := strconv.Atoi(levels[0])
	if err != nil {
		fmt.Println("Error converting first level to integer:", err) // Debug print
		return false
	}
	increasing := true
	decreasing := true

	for i := 1; i < len(levels); i++ {
		current, err := strconv.Atoi(levels[i])
		if err != nil {
			fmt.Println("Error converting level to integer:", err) // Debug print
			return false
		}
		diff := current - first

		if diff < -3 || diff > 3 || diff == 0 {
			return false
		}

		if current > first {
			decreasing = false
		} else if current < first {
			increasing = false
		}

		first = current
	}

	return increasing || decreasing
}
