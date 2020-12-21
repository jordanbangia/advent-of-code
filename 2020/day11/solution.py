from copy import deepcopy

grid = []
with open("input.txt", "r") as f:
    for line in f:
        line = line.strip()
        grid += [list(line)]


DELTAS = (-1, -1), (0, -1), (1, -1), (-1, 0), (1, 0), (-1, 1), (0, 1), (1, 1)
TAKEN, FREE = "#L"


def seat(r, c, grid):
    if r < 0 or c < 0:
        return
    try:
        return grid[r][c]
    except:
        return


def count_taken_adjacent(r, c, grid):
    count = 0
    for delta_row, delta_col in DELTAS:
        if seat(r + delta_row, c + delta_col, grid) == TAKEN:
            count += 1
    return count


def count_taken_sight(r, c, grid):
    count = 0
    for delta_row, delta_col in DELTAS:
        row = r + delta_row
        col = c + delta_col
        while True:
            s = seat(row, col, grid)
            if s in [None, FREE]:
                break
            if s == TAKEN:
                count += 1
                break
            row += delta_row
            col += delta_col
    return count


def iter(grid, count_fn, max_occupied):
    changed = False
    next_state = deepcopy(grid)

    for i, row in enumerate(grid):
        for j, seat in enumerate(row):
            occupied = count_fn(i, j, grid)
            if seat == FREE and occupied == 0:
                next_state[i][j] = TAKEN
                changed = True
            elif seat == TAKEN and occupied >= max_occupied:
                next_state[i][j] = FREE
                changed = True
    return next_state, changed


def solve_first(grid):
    changed = True
    while changed:
        grid, changed = iter(grid, count_taken_adjacent, 4)

    print("1 occupied seats:", sum(row.count(TAKEN) for row in grid))


def solve_second(grid):
    changed = True
    while changed:
        grid, changed = iter(grid, count_taken_sight, 5)

    print("2 occupied seats:", sum(row.count(TAKEN) for row in grid))


solve_first(grid)
solve_second(grid)