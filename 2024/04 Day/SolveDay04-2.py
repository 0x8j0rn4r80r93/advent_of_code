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
