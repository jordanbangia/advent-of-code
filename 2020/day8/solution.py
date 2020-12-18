instr = []
with open("input.txt", "r") as f:
    for line in f:
        line = line.strip()
        instruction = line.split(" ")
        instr += [(instruction[0], int(instruction[1]))]

accumulator = 0

# dict of instruction idex to bool of seen.  If not in dict its unseen
seen = dict()


def run(code):
    accumulator = 0

    i = 0

    seen = set()

    while True:
        if i == len(code):
            return 0, accumulator
        if i in seen:
            return 1, accumulator
        seen.add(i)
        instr, arg = code[i]
        if instr == "acc":
            accumulator += arg
            i += 1
        elif instr == "jmp":
            i += arg
        elif instr == "nop":
            i += 1


def solve_second(code):
    for i, (instr, arg) in enumerate(code):
        if instr in ["jmp", "nop"]:
            if instr == "jmp":
                code[i] = ("nop", arg)
            elif instr == "nop":
                code[i] = ("jmp", arg)
            retcode, retval = run(code)
            if retcode == 0:
                print("at index", i)
                return retval
            else:
                code[i] = instr, arg


print("loop found, accumlator is:", run(instr))

print(solve_second(instr))