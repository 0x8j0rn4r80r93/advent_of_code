# Print Queue Puzzle Solution

This repository contains a solution for the "Print Queue" puzzle. The problem is about determining if certain updates (lists of page numbers) adhere to given ordering rules, and then summing the "middle page number" of each correctly-ordered update.

## Overview of the Problem

The puzzle is set in a scenario where pages of a safety manual must be printed in a specific order. Each rule `X|Y` dictates that if both pages `X` and `Y` appear in an update, `X` must be printed at some point before `Y`. We apply these rules to each update (a list of pages) and determine whether the given order of pages is correct.

For each correctly-ordered update, we identify the middle page number (based on the update’s length) and sum all such middle pages across all correctly-ordered updates.

**Key Points:**
- Ignore rules that involve pages not present in the current update.
- Only sum the middle page number from updates that fully comply with the rules.
- The final answer is the sum of all these middle page values.

## Steps to Solve

1. **Parsing Input:**
   - The input file contains two sections:  
     1. **Rules Section:** Each line contains a rule in the form `X|Y`.
     2. **Updates Section:** Each subsequent line lists page numbers for one update, separated by commas.
   
   - We first read all lines and treat those with the `|` character as rules. Once we encounter a line that does not match the rule pattern, we consider all subsequent lines as updates.
   
   - Each rule is stored in a data structure mapping `X` to an array of pages `Y` that must come after `X`.

2. **Applying Rules:**
   - For each update, we create a map of `page_number -> index` to easily check the order of any two pages.
   - We iterate through every rule. If both `X` and `Y` of a rule are in the current update, we verify that `indexOf(X) < indexOf(Y)`. If not, the update is disqualified.

3. **Finding the Middle Page:**
   - If an update passes all applicable rules, we find its middle page. For an update with `n` pages, the middle index is `floor(n/2)` (0-based indexing).
   - Add this middle page number to our running total.

4. **Output the Result:**
   - After processing all updates, print out the final total. This total is the sum of the middle pages from every correctly-ordered update.

## Example

Given a small set of rules and updates:
47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47


We apply the logic and find which updates are ordered correctly and then sum their middle pages. For the given test data, the solution was `5129`.

## Common Pitfalls and Troubleshooting

- **Incorrect Separation of Rules and Updates:**  
  Ensure that once you begin reading updates, you do not return to reading rules.
  
- **Handling Missing Pages:**  
  Only apply rules when both pages `X` and `Y` exist in the update. Otherwise, ignore that rule for the current update.
  
- **Ensuring the Correct Middle Page:**  
  Double-check that you always select the middle index correctly and do not confuse it with zero-based or one-based indexing.

- **Confirming With the Full Input:**  
  Partial or test data might not give the final answer. Use the full puzzle input as provided.

## Running the Script

1. Make sure you have PHP installed and added to your system PATH.
2. Place `input.txt` in the same directory as `SolveDay05.php`.
3. Run the script:
   ```bash
   php SolveDay05.php
   ```
4. The script will read the input file, process the data, and output the final sum of middle pages.