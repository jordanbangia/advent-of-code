
from collections import defaultdict


def parse_input(i):
    g = []
    with open(f'{i}.txt', 'r') as f:
        for l in f.readlines():
            g.append(
                [c for c in l.strip()]
            )
    return g


def check_for_symbol(r, c, g) -> (bool, list[tuple[int, int]]):
    has_symbol = False
    gears = []
    for i in [-1, 0 , 1]:
        for j in [-1, 0, 1]:
            if i == j == 0:
                continue
            
            r_p = r + i
            c_p = c + j

            if r_p <= -1 or r_p >= len(g):
                continue
            if c_p <= -1 or c_p >= len(g[0]):
                continue
        
            if g[r_p][c_p] != '.' and not g[r_p][c_p].isdigit():
                has_symbol = True
                if g[r_p][c_p] == '*':
                    gears.append((r_p, c_p))
    return has_symbol, gears

def part(i) -> tuple[int, int]:
    g = parse_input(i)

    part_num_sum = 0

    # every time we find a num adjacent to a gear,
    # store that in this dict with the coordinate
    gear_nums = defaultdict(list)

    for i, row in enumerate(g):
        n = ''
        is_part_num = False
        gear_symbols = []
        for j, c in enumerate(row):
            if c.isdigit():
                n += c
                is_next_to_sybmol, symbol_locations = check_for_symbol(i, j, g)
                
                is_part_num = is_part_num or is_next_to_sybmol
                gear_symbols.extend(symbol_locations)
    
            elif n != '' and is_part_num:
                print("found part", int(n))
                part_num_sum += int(n)

                if len(gear_symbols) > 0:
                    gears = set(gear_symbols)
                    for gear in gears:
                        gear_nums[gear].append(int(n))

                
                n = ''
                is_part_num = False
                gear_symbols = []
            elif n != '':
                print("not including", n)
                
                n = ''
                is_part_num = False
                gear_symbols = []
            else:
                # we never started a number, so keep moving
                continue
        
        if n != '':
            if is_part_num:
                print("found part", n)
                part_num_sum += int(n)
                if len(gear_symbols) > 0:
                    gears = set(gear_symbols)
                    for gear in gears:
                        gear_nums[gear].append(int(n))
            else:
                print("not including", n)
            

    gear_ratios_prod = 0
    for gear, nums in gear_nums.items():
        if len(nums) != 2:
            continue
        print(gear, nums)
        gear_ratios_prod += nums[0] * nums[1]



    return (part_num_sum, gear_ratios_prod)


print(part('part_1'))