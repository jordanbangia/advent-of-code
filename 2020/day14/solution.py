def read_input():
    data = []
    with open("input.txt", "r") as f:
        for line in f:
            arg, val = line.strip().split(" = ")
            if arg == "mask":
                data += [(arg, val)]
            else:
                data += [(int(arg[4:-1]), val)]
    return data


def bitstring(x, len):
    return bin(x)[2:].zfill(len)


def solve_one(code):
    mem = {}
    mask = ""
    for instr, val in code:
        if instr == "mask":
            mask = val
            continue

        value = bitstring(int(val), len(mask))
        # copy the value in X, or use the value from mask
        value = (v if m == "X" else m for v, m in zip(value, mask))

        value = int("".join(value), 2)
        mem[instr] = value
    return sum(mem.values())


def solve_two(code):
    mem = {}
    mask = ""
    for instr, val in code:
        if instr == "mask":
            mask = val
            continue

        addr = bitstring(instr, len(mask))
        addr_template = ""
        for mask_bit, addr_bit in zip(mask, addr):
            if mask_bit == "0":
                addr_template += addr_bit
            elif mask_bit == "1":
                addr_template += "1"
            else:
                addr_template += "{}"

        floating_length = mask.count("X")
        for f in range(2 ** floating_length):  # 2^num xs possible new values
            addr = addr_template.format(*bitstring(f, floating_length))
            addr = int(addr, 2)
            mem[addr] = int(val)
    return sum(mem.values())


code = read_input()
print("first answer", solve_one(code))
print("second answer", solve_two(code))