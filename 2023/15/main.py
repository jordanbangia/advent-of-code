
from collections import defaultdict


def hash_str(s: str) -> int:
    current = 0
    for c in s:
        current += ord(c)
        current *= 17
        current = current % 256
    return current

def part_1(i):
    with open(f"{i}.txt", "r") as f:
        input_str = f.readline().strip()

        t = 0
        for p in input_str.split(","):
            t += hash_str(p.strip())
        return t
    
def part_2(i):
    boxes = defaultdict(list)

    input_str = ""
    with open(f"{i}.txt", "r") as f:
        input_str = f.readline().strip()
    
    for instruction in input_str.split(","):
        instruction = instruction.strip()
        
        label = ""
        operation = ""
        focal_length = None
        if "-" in instruction:
            parts = instruction.split("-")
            label = parts[0].strip()
            operation = "-"
        else:
            parts = instruction.split("=")
            label = parts[0].strip()
            operation = "="
            focal_length = int(parts[1].strip())


        box = hash_str(label)
        indx_of_label = next((i for i in range(len(boxes[box])) if boxes[box][i][0] == label), None)

        if operation == "-":
            if indx_of_label is not None:
                boxes[box].pop(indx_of_label)
        elif operation == "=":
            if indx_of_label is not None:
                boxes[box][indx_of_label][1] = focal_length
            else:
                boxes[box].append([label, focal_length])

    focusing_power = 0
    for i, b in boxes.items():
        if len(b) == 0:
            continue
    
        for k, lens in enumerate(b):
            focusing_power += (1+i) * (1+k) * lens[1]
    return focusing_power
    
print(part_1("input"))

print(part_2("input"))