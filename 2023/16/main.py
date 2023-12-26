
def parse_input(i):
    g = {}

    with open(f"{i}.txt", "r") as f:
        num_cols = -1
        num_rows = -1
        lines = f.readlines()
        num_rows = len(lines)
        num_cols = len(lines[0])
        for x, line in enumerate(lines):
            for y, c in enumerate(line):
                g[(x, y)] = c
        return g, num_rows, num_cols

RIGHT = (0, 1)
LEFT = (0, -1)
UP = (-1, 0)
DOWN = (1, 0)

def move(t: tuple[int, int], dir: tuple[int, int]):
    return (t[0] + dir[0], t[1] + dir[1])


def move_beam_through_grid(g: dict[tuple[int, int], str], start_position: tuple[tuple[int, int], tuple[int, int]]):

    # current location + direction of travel
    beams = [start_position]

    seen_states = set()

    i = 0
    while len(beams) > 0:
        i += 1
        beam = beams.pop(0)
        current_location = beam[0]
        direction = beam[1]

        if beam in seen_states:
            # we've already been to this cell going this direction
            # this likely implies a loop, no sense continuing
            continue
        if current_location in g:
            seen_states.add(beam)

        next_location = move(current_location, direction)

        if next_location not in g:
            continue

        next_cell = g[next_location]

        if next_cell == '.':
            beams.append((next_location, direction))  
        elif next_cell == '\\':
            if direction == RIGHT:
                # --->\
                #     |
                #     v
                beams.append((next_location, DOWN))
            elif direction == LEFT:
                # ^
                # |
                # \ <---
                beams.append((next_location, UP))
            elif direction == UP:
                # <-- \
                #     ^
                #     |
                beams.append((next_location, LEFT))
            elif direction == DOWN:
                #  |
                #  v
                #  \ -->
                beams.append((next_location, RIGHT))
        elif next_cell == '/':
            if direction == RIGHT:
                #     ^
                #     |
                # --> /
                beams.append((next_location, UP))
            elif direction == LEFT:
                # / <---
                # |
                # v
                beams.append((next_location, DOWN))
            elif direction == UP:
                # / -->
                # ^
                # |
                beams.append((next_location, RIGHT))
            elif direction == DOWN:
                #   |
                #   v
                # <-/
                beams.append((next_location, LEFT))
        elif next_cell == '|':
            if direction == UP or direction == DOWN:
                beams.append((next_location, direction))
            else:
                beams.append((next_location, UP))
                beams.append((next_location, DOWN))
        elif next_cell == '-':
            if direction == RIGHT or direction == LEFT:
                beams.append((next_location, direction))
            else:
                beams.append((next_location, LEFT))
                beams.append((next_location, RIGHT))

    return len({b[0] for b in seen_states})

def part_1(i):
    grid, _, _ = parse_input(i)
    return move_beam_through_grid(grid, ((0, -1), RIGHT))

def part_2(i):
    grid, num_rows, num_cols = parse_input(i)

    max_energy = -1

    for i in range(num_rows):
        max_energy = max(
            max_energy,
            move_beam_through_grid(grid, ((i, -1), RIGHT)),
            move_beam_through_grid(grid, ((i, num_cols), LEFT))
        )
    for i in range(num_cols):
        max_energy = max(
            max_energy,
            move_beam_through_grid(grid, ((-1, i), DOWN)),
            move_beam_through_grid(grid, ((num_rows, i), UP))
        )
    
    return max_energy

print(part_1("input"))
print(part_2("input"))