inputData = []

with open("input.txt", "r") as f:
    for line in f:
        inputData += [int(line)]

for x in inputData:
    for y in inputData:
        if x + y == 2020:
            print(x, y, x * y)
