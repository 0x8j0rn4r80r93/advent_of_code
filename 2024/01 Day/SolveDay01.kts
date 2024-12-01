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

    // Sort both lists
    leftList.sort()
    rightList.sort()

    println("Left List (sorted): $leftList")
    println("Right List (sorted): $rightList")

    // Calculate the total distance
    println("Calculating total distance...")
    val totalDistance = leftList.indices.sumOf { index ->
        val distance = kotlin.math.abs(leftList[index] - rightList[index])
        println("Pair: ${leftList[index]} and ${rightList[index]}, Distance: $distance")
        distance
    }

    // Print the result
    println("Total Distance: $totalDistance")
}

main()