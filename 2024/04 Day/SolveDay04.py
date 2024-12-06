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
