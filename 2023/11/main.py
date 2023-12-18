
def parse_input(i):    
    with open(f"{i}.txt", "r") as f:
        universe = [line.strip() for line in f.readlines()]
        return universe
    

def expand_universe(universe: list[str]):
    # first expand every row:
    expanded_universe = []
    for l in universe:
        if len(set(l)) == 1 and l[0] == '.':
            expanded_universe.append(l)
        expanded_universe.append(l)

    # next expand the columns
    rotated_universe = []
    for i in range(len(expanded_universe[0])):
        c = [g[i] for g in expanded_universe]
        if len(set(c)) == 1 and c[0] == '.':
            rotated_universe.append(c)
        rotated_universe.append(c)

    # rotate it back for convenience
    rotated_back = []
    for i in range(len(rotated_universe[0])):
        rotated_back.append(''.join([g[i] for g in rotated_universe]))
    return rotated_back

def print_universe(universe: list[str]):
    for g in universe:
        print(g)

def find_galaxy_locations(universe: list[str]):
    galaxies = {}
    x = 1
    for i, g in enumerate(universe):
        for j, c in enumerate(g):
            if c == '#':
                galaxies[x] = (i,j)
                x += 1
    return galaxies


movement_options = [
    (1, 0),
    (-1, 0),
    (0, 1),
    (0, -1),
]


def part_1(i):
    universe = expand_universe(parse_input(i))
    # print_universe(universe)

    galaxies = find_galaxy_locations(universe)
    # print(galaxies)

    galaxy_ids = list(galaxies.keys())

    print(len(galaxy_ids))

    steps = []
    for i in range(len(galaxy_ids)):
        for k in range(i+1, len(galaxy_ids)):

            galaxy_i = galaxies[galaxy_ids[i]]
            galaxy_k = galaxies[galaxy_ids[k]]

            st = abs(galaxy_i[0] - galaxy_k[0]) + abs(galaxy_i[1] - galaxy_k[1])

            print(f"From {galaxy_ids[i]} to {galaxy_ids[k]} takes {st}")

            steps.append(
                st
            )       
    return sum(steps)


def find_expansion_rows_and_cols(universe: list[str]) -> tuple[list[int], list[int]]:
    # we want to encode in the universe that it you traverse across it in certain ways
    # you hit the expansions
    # we don't want to add all the rows for expansions, we just want to track that they exist

    # if row i is an expansion universe, traversing across i is the same as traversing across expansion factor i
    # so lets just track which rows are "expanded" rows and which columns are expanded

    expanded_rows = [
        i for i, l in enumerate(universe) if len(set(l)) == 1 and l[0] == '.'
    ]

    expanded_cols = []
    for i in range(len(universe[0])):
        if len({c[i] for c in universe}) == 1 and universe[0][i] == '.':
            expanded_cols.append(i)

    return expanded_rows, expanded_cols


def part_2(i, expansion_factor: int):
    universe = parse_input(i)
    expansion_rows, expansion_cols = find_expansion_rows_and_cols(universe)

    print(len(expansion_rows), len(expansion_cols))
    
    galaxies = find_galaxy_locations(universe)
    galaxy_ids = list(galaxies.keys())
    print(len(galaxy_ids))

    steps_total = 0
    
    for i in range(len(galaxy_ids)):
        for k in range(i+1, len(galaxy_ids)):
            galaxy_i = galaxy_ids[i]
            galaxy_k = galaxy_ids[k]

            gir, gic = galaxies[galaxy_i]
            gkr, gkc = galaxies[galaxy_k]

            row_traversal = abs(gkr - gir)
            expansion_row_traversals = len([x for x in expansion_rows if min(gkr, gir) <= x <= max(gkr, gir)])

            rows_traversed = row_traversal - expansion_row_traversals + expansion_row_traversals*expansion_factor

            col_traversal = abs(gkc - gic)
            expansion_col_traversals = len([x for x in expansion_cols if min(gkc, gic) <= x <= max(gkc, gic)])

            cols_traversed = col_traversal - expansion_col_traversals + expansion_col_traversals * expansion_factor

            steps = cols_traversed + rows_traversed

            # print(f"Steps from {galaxy_i} to {galaxy_k} = {steps}")

            steps_total += steps
    
    return steps_total
            

# print(part_1("input"))
print(part_2("input", 1000000))