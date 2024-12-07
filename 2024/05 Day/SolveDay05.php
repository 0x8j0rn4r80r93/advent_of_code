<?php
$rules = [];
$updates = [];

// Step 1: Reading input
$lines = file('input.txt', FILE_IGNORE_NEW_LINES | FILE_SKIP_EMPTY_LINES);
$isUpdateSection = false;

foreach ($lines as $line) {
    if (strpos($line, '|') !== false && !$isUpdateSection) {
        // This is a rule line
        list($x, $y) = explode('|', $line);
        $x = (int) $x;
        $y = (int) $y;
        $rules[$x][] = $y;
    } else {
        // We are now reading updates
        $isUpdateSection = true;
        $updatePages = array_map('intval', explode(',', $line));
        $updates[] = $updatePages;
    }
}

// Step 2 & 3: Check each update
$total = 0;
foreach ($updates as $update) {
    // Build a quick lookup map for indices
    $indexMap = [];
    foreach ($update as $i => $p) {
        $indexMap[$p] = $i;
    }

    $isCorrect = true;
    // Filter and check rules
    foreach ($rules as $x => $dependents) {
        foreach ($dependents as $y) {
            if (isset($indexMap[$x]) && isset($indexMap[$y])) {
                if ($indexMap[$x] >= $indexMap[$y]) {
                    $isCorrect = false;
                    break 2; // Break out of both loops
                }
            }
        }
    }

    if ($isCorrect) {
        // Step 4: Get the middle page
        $midIndex = floor(count($update)/2);
        $total += $update[$midIndex];
    }
}

// Step 5: Print result
echo $total;
