
def parse_board(filename: str):
    board = []
    with open(filename, 'r') as f:
        for line in f.readlines():
            board.append(list(line.rstrip()))
    

    max_len = max(len(r) for r in board)
    for i, _ in enumerate(board):
        if len(board[i]) < max_len:
            board[i].extend(list(' ' * (max_len - len(board[i]))))
    
    return board

def parse_commands(filename: str):
    commands = []
    with open(filename, 'r') as f:
        curr = ""
        for c in f.read():
            if c in ('R', 'L'):
                commands.append(int(curr))
                commands.append(c)
                curr = ""
            else:
                curr += c
        if curr != "":
            commands.append(int(curr))
    return commands

def pre_compute_board(_old):
    # for each space on the board that isn't a . or #
    # compute what the next space to loop around too
    # actually only really need to compute this for edge squares
    # empty spaces have [[next square going left, next square going up, next square going right, next square going down]]

    board = [[c for c in r] for r in _old]

    for x, row in enumerate(board):
        for y, c in enumerate(row):
            if board[x][y] in ('.', '#'):
                continue

            i = x
            while board[i][y] not in ('.', '#'):
                i -= 1
                if i == -1:
                    i = len(board) - 1
            
            j = x
            while board[j][y] not in ('.', '#'):
                j += 1
                if j == len(board):
                    j = 0
            
            k = y
            while board[x][k] not in ('.', '#'):
                k -= 1
                if k == -1:
                    k = len(row) - 1
            
            l = y
            while board[x][l] not in ('.', '#'):
                l += 1
                if l == len(row):
                    l = 0
            
            board[x][y] = [(x,l), (j, y), (x,k), (i, y)]

    return board


def increment(x, y, d, board):
    # up -> -x
    # down -> +x
    # left -> -y
    # right -> +y
    
    if d == 0:
        y += 1
    elif d == 1:
        x += 1
    elif d == 2:
        y -= 1
    else:
        x -= 1

    if x == -1:
        x = len(board) - 1
    elif x == len(board):
        x = 0
    
    if y == -1:
        y = len(board[0]) - 1
    elif y == len(board[0]):
        y = 0
    
    return (x, y)


directions = {'R': 0, 'D': 1, 'L': 2, 'U': 3}

field = ('.', '#', 'X', 'A')
change_dirs = ('L', 'R')

def part_1(board, commands):
    board = pre_compute_board(board)

    # start position
    x, y = 0,0
    if board[x][y] not in field:
        x,y = board[x][y][0] # go to the right most

    board[x][y] = 'X'

    d = 0 # start direction is right

    for i, command in enumerate(commands):
        if command in change_dirs:
            if command == 'L':
                d = (d-1) % 4
            else:
                d = (d+1) % 4
        else:
            moves = command
            while moves > 0:
                next_x, next_y = increment(x, y, d, board)

                if board[next_x][next_y] in ('.', 'X', 'A'):
                    moves -= 1
                    x, y = next_x, next_y
                    board[x][y] = 'X'
                elif board[next_x][next_y] == '#':
                    moves = 0
                else:
                    # we're off the field, so we can use the jumps
                    next_x, next_y = board[next_x][next_y][d]
                    if board[next_x][next_y] == '#':
                        moves = 0
                    else:
                        x, y = next_x, next_y
                        board[x][y] = 'X'
                        moves -= 1
    
    board[x][y] = 'A'
    
    for row in board:
        print(''.join([c if c in ('.', '#', 'X', 'A') else ' ' for c in row]))

    return (x+1) * 1000 + (y+1) * 4 + d


# print(part_1(parse_board('input_board.txt'), parse_commands('input_moves.txt')))

def part_2(board, commands, size=4):
    sections = [[None, None, None, None], [None, None, None, None], [None, None, None, None]]

    for x in range(3):
        for y in range(4):
            if board[x*3][y*4] == ' ':
                sections[x][y] = None
            else:
                sections[x][y] = [
                    [board[i+x*3][j+y*4] for j in range(4)] for i in range(4)
                ]

    print(sections)        


print(part_2(parse_board('sample_board.txt'), parse_commands('sample_moves.txt')))


            



