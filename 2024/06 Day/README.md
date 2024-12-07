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
