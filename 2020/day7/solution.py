import re

with open("input.txt", "r") as f:
    lines = f.read().splitlines()

bags = {}
bag_count = 0

# make map of colour to bags it contains
for line in lines:
    color = re.match(r"(.+?) bags contain", line)[1]
    bags[color] = re.findall(r"(\d+?) (.+?) bags?", line)


def has_shiny_gold(colour):
    if colour == "shiny gold":
        return True
    else:
        # recursively try sub bags
        return any(has_shiny_gold(c) for _, c in bags[colour])


for bag in bags:
    if has_shiny_gold(bag):
        bag_count += 1

print("Part 1: " + str(bag_count - 1))
print(bags["clear tan"])


def count_bags(bag_type, level=0):
    print(level, bag_type)
    return 1 + sum(
        int(number) * count_bags(colour, level + 1) for number, colour in bags[bag_type]
    )


print("Part 2: " + str(count_bags("shiny gold") - 1))
