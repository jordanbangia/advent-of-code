inputData = []

# lines have format first-second letter: pass
# a password is valid iff letter appears at first index or second index
# its invalid if it appears at both


def parse_line(line: str) -> dict:
    result = dict()

    parts = line.strip().split(": ")
    result["password"] = parts[1]

    max_min, letter = parts[0].split(" ")
    result["letter"] = letter

    min_int, max_int = max_min.split("-")
    result["first"] = int(min_int)
    result["second"] = int(max_int)
    return result


def is_password_valid(password: str, letter: str, first: int, second: int) -> bool:
    first_is_expected = password[first - 1] == letter
    second_is_expected = password[second - 1] == letter
    return first_is_expected ^ second_is_expected


valid = 0
with open("input.txt", "r") as f:
    for line in f:
        entry = parse_line(line)
        if is_password_valid(
            entry["password"], entry["letter"], entry["first"], entry["second"]
        ):
            valid += 1
print(valid)
