# Mull It Over Challenge - Java Solution

This repository contains a Java-based solution for the challenge **"Mull It Over"**. The challenge involves analyzing corrupted memory to extract valid `mul` instructions and compute their results based on specified conditions.

---

## Problem Statement

### Part 1: Extract Valid Multiplication Instructions
Given a corrupted memory string, identify valid multiplication instructions in the format `mul(X,Y)`, where `X` and `Y` are integers. Ignore invalid instructions and sum the results of the valid multiplications.

**Example Input:**
xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))


**Valid Instructions:**
- `mul(2,4)` -> Result: 8
- `mul(5,5)` -> Result: 25
- `mul(11,8)` -> Result: 88
- `mul(8,5)` -> Result: 40

**Output:**
Total Sum: 161

### Part 2: Handle Conditional Statements
Extend the solution to handle `do()` and `don't()` instructions:
- `do()` enables future `mul` instructions.
- `don't()` disables future `mul` instructions.

Only enabled `mul` instructions are processed, and the most recent condition applies.

**Example Input:**
xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))

**Valid Instructions (considering conditions):**
- `mul(2,4)` -> Result: 8
- `mul(8,5)` -> Result: 40

**Output:**
Total Sum: 48


---

## Solution

The solution consists of a Java script that:
1. Reads input from a file (`input.txt`).
2. Extracts and processes valid `mul` instructions using regex patterns.
3. Handles enabling/disabling conditions for instructions in Part 2.

---

## Code

### Main.java
```java
import java.nio.file.Files;
import java.nio.file.Paths;
import java.io.IOException;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

public class Main {
    public static void main(String[] args) {
        task1();
        task2();
    }

    public static void task1() {
        String filePath = "../input.txt";
        try {
            String content = new String(Files.readAllBytes(Paths.get(filePath)));
            Pattern pattern = Pattern.compile("mul\\((\\d+),(\\d+)\\)");
            Matcher matcher = pattern.matcher(content);
            int totalSum = 0;
            while (matcher.find()) {
                int num1 = Integer.parseInt(matcher.group(1));
                int num2 = Integer.parseInt(matcher.group(2));
                totalSum += num1 * num2;
            }
            System.out.println("Task 1 | Total sum: " + totalSum);
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    public static void task2() {
        String filePath = "../input.txt";
        try {
            String content = new String(Files.readAllBytes(Paths.get(filePath)));
            Pattern pattern = Pattern.compile("mul\\((\\d+),(\\d+)\\)|don't\\(\\)|do\\(\\)");
            Matcher matcher = pattern.matcher(content);
            int totalSum = 0;
            boolean isDisabled = false;
            while (matcher.find()) {
                String match = matcher.group();
                if (match.equals("don't()")) {
                    isDisabled = true;
                } else if (match.equals("do()")) {
                    isDisabled = false;
                } else if (!isDisabled) {
                    int num1 = Integer.parseInt(matcher.group(1));
                    int num2 = Integer.parseInt(matcher.group(2));
                    totalSum += num1 * num2;
                }
            }
            System.out.println("Task 2 | Total sum: " + totalSum);
        } catch (IOException e) {
            e.printStackTrace();
        }
    }
}
```