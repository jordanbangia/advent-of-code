"""
This feels similar to one of the earlier problems
The plan:
- I will create the grid that they are considering
- I will then walk across each row, and add up all the elements st. that
    - if num walls seen is even, they don't count
    - if num walls seen is odd, they do count


It's actually not that, the original question is an easier extension of this one
This uses something called the shoelace theoreum for calculating the area of an
arbitrary closed polygon
Hell I could have punched this in shapely and asked for its area
"""


directions = {
    "R": (1, 0),
    "L": (-1, 0),
    "U": (0, -1),
    "D": (0, 1),
}


def calculate_fill_using_shoe_theorum(instructions):
    current = (0, 0)
    trench = [current]
    perimeter = 0
    for instruction in instructions:
        direction, length = instruction
        d = directions[direction]
        current = (
            current[0] + length * d[0],
            current[1] + length * d[1],
        )
        trench.append(current)
        perimeter += length

    print(trench)
    print(perimeter)

    t = perimeter + len(trench) - 1
    print(t)
    t = (perimeter / 2) + 1
    print(t)
    for i in range(1, len(trench)):
        trig = (trench[i - 1][0] * trench[i][1] - trench[i - 1][1] * trench[i][0]) / 2
        print(i, trig)
        t += trig

    return t


def part_1(i):
    instructions = []

    with open(f"{i}.txt", "r") as f:
        for line in f.readlines():
            parts = line.strip().split(" ")

            direction = parts[0]
            length = int(parts[1])

            instructions.append((direction, length))

    return calculate_fill_using_shoe_theorum(instructions)


def part_2(i):
    instructions = []

    direction_digit = {
        "0": "R",
        "1": "D",
        "2": "L",
        "3": "U",
    }

    with open(f"{i}.txt", "r") as f:
        for line in f.readlines():
            parts = line.strip().split(" ")

            hex_code = (
                parts[2].strip().replace("#", "").replace("(", "").replace(")", "")
            )
            direction = direction_digit[hex_code[-1]]

            length = int(hex_code[:5], 16)

            print(direction, length)

            instructions.append((direction, length))

    return calculate_fill_using_shoe_theorum(instructions)


# print(part_1("input"))
print(part_2("input"))
