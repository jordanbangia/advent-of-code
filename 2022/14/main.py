
from collections import defaultdict


class KeyDefaultDict(defaultdict):
    floor = None

    def set_floor(self, floor):
        self.floor = floor

    def __missing__(self, key):
        if self.floor and key[1] == self.floor:
            ret = self[key] = '#'
            return ret
        else:
            ret = self[key] = '.'
            return ret

def parse_input(file_name: str) -> dict[str]:
    d = KeyDefaultDict()
    
    with open(file_name, "r") as f:
        for line in f.readlines():
            line = line.strip()

            parts = line.split(" -> ")
            for i in range(len(parts)-1):
                coord1 = parts[i].split(',')
                coord2 = parts[i+1].split(',')

                coord1 = (int(coord1[0]), int(coord1[1]))
                coord2 = (int(coord2[0]), int(coord2[1]))

                if coord1[0] == coord2[0]:
                    start = min(coord1[1], coord2[1])
                    end = max(coord1[1], coord2[1])
                    for i in range(start, end+1):
                        d[(coord1[0], i)] = '#'
                else:
                    start = min(coord1[0], coord2[0])
                    end = max(coord1[0], coord2[0])
                    for i in range(start, end+1):
                        d[(i, coord1[1])] = '#'
    
    return d

def print_grid(d, floor=None):
    if floor is None:
        floor = 10

    z = [r for r, _ in d.keys()]

    for j in range(floor+1):
        print(''.join(d[(x,j)] for x in range(min(z), max(z) + 1)))


def part_1(file_name: str):
    print("Part 1:")
    d = parse_input(file_name)
    lowest_rock = max(c for _, c in d.keys())
    print("lowest rock is at:", lowest_rock)

    came_to_rest = 0
    sand_into_abyss = False
    while not sand_into_abyss:
        # sand starts at (500, 0)
        curr_pos = (500, 0)
        at_rest = False
        while not at_rest:
            if curr_pos[1] > lowest_rock:
                # we've passed into the abyss
                sand_into_abyss = True
                break

            next_pos = (curr_pos[0], curr_pos[1] + 1)
            if d[next_pos] == '.':
                # next space is empty, can tick down
                curr_pos = next_pos
                continue
            
            # its a blocked, so check down to the left
            next_pos = (curr_pos[0]-1, curr_pos[1] +1)
            if d[next_pos] == '.':
                curr_pos = next_pos
                continue
            
            next_pos = (curr_pos[0] + 1, curr_pos[1] + 1)
            if d[next_pos] == '.':
                curr_pos = next_pos
                continue
            
            # couldn't go striaght down, down the left or down to the right
            # so we must be coming to a rest
            d[curr_pos] = 'o'
            at_rest = True
            came_to_rest += 1
    print("units came to rest:", came_to_rest)


def part_2(file_name: str):
    print("Part 2:")
    d = parse_input(file_name)
    lowest_rock = max(c for _, c in d.keys())
    print("lowest rock is at:", lowest_rock)
    
    floor = lowest_rock + 2
    print("floor is at:", floor)
    
    d.set_floor(floor)
    
    came_to_rest = 0
    while d[(500, 0)] != 'o':
        # sand starts at (500, 0)
        curr_pos = (500, 0)
        at_rest = False
        while not at_rest:
            next_pos = (curr_pos[0], curr_pos[1] + 1)
            if d[next_pos] == '.':
                # next space is empty, can tick down
                curr_pos = next_pos
                continue
            
            # its a blocked, so check down to the left
            next_pos = (curr_pos[0]-1, curr_pos[1] +1)
            if d[next_pos] == '.':
                curr_pos = next_pos
                continue
            
            next_pos = (curr_pos[0] + 1, curr_pos[1] + 1)
            if d[next_pos] == '.':
                curr_pos = next_pos
                continue
            
            # couldn't go striaght down, down the left or down to the right
            # so we must be coming to a rest
            d[curr_pos] = 'o'
            at_rest = True
            came_to_rest += 1
        # if came_to_rest == 93:
        #     print_grid(d, floor)
        #     return
    # print_grid(d)
    print("units came to rest:", came_to_rest)


test_file = "input.txt"
part_1(test_file)
part_2(test_file)