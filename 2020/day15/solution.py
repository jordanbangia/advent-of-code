from collections import defaultdict

input_txt = [20, 0, 1, 11, 6, 3]
# input_txt = [0, 3, 6]


def solution(turns):
    turn = 1

    # each entry in memory is an array of the last times that a number was seen
    memory = dict()

    for i in input_txt:
        if i not in memory:
            memory[i] = []
        memory[i].insert(0, turn)
        turn += 1

    said = [i for i in input_txt]

    for i in range(turn, turns + 1):
        last_spoken = said[-1]
        if len(memory[last_spoken]) == 1:
            # first time last spoken was spoken
            said += [0]
            memory[0].insert(0, i)
        else:
            # this number was been said more than once
            next_spoken = memory[last_spoken][0] - memory[last_spoken][1]
            said += [next_spoken]
            if next_spoken not in memory:
                memory[next_spoken] = []
            memory[next_spoken].insert(0, i)

    print("solution for turns", turns, ":", said[-1])


def solve_for_fast():
    mem = {n: i for i, n in enumerate(input_txt[:-1], 1)}
    mem = defaultdict(bool, mem)
    last = input_txt[-1]
    print(mem)
    for turn in range(len(input_txt), 30000000):
        current = mem[last]
        current = turn - current if current else 0
        mem[last] = turn
        last = current
    return current


solution(2020)
print(solve_for_fast())
