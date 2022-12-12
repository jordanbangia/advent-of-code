

def parse_input(file_name: str) -> tuple[list[list[int]], list[tuple[int, int]], tuple[int, int]]:
    grid = []

    starts = []
    end = None

    with open(file_name, "r") as f:
        for i, line in enumerate(f.readlines()):
            line = line.strip()

            row = []
            for j, c in enumerate(line):
                if c == "S":
                    c = 'a'
                elif c == "E":
                    end = (i, j)
                    c = 'z'
                if c == 'a':
                    starts.append((i, j))
                row.append(ord(c) - ord('a'))
            grid.append(row)
    
    return grid, starts, end

def find_path(grid, start, end):
    # BFS ?
    
    visited = set()

    # each ele in q is [i,j, moves_so_far]
    q = [[start[0], start[1], 0]]

    while len(q) > 0:
        n = q.pop(0)

        i, j = n[0], n[1]

        if (i,j) in visited:
            continue

        visited.add((i,j))
        
        n_val = grid[i][j]

        next_coords = []
        if i == 0:
            next_coords.append((i+1, j))
        elif i == len(grid) - 1:
            next_coords.append((i-1, j))
        else:
            next_coords.append((i+1, j))
            next_coords.append((i-1, j))
        if j == 0:
            next_coords.append((i, j+1))
        elif j == len(grid[0]) - 1:
            next_coords.append((i, j-1))
        else:
            next_coords.append((i, j+1))
            next_coords.append((i, j-1))

        # check the possible options
        for coord in next_coords:
            if grid[coord[0]][coord[1]] <= n_val + 1:
                if coord[0] == end[0] and coord[1] == end[1]:
                    # we've hit the end, don't really need to search any further
                    # since we're doing BFS, the first to hit is guaranteed to be the smallest
                    return n[2] + 1

                # its a valid square to check out
                q.append([coord[0], coord[1], n[2] + 1])
    
    # no psosible way to make it
    return float('inf')
        

grid, starts, end = parse_input("input.txt")

print("Part 1:", find_path(grid, starts[0], end))

print("searching across", len(starts), "starting points")
steps_from_a = [find_path(grid, s, end) for s in starts]
print("Part 2:", min(steps_from_a))