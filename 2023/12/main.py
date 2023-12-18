


from functools import cache


def parse_input(i):
    rows = []
    with open(f"{i}.txt", "r") as f:
        for line in f.readlines():
            parts = line.strip().split(" ")
            springs = parts[0]
            groups = tuple(int(j) for j in parts[1].strip().split(","))
            rows.append((springs, groups))
    return rows

@cache
def compute_arrangments(springs, groups):
    if not groups:
        return '#' not in springs
    
    # how many springs we have in total
    springs_len = len(springs)

    # current group that we're evaluating
    group_len = groups[0]

    # if the sum of all the groups (i.e all the broken springs) 
    # + the number of broken groups(which would tell us how many operationsl ones need to exist)
    # is greater than all the springs, then we've hit a condition that doesn't work
    # so can return 0 from here
    if springs_len - sum(groups) - len(groups) + 1 < 0:
        return 0
    
    # check if there are any operational springs in the current group
    # of interest
    has_holes = any(springs[x] == '.' for x in range(group_len))
    
    # if the remaining number of spings == the length of this group
    if springs_len == group_len:
        # we return 1 way to make this if all springs are broken
        # otherwise we return 0 since there's an operational spring
        # breaking the chain
        return 0 if has_holes else 1
    
    # springs in group_len if there are no holes (i.e. there are only # or ? springs)
    # and if the spring right after the group isn't broken, as that would make
    # the group count a lie
    can_use = not has_holes and (springs[group_len] != '#')

    # if the spring at the beginning is broken
    if springs[0] == '#':
        # then we have to use this character ina  group
        # assume that we take all the springs required to fill out this group
        # and remove any leading operational springs as they aren't useful
        # and then reduce the groups required by 1 since we've filled it
        # all this assuming that this sequence is usable, has no holes and the
        # next spring isn't broken
        return compute_arrangments(
            springs[group_len+1:].lstrip('.'), tuple(groups[1:])
        ) if can_use else 0
    
    # from here, the first character is '?'
    # first we calculate based on the condition that we don't use the '?' in the group
    skip = compute_arrangments(
        springs[1:].lstrip('.'), groups
    )
    
    # if the sequence of character's are unusable for the group, then we just return skip
    if not can_use:
        return skip
    
    # if it is usable, we return skip + the count assuming we fill the group
    return skip + compute_arrangments(
        springs[group_len+1:].lstrip('.'), tuple(groups[1:])
    )



def part_1(i):
    t = 0
    for springs, groups in parse_input(i):
        t += compute_arrangments(springs, groups)
    return t

def part_2(i):
    t = 0
    for springs, groups in parse_input(i):
        springs = "?".join([springs]*5).lstrip(".")
        groups = tuple([int(g) for g in groups] * 5)
        t += compute_arrangments(springs, groups)
    return t

print(part_1("input"))
print(part_2("input"))