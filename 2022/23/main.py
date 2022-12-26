
from collections import defaultdict


def parse_input(filename: str):
    d = defaultdict(lambda: '.')

    elves = []

    with open(filename, 'r') as f:
        for x, line in enumerate(f.readlines()):
            for y, c in enumerate(line.strip()):
                d[(x,y)] = c

                if c == '#':
                    elves.append((x,y))
    
    return d, elves

def print_board(board, elves):
    min_x = float('inf')
    max_x = 0
    min_y = float('inf')
    max_y = 0
    for elf in elves.values():
        min_x = min(min_x, elf[0])
        max_x = max(max_x, elf[0])

        min_y = min(min_y, elf[1])
        max_y = max(max_y, elf[1])

    for x in range(min_x, max_x+1):
        print(''.join(board[(x,y)] if not isinstance(board[(x,y)], int) else '#' for y in range(min_y, max_y+1)))

def alone(board, elf):
    for x in range(-1, 2, 1):
        for y in range(-1, 2, 1):
            if x == 0 and y == 0:
                continue
            if board[(elf[0] + x, elf[1] + y)] != '.':
                return False
    return True

def count_empty(board, elves):
    min_x = float('inf')
    max_x = 0
    min_y = float('inf')
    max_y = 0
    for elf in elves.values():
        min_x = min(min_x, elf[0])
        max_x = max(max_x, elf[0])

        min_y = min(min_y, elf[1])
        max_y = max(max_y, elf[1])

    c = 0
    for i in range(min_x, max_x+1, 1):
        for j in range(min_y, max_y+1, 1):
            if board[(i,j)] == '.':
                c += 1
    return c

move_checks = [
    {'name': 'north', 'dirs': [(-1, -1), (-1, 0), (-1, 1)], 'move': (-1, 0)},  # north
    {'name': 'south', 'dirs': [(1, -1), (1, 0), (1, 1)], 'move': (1, 0)}, # south
    {'name': 'west', 'dirs': [(0, -1), (-1, -1), (1, -1)], 'move': (0, -1)}, # west
    {'name': 'east', 'dirs': [(0, 1), (-1, 1), (1, 1)], 'move': (0, 1)} # east
]

def part_1(filename: str, its: int = 10):
    board, elves = parse_input(filename)

    elves = {i: e for i,e in enumerate(elves)}
    
    
    # print_board(board, elves)

    for k in range(its):
        print(k, [m['name'] for m in move_checks])

        moves = defaultdict(list)

        # first half
        # compute each move that the elf could do
        for indx, elf in elves.items():
                if alone(board, elf):
                    continue
                for move in move_checks:
                    if all(
                        board[(elf[0] + d[0], elf[1] + d[1])] == '.' for d in move['dirs']
                    ):
                        moves[(elf[0] + move['move'][0], elf[1] + move['move'][1])].append(indx)
                        break

        # second half, execute the moves
        for new_point, e in moves.items():
            if len(e) > 1:
                continue
            old_point = elves[e[0]]
            board[old_point] = '.'
            elves[e[0]] = new_point
            board[new_point] = '#'
        
        move_checks.append(move_checks.pop(0))
        # print_board(board, elves)

    
    return count_empty(board, elves)

    
def part_2(filename):
    board, elves = parse_input(filename)

    elves = {i: e for i,e in enumerate(elves)}
    
    i = 0
    while True:
        i += 1
        print(i, [m['name'] for m in move_checks])

        moves = defaultdict(list)

        # first half
        # compute each move that the elf could do
        for indx, elf in elves.items():
                if alone(board, elf):
                    continue
                for move in move_checks:
                    if all(
                        board[(elf[0] + d[0], elf[1] + d[1])] == '.' for d in move['dirs']
                    ):
                        moves[(elf[0] + move['move'][0], elf[1] + move['move'][1])].append(indx)
                        break

        # second half, execute the moves
        elf_moved = False
        for new_point, e in moves.items():
            if len(e) > 1:
                continue
            elf_moved = True
            old_point = elves[e[0]]
            board[old_point] = '.'
            elves[e[0]] = new_point
            board[new_point] = '#'
        
        if not elf_moved:
            print_board(board, elves)
            return i
        
        move_checks.append(move_checks.pop(0))

        # if i == 300:
        #     print_board(board, elves)
        #     return -1




print(part_1('input.txt'))

print(part_2('input.txt'))

