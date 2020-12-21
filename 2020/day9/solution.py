data = []
with open("input.txt", "r") as f:
    for line in f:
        line = line.strip()
        data += [int(line)]

preamble_length = 25


def valid_next_values(preamble: []) -> set:
    valids = set()
    for i in range(len(preamble)):
        for y in range(i + 1, len(preamble)):
            valids.add(preamble[i] + preamble[y])
    return valids


invalid = 0
for i in range(preamble_length, len(data)):
    preamble_start = i - preamble_length
    valids = valid_next_values(data[preamble_start : preamble_start + preamble_length])
    if data[i] not in valids:
        print("first data point that is not valid:", data[i])
        invalid = data[i]
        break


start = 0
end = 0

while sum(data[start:end]) != invalid:
    current_total = sum(data[start:end])
    if current_total > invalid:
        start += 1
    elif current_total < invalid:
        end += 1

print("segment starts at ", start, "and ends at ", end)

print(data[start:end])
print(sum(data[start:end]))
print(min(data[start:end]) + max(data[start:end]))