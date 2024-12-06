# Advent of Code 2024 - Day 4: Ceres Search

## Problem Statement

### Part 1
The Chief is missing, and you're at the Ceres monitoring station. A small Elf approaches with a word search puzzle, asking for help finding the word **XMAS** in the grid. The word can appear in any direction:
- Horizontal (left-to-right or right-to-left),
- Vertical (top-to-bottom or bottom-to-top),
- Diagonal (any diagonal direction).

Your task is to count **all occurrences** of the word **XMAS** in the puzzle grid.

#### Example:
Given the following grid:
MMMSXXMASM  
MSAMXMSMSA  
AMXSXMAAMM  
MSAMASMSMX  
XMASAMXAMM  
XXAMMXXAMA  
SMSMSASXSS  
SAXAMASAAA  
MAMMMXMMMM  
MXMXAXMASX  

The word **XMAS** appears **18 times** in total.

---

### Part 2
In Part 2, the task changes: instead of finding **XMAS**, you now need to find patterns that form an **"X-MAS"** shape. An **"X-MAS"** consists of two **MAS** sequences forming an "X", such as:
M.S .A. M.S

Each **MAS** can be written forwards or backwards.

#### Example:
Given the following grid:
.M.S......  
..A..MSMS.  
.M.S.MAA..  
..A.ASMSM.  
.M.S.M....  
..........  
S.S.S.S.S.  
.A.A.A.A..  
M.M.M.M.M.  
..........  


The **X-MAS** pattern appears **9 times**.

---

## Solution

### Part 1 - Counting "XMAS"
This part involves searching the grid for occurrences of the word **XMAS** in all possible directions:
1. Horizontal (left-to-right and right-to-left),
2. Vertical (top-to-bottom and bottom-to-top),
3. Diagonal (all four diagonal directions).

#### Python Script for Part 1:
```python
import os

# Get the directory of the current script
script_dir = os.path.dirname(os.path.abspath(__file__))

# Build the absolute path to the input file
file_path = os.path.join(script_dir, 'input.txt')

# Reading the file content into a list of lines
with open(file_path, 'r') as file:
    grid = [line.strip() for line in file.readlines()]

# Define the word to find
word_to_find = "XMAS"

# Function to count occurrences of the word in all possible directions
def count_word_occurrences(grid, word):
    word_length = len(word)
    total_count = 0
    rows, cols = len(grid), len(grid[0])

    # Function to check in a specific direction
    def check_direction(r, c, dr, dc):
        for i in range(word_length):
            nr, nc = r + dr * i, c + dc * i
            if nr < 0 or nr >= rows or nc < 0 or nc >= cols or grid[nr][nc] != word[i]:
                return 0
        return 1

    # Directions: (dr, dc)
    directions = [
        (0, 1),   # Horizontal right
        (1, 0),   # Vertical down
        (0, -1),  # Horizontal left
        (-1, 0),  # Vertical up
        (1, 1),   # Diagonal down-right
        (-1, -1), # Diagonal up-left
        (1, -1),  # Diagonal down-left
        (-1, 1)   # Diagonal up-right
    ]

    # Iterate over each cell in the grid
    for r in range(rows):
        for c in range(cols):
            for dr, dc in directions:
                total_count += check_direction(r, c, dr, dc)

    return total_count

# Count the occurrences of the word "XMAS"
xmas_count = count_word_occurrences(grid, word_to_find)

# Print the result
print(f"The word '{word_to_find}' occurs {xmas_count} times in the grid.")
```

Part 2 - Counting "X-MAS"
For this part, you need to find two overlapping MAS diagonals forming an "X". Both diagonals must share the same center (A) and can appear forwards or backwards.

Python Script for Part 2:
```python
import os

# Get the directory of the current script
script_dir = os.path.dirname(os.path.abspath(__file__))

# Build the absolute path to the input file
file_path = os.path.join(script_dir, 'input.txt')

# Read the grid
with open(file_path, 'r') as file:
    grid = [line.strip() for line in file.readlines()]

# Function to count "X-MAS" patterns
def count_xmas_patterns(grid):
    rows, cols = len(grid), len(grid[0])
    total_count = 0

    # Iterate over each potential center of the "X"
    for r in range(1, rows - 1):  
        for c in range(1, cols - 1): 
            if grid[r][c] == 'A':  # Center of the "X"
                # Check both diagonals for "MAS"
                top_left_to_bottom_right = (
                    (grid[r-1][c-1] == 'M' and grid[r+1][c+1] == 'S') or
                    (grid[r-1][c-1] == 'S' and grid[r+1][c+1] == 'M')
                )
                top_right_to_bottom_left = (
                    (grid[r-1][c+1] == 'M' and grid[r+1][c-1] == 'S') or
                    (grid[r-1][c+1] == 'S' and grid[r+1][c-1] == 'M')
                )
                # Count only if both diagonals form valid "MAS"
                if top_left_to_bottom_right and top_right_to_bottom_left:
                    total_count += 1

    return total_count

# Count "X-MAS" patterns in the grid
xmas_count = count_xmas_patterns(grid)

# Print the result
print(f"The 'X-MAS' pattern appears {xmas_count} times in the grid.")
```

#### Results
Part 1 Answer: 2613<br>
Part 2 Answer: 1905