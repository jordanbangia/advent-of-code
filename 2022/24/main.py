
from functools import cache


def parse_input(filename: str):
    board = []
    with open(filename, 'r') as f:
        for line in f.readlines():
            board.append(list(line.strip()))

    for i in range(len(board)):
        for j in range(len(board[0])):
            if board[i][j] == '#':
                continue
            elif board[i][j] == '.':
                board[i][j] = []
            else:
                board[i][j] = [board[i][j]]
    return board

positions = [(1, 0), (-1, 0), (0, 1), (0, -1), (0, 0)] # last option is to consider sitting


def sim(board, start_position, end_position, start_t = 0) -> int:
    m = len(board) - 1
    n = len(board[0]) - 1

    print(start_position)
    print(end_position)

    @cache
    def board_at_time(t):
        if t == 0:
            return board

        m = len(board)
        n = len(board[0])
        
        next_board = []
        for i, row in enumerate(board):
            r = []
            for j, c in enumerate(row):
                r.append([])
            next_board.append(r)

        for i, row in enumerate(board):
            for j, c in enumerate(row):
                if board[i][j] == '#':
                    next_board[i][j] = '#'
                else:
                    for blizzard in board[i][j]:
                        if blizzard == '<':
                            next_j = j
                            k = t
                            while k > 0:
                                k -= 1
                                next_j -= 1
                                if next_j == 0:
                                    next_j = n-2
                            next_board[i][next_j].append(blizzard)
                        elif blizzard == '^':
                            next_i = i
                            k = t
                            while k > 0:
                                k -= 1
                                next_i -= 1
                                if next_i == 0:
                                    next_i = m - 2
                            next_board[next_i][j].append(blizzard)
                        elif blizzard == '>':
                            next_j = j
                            k = t
                            while k > 0:
                                k -= 1
                                next_j += 1
                                if next_j == n-1:
                                    next_j = 1
                            next_board[i][next_j].append(blizzard)
                        elif blizzard == 'v':
                            next_i = i
                            k = t
                            while k > 0:
                                k -= 1
                                next_i += 1
                                if next_i == m-1:
                                    next_i = 1
                            next_board[next_i][j].append(blizzard)
        return next_board

    seen_states = set()

    # tuple of current position, current steps taken
    q = [(start_position, start_t)]
    while len(q) > 0:
        e = q.pop(0)

        current_pos = e[0]
        current_t = e[1]

        if current_pos == end_position:
            return current_t

        # avoid 
        if e in seen_states:
            continue
        seen_states.add(e)

        next_board = board_at_time(current_t + 1)

        for position in positions:
            next_pos = (current_pos[0] + position[0], current_pos[1] + position[1])
            if next_pos[0] < 0 or next_pos[0] > m or next_pos[1] < 0 or next_pos[1] > n:
                # out of bounds in up/down
                continue
            if board[next_pos[0]][next_pos[1]] == '#':
                # hit a wall, can't go that way
                continue
            if len(next_board[next_pos[0]][next_pos[1]]) != 0:
                # hit a blizzard, can't go
                continue
            # we can go this way
            q.append((next_pos, current_t + 1))

def part_1(filename):
    board = parse_input(filename)
    start_position = (0, [i for i in range(len(board[0])) if len(board[0][i]) == 0][0])
    end_position = (len(board) - 1, [i for i in range(len(board[0])) if len(board[len(board) - 1][i]) == 0][0])

    return sim(board, start_position, end_position)

def part_2(filename):
    board = parse_input(filename)

    start = (0, [i for i in range(len(board[0])) if len(board[0][i]) == 0][0])
    end = (len(board) - 1, [i for i in range(len(board[0])) if len(board[len(board) - 1][i]) == 0][0])

    # first sim start -> end
    start_to_end = sim(board, start, end, 0)
    print("Start to end:", start_to_end)
    # then go back for the snacks
    back_for_snacks = sim(board, end, start, start_to_end)
    print("back for snacks:", back_for_snacks, back_for_snacks - start_to_end)
    # then return with the snacks
    returned = sim(board, start, end, back_for_snacks)
    print("returned:", returned, returned - back_for_snacks)

    return returned

# print(part_1('input.txt'))
print(part_2('input.txt'))