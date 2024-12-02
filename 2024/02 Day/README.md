# Advent of Code 2024 - Day 2: Red-Nosed Reports

## --- Day 2: Red-Nosed Reports ---

The first location The Historians want to search isn't far from the Chief Historian's office. While the Red-Nosed Reindeer nuclear fusion/fission plant doesn't seem to have any sign of the Chief Historian, the engineers there request help analyzing unusual data from the Red-Nosed reactor.

The unusual data (your puzzle input) consists of many reports, each report being a list of numbers called **levels**, separated by spaces. For example:
7 6 4 2 1 
1 2 7 8 9 
9 7 6 2 1 
1 3 2 4 5 
8 6 4 4 1 
1 3 6 7 9


Each report is considered **safe** if:
1. The levels are either all increasing or all decreasing.
2. Any two adjacent levels differ by at least `1` and at most `3`.

### Example Analysis

Using the above example:
- `7 6 4 2 1`: **Safe** (all levels are decreasing by `1` or `2`).
- `1 2 7 8 9`: **Unsafe** (the jump from `2` to `7` is an increase of `5`).
- `9 7 6 2 1`: **Unsafe** (the drop from `6` to `2` is a decrease of `4`).
- `1 3 2 4 5`: **Unsafe** (it switches from increasing to decreasing).
- `8 6 4 4 1`: **Unsafe** (adjacent levels `4` and `4` are neither an increase nor a decrease).
- `1 3 6 7 9`: **Safe** (all levels increase by `1`, `2`, or `3`).

In this example, **2 reports** are safe.

### Puzzle Solution

The script determines how many reports are safe in the input data. For my input, the number of safe reports was **526**.

---

## --- Part Two ---

The engineers realized they forgot to account for the **Problem Dampener**.

The Problem Dampener allows the reactor safety systems to tolerate **a single bad level** in an otherwise safe report. Essentially, it’s like the bad level was never there. The same rules apply as before, except removing one level from a report can now make it count as safe.

### Example Analysis with Problem Dampener

Using the same example:
- `7 6 4 2 1`: **Safe** without removing any level.
- `1 2 7 8 9`: **Unsafe** (no single removal makes it safe).
- `9 7 6 2 1`: **Unsafe** (no single removal makes it safe).
- `1 3 2 4 5`: **Safe** (remove the second level, `3`).
- `8 6 4 4 1`: **Safe** (remove the third level, `4`).
- `1 3 6 7 9`: **Safe** without removing any level.

With the Problem Dampener, **4 additional reports** are now safe.

### Updated Puzzle Solution

The script was updated to account for the Problem Dampener. The final count of safe reports after applying the Problem Dampener was determined as part of the solution:

**526 safe reports** (initial count) + **40 extra safe reports** (with Dampener) = **566 total safe reports**.

Both parts of this puzzle are complete! They provide two gold stars: **.

---

## Solution in Go

Below is the Go implementation used to solve this problem:

```go
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

	// Try removing each level one by one
	for i := 0; i < len(levels); i++ {
		modifiedLevels := append([]string{}, levels[:i]...)
		modifiedLevels = append(modifiedLevels, levels[i+1:]...)
		if isSafe(strings.Join(modifiedLevels, " "), 0) {
			return true
		}
	}

	return false
}
```