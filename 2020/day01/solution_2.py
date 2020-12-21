inputData = []

with open("input.txt", "r") as f:
    for line in f:
        inputData += [int(line)]

for x in inputData:
    for y in inputData:
        for z in inputData:
            if x + y + z == 2020:
                print(x, y, z, x * y * z)
