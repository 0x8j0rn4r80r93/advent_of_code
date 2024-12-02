# Advent of Code - Day 1: Historian Hysteria

## Introduction

The Chief Historian has gone missing, and the Elvish Senior Historians need your help! By reconciling two lists of historically significant location IDs, we can uncover the truth about the Chief Historian's journey.

Advent of Code presents two puzzles each day, granting stars upon successful completion. Solve both parts of the challenge for Day 1 and help save Christmas!

---

### Part One

The lists of location IDs from two groups of Elves must be reconciled by pairing the smallest number from one list with the smallest number from the other. For each pair, calculate the distance between the numbers and sum these distances to find the **total distance**.

**Example Input:**

<table>
  <thead>
    <tr>
      <th>Left List</th>
      <th>Right List</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td>3</td>
      <td>4</td>
    </tr>
    <tr>
      <td>4</td>
      <td>3</td>
    </tr>
    <tr>
      <td>2</td>
      <td>5</td>
    </tr>
    <tr>
      <td>1</td>
      <td>3</td>
    </tr>
    <tr>
      <td>3</td>
      <td>9</td>
    </tr>
    <tr>
      <td>3</td>
      <td>3</td>
    </tr>
  </tbody>
</table>


Maybe the lists are only off by a small amount! To find out, pair up the numbers and measure how far apart they are. **Pair up the smallest number in the left list with the smallest number in the right list**, then the second-smallest left number with the second-smallest right number, and so on.

Within each pair, figure out how far apart the two numbers are; you'll need to add up all of those distances. For example:
- If you pair up a `3` from the left list with a `7` from the right list, the distance apart is `4`.
- If you pair up a `9` with a `3`, the distance apart is `6`.

In the example list above, the pairs and distances would be as follows:

| Pair            | Distance |
|------------------|----------|
| `1` and `3`     | `2`      |
| `2` and `3`     | `1`      |
| `3` and `3`     | `0`      |
| `3` and `4`     | `1`      |
| `3` and `5`     | `2`      |
| `4` and `9`     | `5`      |

### Total Distance

To find the total distance between the left list and the right list, add up the distances between all of the pairs you found. In the example above, this is:

2 + 1 + 0 + 1 + 2 + 5 = 11


### Part Two

Analyze the lists further to compute a **similarity score**. Multiply each number in the left list by the number of times it appears in the right list and sum these values.

**Example Input:**


**Example Calculation:**
- Count occurrences of each number in the right list:
    - 3 appears 3 times
    - 4 appears 1 time
    - 5 appears 1 time
    - 9 appears 1 time

- Multiply each number in the left list by its frequency in the right list:
    - 3 × 3 = 9
    - 4 × 1 = 4
    - 2 × 0 = 0
    - 1 × 0 = 0
    - 3 × 3 = 9
    - 3 × 3 = 9

**Similarity Score:** `9 + 4 + 0 + 0 + 9 + 9 = 31`

---

## Results

For your input:

- **Part One - Total Distance:** `2031679`
- **Part Two - Similarity Score:** `19678534`

---

## Solution in Kotlin

Below is the Kotlin implementation for solving both parts of the Day 1 challenge.

```kotlin
import java.io.File

fun main() {
    println("Script started.") // Confirm the script is running

    // Read input file
    val inputFile = File("input.txt")
    if (!inputFile.exists()) {
        println("Error: input.txt file not found!")
        return
    }

    println("Reading input file...")

    // Parse the input
    val leftList = mutableListOf<Int>()
    val rightList = mutableListOf<Int>()

    inputFile.forEachLine { line ->
        println("Processing line: $line") // Print each line being processed
        val (left, right) = line.split("\\s+".toRegex()).map { it.toInt() }
        leftList.add(left)
        rightList.add(right)
    }

    println("Left List (unsorted): $leftList")
    println("Right List (unsorted): $rightList")

    // Sort both lists for part one
    leftList.sort()
    rightList.sort()

    println("Left List (sorted): $leftList")
    println("Right List (sorted): $rightList")

    // Calculate the total distance for part one
    println("Calculating total distance for Part One...")
    val totalDistance = leftList.indices.sumOf { index ->
        val distance = kotlin.math.abs(leftList[index] - rightList[index])
        println("Pair: ${leftList[index]} and ${rightList[index]}, Distance: $distance")
        distance
    }
    println("Total Distance (Part One): $totalDistance")

    // Calculate the similarity score for part two
    println("Calculating similarity score for Part Two...")
    val rightListFrequency = rightList.groupingBy { it }.eachCount() // Count occurrences in the right list
    println("Right List Frequencies: $rightListFrequency")

    val similarityScore = leftList.sumOf { leftNumber ->
        val countInRight = rightListFrequency.getOrDefault(leftNumber, 0)
        val contribution = leftNumber * countInRight
        println("Left Number: $leftNumber, Count in Right List: $countInRight, Contribution: $contribution")
        contribution
    }
    println("Similarity Score (Part Two): $similarityScore")

    // Print both results
    println("Results:")
    println("Part One - Total Distance: $totalDistance")
    println("Part Two - Similarity Score: $similarityScore")
}

main()
```