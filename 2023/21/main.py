def parse_input(i) -> tuple[dict[tuple[int, int], str], int, int, tuple[int, int]]:
    g = {}

    with open(f"{i}.txt", "r") as f:
        start = None
        lines = f.readlines()

        num_rows = len(lines)
        num_cols = len(lines[0].strip())

        for r, line in enumerate(lines):
            for c, x in enumerate(line.strip()):
                if x == "S":
                    start = (r, c)
                    g[(r, c)] = "."
                else:
                    g[(r, c)] = x

        return g, num_rows, num_cols, start


directions = [
    (1, 0),
    (-1, 0),
    (0, 1),
    (0, -1),
]


def add(p: tuple[int, int], d: tuple[int, int]):
    return (p[0] + d[0], p[1] + d[1])


def available_positions_after_n_steps(i, n=1, debug=False) -> int:
    grid, num_rows, num_cols, start = parse_input(i)

    current_positions = set([start])

    steps = 0
    while steps < n:
        next_positions = set()
        if debug:
            print(current_positions)
        for position in current_positions:
            for direction in directions:
                t = add(position, direction)
                if t in grid and grid[t] != "#":
                    next_positions.add(t)
        current_positions = next_positions
        steps += 1

    if debug:
        for row in range(num_rows):
            line = "".join(
                [
                    grid[(row, col)] if (row, col) not in current_positions else "O"
                    for col in range(num_cols)
                ]
            )
            print(line)

    return len(current_positions)


print(available_positions_after_n_steps("input", 1000))


"""
For part 2 - we need to go large number of steps
I don't think its about finding a pattern

For a given position we can know which squares are achievable in n steps
We need to go 26501365
65 + 131 * something

There's some amount of repeated functionality
And if we consider the input, there's a straight line 
up down / left right of the start position.

So there's some amount of periodicity to everything
And some parity to which cells could actually be filled

636350496972143


I read a bunch and could not figure out how to solve this one.  Brutal.

"""
