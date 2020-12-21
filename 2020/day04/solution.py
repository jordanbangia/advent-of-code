required_fields = [
    "byr",
    "iyr",
    "eyr",
    "hgt",
    "hcl",
    "ecl",
    "pid",
]

valid_eye_colors = ["amb", "blu", "brn", "gry", "grn", "hzl", "oth"]
valid_hair_colors = [
    "0",
    "1",
    "2",
    "3",
    "4",
    "5",
    "6",
    "7",
    "8",
    "9",
    "a",
    "b",
    "c",
    "d",
    "e",
    "f",
]


def parse_file_for_passports() -> []:
    entry = {}

    entries = []
    with open("input.txt", "r") as f:
        for line in f:
            line = line.strip()
            if len(line) == 0:
                entries += [entry]
                entry = {}
                continue

            pairs = line.split(" ")
            for p in pairs:
                p = p.strip()
                key, value = p.split(":")
                entry[key.strip()] = value.strip()
    entries += [entry]
    return entries


def is_valid_year(year: str, min_year: int, max_year: int) -> bool:
    if len(year) != 4:
        return False
    return min_year <= int(year) <= max_year


def is_valid_height(height: str) -> bool:
    if "cm" in height:
        height_in_cm = int(height.strip("cm"))
        return 150 <= height_in_cm <= 193
    elif "in" in height:
        height_in_inches = int(height.strip("in"))
        return 59 <= height_in_inches <= 76
    return False


def is_valid_hair_color(hair_color: str) -> bool:
    if len(hair_color) != 7:
        return False
    if hair_color[0] != "#":
        return False

    hair_color = hair_color.lstrip("#")
    for l in hair_color:
        if l not in valid_hair_colors:
            return False
    return True


def is_valid_eye_color(eye_color: str) -> bool:
    return eye_color in valid_eye_colors


def is_valid_passport_id(passport_id: str) -> bool:
    if len(passport_id) != 9:
        return False
    return passport_id.isnumeric()


def is_valid_entry(entry: dict) -> bool:
    for field in required_fields:
        if field not in entry:
            return False

    # thse are for solution 2
    if not is_valid_year(entry["byr"], 1920, 2002):
        print("birth year is invalid", entry)
        return False

    if not is_valid_year(entry["iyr"], 2010, 2020):
        print("issue year is invalid", entry)
        return False

    if not is_valid_year(entry["eyr"], 2020, 2030):
        print("expiry year is invalid", entry)
        return False

    if not is_valid_height(entry["hgt"]):
        print("height is invalid", entry)
        return False

    if not is_valid_hair_color(entry["hcl"]):
        print("hair colour is invalid", entry)
        return False

    if not is_valid_eye_color(entry["ecl"]):
        print("eye colour is invalid", entry)
        return False

    if not is_valid_passport_id(entry["pid"]):
        print("passport id is invalid", entry)
        return False

    return True


entries = parse_file_for_passports()
valid = 0
for entry in entries:
    if is_valid_entry(entry):
        valid += 1
print(valid)