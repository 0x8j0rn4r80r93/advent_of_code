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
