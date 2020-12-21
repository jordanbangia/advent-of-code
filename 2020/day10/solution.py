from collections import defaultdict

data = []
with open("input.txt", "r") as f:
    for line in f:
        line = line.strip()
        data += [int(line)]
data.extend([0, max(data) + 3])
data = sorted(data)

print(data)


def solve_first():
    current_joltage = 0
    diffs = {1: 0, 2: 0, 3: 0}

    for _, rating in enumerate(data[1:]):
        diff = rating - current_joltage
        if diff > 3:
            print("SOMETHING IS WRONG")
        diffs[diff] += 1
        current_joltage = rating

    print(current_joltage)
    print(diffs)
    print(diffs[1] * diffs[3])


cache = defaultdict(int, {0: 1})


# we make a cache that stores the number of ways to get the given voltage
# 0 (plug) is 1
# 1 requires 1
# 2 can be done with 2
# each cache entrry is the sum of the cache entries for the given joltage value - 1 2 and 3
def solve_second():
    for i in data[1:]:
        cache[i] = sum(cache[i] for i in range(i - 3, i))
    # print(cache)
    print(cache[data[-1]])


solve_first()
solve_second()