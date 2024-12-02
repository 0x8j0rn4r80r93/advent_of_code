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
