max_rows = 128
max_columns = 8


def parse_input_file() -> []:
    out = []
    with open("input.txt", "r") as f:
        for line in f:
            line = line.strip()
            out += [line]
    return out


def input_to_seat(l: str) -> (int, int, int):
    bin_ = ""
    for i in l:
        if i == "F":
            bin_ += "0"
        elif i == "B":
            bin_ += "1"
        elif i == "L":
            bin_ += "0"
        elif i == "R":
            bin_ += "1"

    row = bin_[0:7]
    col = bin_[7:11]

    return int(row, base=2), int(col, base=2), int(bin_, base=2)


input_passes = parse_input_file()

seat_ids = []
for p in input_passes:
    row, col, seat_id = input_to_seat(p)
    seat_ids += [seat_id]

seat_ids.sort()
print(seat_ids)


for i in range(1, len(seat_ids) - 2):
    # print(seat_ids[i])
    if seat_ids[i + 1] - seat_ids[i] != 1:
        print(seat_ids[i] + 1)
