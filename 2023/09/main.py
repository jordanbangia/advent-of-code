
def parse_input(i):
    with open(f"{i}.txt", "r") as f:
        return [
            [int(a) for a in l.strip().split()] for l in f.readlines()
        ]


def _extrapolate_forward(v):

    diffs = [v]
    while not all(a == 0 for a in diffs[-1]):
        d = diffs[-1]
        n = [d[i] - d[i-1] for i in range(1, len(d))]
        diffs.append(n)
        print(diffs)
    
    diffs = list(reversed(diffs))

    # once we've hit all zeroes, we work backwards

    for i, d in enumerate(diffs):
        if i == 0:
            d.append(0)
        else:
            below = diffs[i-1]
            d.append(d[-1] + below[-1])
    print(diffs[-1][-1])
    return diffs[-1][-1]



def part_1(i):
    history_values = parse_input(i)
    return sum(
        _extrapolate_forward(h) for h in history_values
    )

def _extrapolate_backwards(v):
    diffs = [v]
    while not all(a == 0 for a in diffs[-1]):
        d = diffs[-1]
        n = [d[i] - d[i-1] for i in range(1, len(d))]
        diffs.append(n)
        print(diffs)

    diffs = list(reversed(diffs))
    for i, d in enumerate(diffs):
        if i == 0:
            d.insert(0, 0)
        else:
            d.insert(0, d[0] - diffs[i-1][0])
    print(diffs[-1][0])
    return diffs[-1][0]


def part_2(i):
    history_values = parse_input(i)
    return sum(
        _extrapolate_backwards(h) for h in history_values
    )

# print(part_1("input"))
print(part_2("input"))