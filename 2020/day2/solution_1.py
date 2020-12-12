inputData = []

# lines have format min-max letter: pass
# a password is valid iff letter appears between min max times (inclusive)
def parse_line(line: str) -> dict:
    result = dict()

    parts = line.strip().split(": ")
    result["password"] = parts[1]

    max_min, letter = parts[0].split(" ")
    result["letter"] = letter

    min_int, max_int = max_min.split("-")
    result["min"] = int(min_int)
    result["max"] = int(max_int)
    return result


def is_password_valid(password: str, letter: str, min: int, max: int) -> bool:
    letter_count = sum([1 if l == letter else 0 for l in password])
    return min <= letter_count <= max


valid = 0
with open("input.txt", "r") as f:
    for line in f:
        entry = parse_line(line)
        if is_password_valid(
            entry["password"], entry["letter"], entry["min"], entry["max"]
        ):
            valid += 1
print(valid)
