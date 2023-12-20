
from collections import defaultdict


def parse_input(i):
    with open(f"{i}.txt", "r") as f:
        return [line.strip() for line in f.readlines()]
    
    
def count_rock_positions(g: list[str]):
    num_rows = len(g)
    num_cols = len(g[0])

    col_rock_pairs = []

    for i in range(num_cols):
        stopping_points = []

        still_rock = num_rows+1
        moving_rocks = 0
        for k in range(num_rows):
            if g[k][i] == "O":
                moving_rocks += 1
            elif g[k][i] == ".":
                continue
            elif g[k][i] == "#":
                if moving_rocks != 0:
                    stopping_points.append(
                        (still_rock, moving_rocks)
                    )
                moving_rocks = 0
                still_rock = num_rows - k
        if moving_rocks != 0:
            stopping_points.append(
                (still_rock, moving_rocks)
            )
        col_rock_pairs.append(stopping_points)
    

    load = 0
    for col_pairs in col_rock_pairs:
        for still_rock, num_rocks in col_pairs:
            p = [
                still_rock - r for r in range(1, num_rocks + 1)
            ]
            load += sum(p)
    
    return load

def part_1(i):
    g = parse_input(i)
    print(count_rock_positions(g))


def roll_north(g):
    # count down the row, and place them after each rock
    num_rows = len(g)
    num_cols = len(g[0])

    for i in range(num_cols):
        stopping_points = []

        still_rock = -1
        moving_rocks = 0
        for k in range(num_rows):
            if g[k][i] == "O":
                moving_rocks += 1
                g[k][i] = '.'
            elif g[k][i] == ".":
                continue
            elif g[k][i] == "#":
                if moving_rocks != 0:
                    stopping_points.append(
                        (still_rock, moving_rocks)
                    )
                moving_rocks = 0
                still_rock = k
        if moving_rocks != 0:
            stopping_points.append(
                (still_rock, moving_rocks)
            )
        
        for still_rock, moving_rocks in stopping_points:
            for x in range(1, moving_rocks + 1):
                if g[still_rock+x][i] not in ('O', '.'):
                    raise ValueError("WTF")
                g[still_rock+x][i] = "O"
    return g

def roll_west(g):
    # count across the row from left to right
    # place them at the rock
    num_rows = len(g)
    num_cols = len(g[0])

    for i in range(num_rows):
        stopping_points = []

        still_rock = -1
        moving_rocks = 0
        for k in range(num_cols):
            if g[i][k] == "O":
                moving_rocks += 1
                g[i][k] = '.'
            elif g[i][k] == ".":
                continue
            elif g[i][k] == "#":
                if moving_rocks != 0:
                    stopping_points.append(
                        (still_rock, moving_rocks)
                    )
                moving_rocks = 0
                still_rock = k
        if moving_rocks != 0:
            stopping_points.append(
                (still_rock, moving_rocks)
            )
        
        for still_rock, moving_rocks in stopping_points:
            for x in range(1, moving_rocks + 1):
                if g[i][still_rock + x] not in ('O', '.'):
                    raise ValueError("WTF")
                g[i][still_rock + x] = "O"
    return g


def roll_south(g):
    # count up the row
    num_rows = len(g)
    num_cols = len(g[0])
    
    for i in range(num_cols):
        stopping_points = []

        still_rock = num_rows
        moving_rocks = 0
        for k in range(num_rows - 1, -1, -1):
            if g[k][i] == 'O':
                g[k][i] = '.'
                moving_rocks += 1
            elif g[k][i] == '.':
                continue
            elif g[k][i] == '#':
                if moving_rocks != 0:
                    stopping_points.append((still_rock, moving_rocks))
                moving_rocks = 0
                still_rock = k
        if moving_rocks != 0:
            stopping_points.append((still_rock, moving_rocks))

        for still_rock, moving_rocks in stopping_points:
            for x in range(1, moving_rocks+1):
                g[still_rock - x][i] = 'O'
    
    return g


def roll_east(g):
    # count across the row from right to left
    # place them at the rock
    num_rows = len(g)
    num_cols = len(g[0])

    for i in range(num_rows):
        stopping_points = []

        still_rock = num_cols
        moving_rocks = 0
        for k in range(num_cols-1, -1, -1):
            if g[i][k] == "O":
                moving_rocks += 1
                g[i][k] = '.'
            elif g[i][k] == ".":
                continue
            elif g[i][k] == "#":
                if moving_rocks != 0:
                    stopping_points.append(
                        (still_rock, moving_rocks)
                    )
                moving_rocks = 0
                still_rock = k
        if moving_rocks != 0:
            stopping_points.append(
                (still_rock, moving_rocks)
            )
        
        for still_rock, moving_rocks in stopping_points:
            for x in range(1, moving_rocks + 1):
                if g[i][still_rock - x] not in ('O', '.'):
                    raise ValueError("WTF")
                g[i][still_rock - x] = "O"
    return g

def count_load(g):
    load = 0
    max_l = len(g)
    for i in range(len(g[0])):
        for k in range(len(g)):
            if g[k][i] == 'O':
                load += max_l - k
    return load

def part_2(i):
    g = parse_input(i)
    g = [[c for c in s] for s in g]

    states = {}
    
    # it appears once we hit a stable point, we'll just continue hitting
    # that stable point over and over again
    # so lets assume once we hit it, we're done and can stop
    jumped = False
    i = 1
    j = 1_000_000_000
    while j > 0:
        print(f"After {i} cylce")
        g = roll_north(g)
        g = roll_west(g)
        g = roll_south(g)
        g = roll_east(g)
        
        j -= 1

        g_as_str = "\n".join(["".join(s) for s in g])
        if g_as_str in states and not jumped:
            print("I've seen state again", i, states[g_as_str])
            cycle = i - states[g_as_str]
            
            print(f"I guess I cycle every {cycle} iterations")
            remaining_its = j
            print(f"I've got {remaining_its} to go")
            
            remaining_its = remaining_its % cycle

            print(f"I'll only need to actually do the remaining {remaining_its} its")

            j = remaining_its
            jumped = True
        else:
            states[g_as_str] = i
            
        i += 1
        

    print("\n".join(["".join(s) for s in g]))
    print("\n")

    load = count_load(g)
    return load


# print(part_1("input"))
print(part_2("input"))

