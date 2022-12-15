import ast
from copy import deepcopy
import itertools

def parse_input(file_name: str) -> list[list]:
    input_pairs = []
    with open(file_name, "r") as f:

        curr = []
    
        for line in f.readlines():
            line = line.strip()
            if line == "":
                input_pairs.append(curr)
                curr = []
            else:
                curr.append(ast.literal_eval(line))
    input_pairs.append(curr)
    return input_pairs

def is_pair_in_right_order(_left: list, _right: list) -> bool:
    left = deepcopy(_left)
    right = deepcopy(_right)
    a = 0
    b = 0

    if len(left) == 0 and len(right) > 0:
        return True
    elif len(left) > 0 and len(right) == 0:
        return False
    
    count = 0
    while a < len(left) and b < len(right):
        count += 1
        if count > 10:
            return None
        left_val = left[a]
        right_val = right[b]

        # print(left_val, "vs", right_val)

        if isinstance(left_val, int) and isinstance(right_val, int):
            if left_val == right_val:
                a += 1
                b += 1
            else:
                return left_val < right_val
        
        elif isinstance(left_val, list) and isinstance(right_val, list):
            sub_results = is_pair_in_right_order(left_val, right_val)
            if sub_results is not None:
                return sub_results
            a += 1
            b += 1

        else:
            if isinstance(left_val, list):
                right[b] = [right_val]
            else:
                left[a] = [left_val]

    if a == len(left) and b == len(right):
        return None
    return a == len(left) and b < len(right)


count = 0
it = parse_input("input.txt")
# for x, i in enumerate(it):
#     res = is_pair_in_right_order(i[0], i[1])
#     print(f"{x+1}: {res}")
#     count += (x+1) if res else 0
# print("Part 1:", count)

all_packets = [[[2]], [[6]]]
for l, r in it:
    all_packets.append(l)
    all_packets.append(r)

# print(all_packets)

swapped = True
while swapped:
    # print(">>> next it")
    swapped = False
    for i in range(len(all_packets) - 1):
        l = all_packets[i]
        r = all_packets[i+1]
        res = is_pair_in_right_order(l, r)
        # print(l, "vs", r, res)
        if not res:
            swapped = True
            all_packets[i], all_packets[i+1] = all_packets[i+1], all_packets[i]

decoder_key = 1
for i, p in enumerate(all_packets):
    # print(p)
    if p == [[2]] or p == [[6]]:
        decoder_key = decoder_key * (i + 1)
print("Part 2:", decoder_key)