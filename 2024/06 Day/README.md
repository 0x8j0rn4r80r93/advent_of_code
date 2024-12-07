# Guard Gallivant Challenge

## Overview

In the "Guard Gallivant" challenge, we are tasked with simulating the patrol route of a guard in a North Pole prototype suit manufacturing lab in the year 1518. The goal is to predict the guard's path and calculate how many distinct positions they visit before leaving the mapped area.

The challenge involves processing a map of the lab (input) and following a set of patrol rules to determine the guard's movement. 

---

## Problem Statement

Given a map:
- `.` represents open floor.
- `#` represents obstacles (crates, desks, etc.).
- The guard's starting position is denoted by a symbol indicating their facing direction:
  - `^` (up)
  - `>` (right)
  - `v` (down)
  - `<` (left)

The guard follows these patrol rules:
1. If there's an obstacle directly in front of the guard, they turn **right** by 90 degrees.
2. Otherwise, the guard moves one step forward in their current direction.

The guard continues patrolling until they leave the mapped area.

Your task is to determine how many unique positions the guard visits before leaving the map.

---

## Example Input

```plaintext
....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...
```

## Example Output
41

The example map results in the guard visiting 41 distinct positions.

## How to Run the Solution
Prerequisites
- Node.js installed on your machine.

Steps
1. Save the input map in a file named input.txt.
2. Save the solution code in a file named SolveDay06.js.
3. Run the following command in your terminal:
```powershell
Get-Content input.txt | node SolveDay06.js
```

## Example Output

4752

## Solution Explanation
The solution involves:

1. Parsing the map from input and locating the guard's initial position and facing direction.
2. Simulating the guard's patrol by applying the movement rules:
3. Turning 90 degrees right when encountering an obstacle.
4. Moving forward otherwise.
5. Tracking all distinct positions the guard visits using a Set data structure.
6. Stopping the simulation when the guard moves out of the mapped area.

Solution Code
```javascript
const readline = require('readline');

const lines = [];
const rl = readline.createInterface({
    input: process.stdin,
    crlfDelay: Infinity
});

rl.on('line', (line) => {
    lines.push(line);
});

rl.on('close', () => {
    solve(lines);
});

function solve(mapLines) {
    // Convert input lines into a 2D array
    const map = mapLines.map(line => line.split(''));

    // Directions: Up(0), Right(1), Down(2), Left(3)
    const directionVectors = [
        [-1, 0], // Up
        [0, 1],  // Right
        [1, 0],  // Down
        [0, -1]  // Left
    ];

    const directionSymbols = { '^':0, '>':1, 'v':2, '<':3 };

    // Find the guard's initial position and direction
    let guardRow = -1, guardCol = -1, direction = -1;
    for (let r = 0; r < map.length; r++) {
        for (let c = 0; c < map[r].length; c++) {
            const cell = map[r][c];
            if (directionSymbols[cell] !== undefined) {
                guardRow = r;
                guardCol = c;
                direction = directionSymbols[cell];
                map[r][c] = '.'; // Replace the guard symbol with floor
                break;
            }
        }
        if (guardRow !== -1) break;
    }

    if (guardRow === -1) {
        console.error("No guard found on the map.");
        return;
    }

    const visited = new Set();
    visited.add(`${guardRow},${guardCol}`);

    while (true) {
        const [dr, dc] = directionVectors[direction];
        const newRow = guardRow + dr;
        const newCol = guardCol + dc;

        // Check if next step is off-map
        if (newRow < 0 || newRow >= map.length || newCol < 0 || newCol >= map[0].length) {
            // The guard leaves the map
            break;
        }

        // If front cell is an obstacle, turn right
        if (map[newRow][newCol] === '#') {
            direction = (direction + 1) % 4;
        } else {
            // Move forward
            guardRow = newRow;
            guardCol = newCol;
            visited.add(`${guardRow},${guardCol}`);
        }
    }

    console.log(visited.size);
}
```
