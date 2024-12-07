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
	dampenerExtraSafe := 0

	for scanner.Scan() {
		line := scanner.Text()

		isSafeReport := isSafe(line, 0) // No errors allowed
		isSafeWithDampenerReport := isSafeWithDampener(line)

		if isSafeReport {
			fmt.Printf("%sLine %d is Safe: %s%s\n", green, lineNumber, line, reset)
			safeCount++
		} else {
			fmt.Printf("%sLine %d is Unsafe: %s%s\n", red, lineNumber, line, reset)
			unsafeCount++
		}

		// Dampener adds to the count only if the report wasn't already safe
		if !isSafeReport && isSafeWithDampenerReport {
			fmt.Printf("%sLine %d is Safe with Dampener: %s%s\n", green, lineNumber, line, reset)
			dampenerExtraSafe++
		}

		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	totalSafeWithDampener := safeCount + dampenerExtraSafe

	fmt.Printf("%sTotal Safe reports: %d%s\n", green, safeCount, reset)
	fmt.Printf("%sTotal Unsafe reports: %d%s\n", red, unsafeCount, reset)
	fmt.Printf("%sTotal Safe reports with Dampener: %d%s\n", green, totalSafeWithDampener, reset)

	fmt.Println("Finished processing the file") // Debug print
}

func isSafe(report string, allowedErrors int) bool {
	levels := strings.Fields(report)
	if len(levels) < 2 {
		return false
	}

	nums := make([]int, len(levels))
	for i, level := range levels {
		n, err := strconv.Atoi(level)
		if err != nil {
			return false
		}
		nums[i] = n
	}

	// Check for increasing sequence
	increasingErrors := 0
	for i := 1; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]
		if diff < 1 || diff > 3 {
			increasingErrors++
		} else if nums[i] <= nums[i-1] {
			increasingErrors++
		}
	}

	// Check for decreasing sequence
	decreasingErrors := 0
	for i := 1; i < len(nums); i++ {
		diff := nums[i-1] - nums[i]
		if diff < 1 || diff > 3 {
			decreasingErrors++
		} else if nums[i] >= nums[i-1] {
			decreasingErrors++
		}
	}

	if increasingErrors <= allowedErrors || decreasingErrors <= allowedErrors {
		return true
	}

	return false
}

func isSafeWithDampener(report string) bool {
	levels := strings.Fields(report)
	if len(levels) < 2 {
		return false
	}

	for i := 0; i < len(levels); i++ {
		modifiedLevels := append([]string{}, levels[:i]...)
		modifiedLevels = append(modifiedLevels, levels[i+1:]...)
		if isSafe(strings.Join(modifiedLevels, " "), 0) {
			return true
		}
	}

	return false
}
