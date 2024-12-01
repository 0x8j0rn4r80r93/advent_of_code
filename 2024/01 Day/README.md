# 🎄 Advent of Code: Day 1 - Historian Hysteria 🎄

The Chief Historian is always present for the big Christmas sleigh launch, but nobody has seen him in months! Last anyone heard, he was visiting locations that are historically significant to the North Pole; a group of Senior Historians has asked you to accompany them as they check the places they think he was most likely to visit.

As each location is checked, they will mark it on their list with a star. They figure the Chief Historian must be in one of the first fifty places they'll look, so in order to save Christmas, you need to help them get fifty stars on their list before Santa takes off on December 25th.

Collect stars by solving puzzles. Two puzzles will be made available on each day in the Advent calendar; the second puzzle is unlocked when you complete the first. Each puzzle grants one star. **Good luck!**

---

## Part 1: Reconciling Lists

You haven't even left yet and the group of Elvish Senior Historians has already hit a problem: their list of locations to check is currently empty. Eventually, someone decides that the best place to check first would be the Chief Historian's office.

Upon pouring into the office, everyone confirms that the Chief Historian is indeed nowhere to be found. Instead, the Elves discover an assortment of notes and lists of historically significant locations! This seems to be the planning the Chief Historian was doing before he left. Perhaps these notes can be used to determine which locations to search?

### The Problem

Throughout the Chief's office, the historically significant locations are listed not by name but by a unique number called the **location ID**. To make sure they don't miss anything, The Historians split into two groups, each searching the office and trying to create their own complete list of location IDs.

There's just one problem: by holding the two lists up side by side (your puzzle input), it quickly becomes clear that the lists aren't very similar. Maybe you can help The Historians reconcile their lists?

### Example

<table>
  <thead>
    <tr>
      <th>Left List</th>
      <th>Right List</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td>3</td>
      <td>4</td>
    </tr>
    <tr>
      <td>4</td>
      <td>3</td>
    </tr>
    <tr>
      <td>2</td>
      <td>5</td>
    </tr>
    <tr>
      <td>1</td>
      <td>3</td>
    </tr>
    <tr>
      <td>3</td>
      <td>9</td>
    </tr>
    <tr>
      <td>3</td>
      <td>3</td>
    </tr>
  </tbody>
</table>


Maybe the lists are only off by a small amount! To find out, pair up the numbers and measure how far apart they are. **Pair up the smallest number in the left list with the smallest number in the right list**, then the second-smallest left number with the second-smallest right number, and so on.

Within each pair, figure out how far apart the two numbers are; you'll need to add up all of those distances. For example:
- If you pair up a `3` from the left list with a `7` from the right list, the distance apart is `4`.
- If you pair up a `9` with a `3`, the distance apart is `6`.

In the example list above, the pairs and distances would be as follows:

| Pair            | Distance |
|------------------|----------|
| `1` and `3`     | `2`      |
| `2` and `3`     | `1`      |
| `3` and `3`     | `0`      |
| `3` and `4`     | `1`      |
| `3` and `5`     | `2`      |
| `4` and `9`     | `5`      |

### Total Distance

To find the total distance between the left list and the right list, add up the distances between all of the pairs you found. In the example above, this is:

2 + 1 + 0 + 1 + 2 + 5 = 11


---

### Your Input

Your actual left and right lists contain many location IDs. **What is the total distance between your lists?**

See input.txt.
