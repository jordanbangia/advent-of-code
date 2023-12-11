

from dataclasses import dataclass
import re


def parse_input(i) -> tuple[list[list[str]], tuple[int, int]]:
    grid = []

    start_position = None

    with open(f"{i}.txt", "r") as f:
        for line in f.readlines():
            grid.append(
                [c for c in line.strip()]
            )
            if "S" in line:
                start_position = (len(grid) - 1, line.index("S"))
        

        return grid, start_position


@dataclass
class Grid:

    g: list[list]

    num_rows: int = -1
    num_cols: int = -1

    def __post_init__(self):
        self.num_rows = len(self.g)
        self.num_cols = len(self.g[0])

    def at(self, pos):
        if pos[0] < 0 or pos[0] > self.num_rows:
            return None
        if pos[1] < 0 or pos[1] > self.num_cols:
            return None
        return self.g[pos[0]][pos[1]]
    
    def clamp(self, pos):
        r = pos[0]
        r = min(max(r, 0), self.num_rows - 1)

        c = pos[1]
        c = min(max(c, 0), self.num_cols - 1)

        return (r, c)



def next_pos(current_pos, grid: Grid):
    pipe = grid.at(current_pos)
    if pipe == '|':
        return {
            (current_pos[0] - 1, current_pos[1]),
            (current_pos[0] + 1, current_pos[1]),
        }
    elif pipe == '-':
        return {
            (current_pos[0], current_pos[1] - 1),
            (current_pos[0], current_pos[1] + 1),
        }
    elif pipe == 'L':
        return {
            (current_pos[0] - 1, current_pos[1]),
            (current_pos[0], current_pos[1] + 1),
        }
    elif pipe == 'J':
        return {
            (current_pos[0] - 1, current_pos[1]),
            (current_pos[0], current_pos[1] - 1)
        }
    elif pipe == '7':
        return {
            (current_pos[0] + 1, current_pos[1]),
            (current_pos[0], current_pos[1] - 1)
        }
    elif pipe == 'F':
        return {
            (current_pos[0] + 1, current_pos[1]),
            (current_pos[0], current_pos[1] + 1)
        }
    elif pipe == '.':
        return {}


def _start_symbol(start_pos, grid: Grid):
    right_symbol = grid.at((start_pos[0], start_pos[1] + 1))
    left_symbol = grid.at((start_pos[0], start_pos[1] - 1))
    up_symbol = grid.at((start_pos[0] - 1, start_pos[1]))
    down_symbol = grid.at((start_pos[0] + 1, start_pos[1]))

    if left_symbol in {'-', 'F', 'L'} and  right_symbol in {'-', 'J', '7'}:
        return '-'
    elif up_symbol in {'7', '|', 'F'} and down_symbol in {'|', 'L', 'J'}:
        return '|'
    elif up_symbol in {'7', '|', 'F'} and right_symbol in {'-', 'J', '7'}:
        return 'L'
    elif up_symbol in {'7', '|', 'F'} and left_symbol in {'-', 'L', 'F'}:
        return 'J'
    elif down_symbol in {'|', 'L', 'J'} and left_symbol in {'-', 'L', 'F'}:
        return '7'
    elif down_symbol in {'|', 'L', 'J'} and right_symbol in {'-', 'J', '7'}:
        return 'F'
    
    raise ValueError("somethings wrong")

def part_1(i):
    grid, start_pos = parse_input(i)
    grid[start_pos[0]][start_pos[1]] = _start_symbol(start_pos, Grid(grid))

    seen_positions = {start_pos}

    grid = Grid(grid)

    print(grid.at(start_pos))

    next_positions = next_pos(start_pos, grid)

    i = 1
    while len(next_positions) == 2 and len(next_positions.intersection(seen_positions)) == 0:
        next_next_pos = set()
        for pos in next_positions:
            n = next_pos(pos, grid).difference(seen_positions)
            assert len(n) == 1
            next_next_pos.add(list(n)[0])
            seen_positions.add(pos)
        i += 1
        next_positions = next_next_pos
        # if i == 3:
        #     break

    # print(next_positions)
    # print(seen_positions)
    # print(i)
    return i

def part_2(i):
    grid, start_pos = parse_input(i)
    grid[start_pos[0]][start_pos[1]] = _start_symbol(start_pos, Grid(grid))

    grid = Grid(grid)
    print("start tile", grid.at(start_pos))

    loop_positions = {start_pos}
    next_positions = next_pos(start_pos, grid)

    # pick any pose connected to the start pos
    cur_pose = list(next_positions)[0]
    last_pose = start_pos
    loop_positions.add(cur_pose)
    while cur_pose != start_pos:
        for p in next_pos(cur_pose, grid):
            if p == last_pose:
                continue
            else:
                loop_positions.add(cur_pose)
                last_pose = cur_pose
                cur_pose = p
                break

    print("num loop tiles", len(loop_positions))


    # replace any tile tha isn't the grid tile as it doesn't matter
    for i in range(grid.num_rows):
        for j in range(grid.num_cols):
            if (i,j) not in loop_positions:
                grid.g[i][j] = '.'

    # we're going to use the even odd rule
    # .E. | ..O. | ..E. | .O.. |  .. E...
    # any '.' in an ODD grouping is counted towards
    # the inside count

    # use a regex to match the 3 boundary conditions
    boundary_regex = re.compile(r"\||F-*J|L-*7")

    inside_count = 0
    for row in grid.g:
       row_str = ''.join(row)
       
       # splits line by one of 3 boundary conditions:
       # | splits into 2 sections
       # anything that looks like F-----J
       # anything that looksl ike L---7
       sections = boundary_regex.split(row_str)

       # the sections will split on these boundary conditions
       # anything in the odd sections will be counted towards our count
       inside_count += sum([
           section.count(".") for section in sections[1::2]
       ])
    
    return inside_count

print(part_1("input"))
print(part_2("input"))