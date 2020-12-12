"""
..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#
"""


def find_trees_for_slope(grid: list, right: int, down: int) -> int:
    line_length = len(grid[0])

    # always start at the top right
    location = [0, 0]
    trees = 0

    location = [
        location[0] + right,
        location[1] + down,
    ]

    its = 0
    while location[1] < len(grid):
        its += 1
        row = grid[location[1]]
        if row[location[0]] == "#":
            trees += 1

        location = [
            location[0] + right,
            location[1] + down,
        ]
        if location[0] >= line_length:
            location[0] = location[0] % line_length
    print(trees)
    return trees


# grid look something like the above
grid = []

with open("input.txt", "r") as f:
    for line in f:
        grid += [list(line.strip())]

one_one = find_trees_for_slope(grid, 1, 1)
three_1 = find_trees_for_slope(grid, 3, 1)
five_1 = find_trees_for_slope(grid, 5, 1)
seven_1 = find_trees_for_slope(grid, 7, 1)
one_2 = find_trees_for_slope(grid, 1, 2)

print(one_one * three_1 * five_1 * seven_1 * one_2)