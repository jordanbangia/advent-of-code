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

# grid look something like the above
grid = []

trajectory_right = 3
trajectory_down = 1

with open("input.txt", "r") as f:
    for line in f:
        grid += [list(line.strip())]

line_length = len(grid[0])
print(line_length)

location = [0, 0]
trees = 0

location = [
    location[0] + trajectory_right,
    location[1] + trajectory_down,
]

its = 0
while location[1] < len(grid):
    print(location)
    its += 1
    row = grid[location[1]]
    if row[location[0]] == "#":
        trees += 1
        grid[location[1]][location[0]] = "X"
    else:
        grid[location[1]][location[0]] = "0"

    location = [
        location[0] + trajectory_right,
        location[1] + trajectory_down,
    ]
    if location[0] >= line_length:
        location[0] = location[0] % line_length

print(its)
print(trees)

with open("output.txt", "w") as f:
    for r in grid:
        f.write(f'{"".join(r)}\n')
