from collections import Counter


def parse_input_file() -> []:
    groups = []
    group = []
    with open("input.txt", "r") as f:
        for line in f:
            line = line.strip()
            if len(line) == 0:
                groups += [group]
                group = []
                continue

            group += [line]
    groups += [group]
    return groups


def count_unique_responses_in_group(group: []) -> int:
    people_in_group = len(group)

    c = Counter()
    for person in group:
        c.update(person)

    all_answered_count = 0
    for question, answer_count in c.items():
        if answer_count == people_in_group:
            all_answered_count += 1
    print(c, all_answered_count)
    return all_answered_count


groups = parse_input_file()

total_answers = 0
for group in groups:
    group_count = count_unique_responses_in_group(group)
    total_answers += group_count

print(total_answers)
