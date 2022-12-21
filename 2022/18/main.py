from collections import defaultdict

def parse_input(filename: str) -> dict:
    d = defaultdict(int)

    with open(filename, 'r') as f:
        for line in f.readlines():
            coord = tuple([int(c) for c in line.strip().split(',')])
            d[coord] = 1
    
    return d



sides = [
    (1, 0, 0),
    (-1, 0, 0),
    (0, 1, 0),
    (0, -1, 0),
    (0, 0, 1),
    (0, 0, -1),
]

def count_surface_area(d: dict):
    c = 0
    keys = list(d.keys())
    for key in keys:
        # for each point, check if there is something beside it
        for side in sides:
            coord = tuple([key[0] + side[0], key[1] + side[1], key[2] + side[2]])
            if d[coord] == 0:
                c += 1
    return c

def count_exterior_surface_area(d: dict):
    keys = list(d.keys())
    maxes = [
        max(k[0] for k in keys)+1,
        max(k[1] for k in keys)+1,
        max(k[2] for k in keys)+1
    ]

    mins = [
        min(k[0] for k in keys)-1,
        min(k[1] for k in keys)-1,
        min(k[2] for k in keys)-1
    ]

    c = 0

    # algo is basically wrapping its way around the object
    # we start at the minimum of all the points
    # so long as we keep hitting air, we keep spreading out and searching
    # - seen stops us from hitting the same point twice
    # once we hit a obsidian point (a value in d) we know we've reached the edge
    # we don't need to keep continuing through that point, cause that would
    # be going into the rock.
    # Instead we keep up the other searches as it goes around that point, eventually
    # hitting all the surface area points from all the sides.

    seen = set()
    queue = [(mins[0], mins[1], mins[2])]
    while queue:
        x, y, z = queue.pop(0)
        if (x, y, z) in seen:
            continue

        # track that we've this point
        seen.add((x, y, z))

        # for all possible adjacent points
        for x_dir, y_dir, z_dir in sides:
            next_x = x+x_dir
            next_y = y+y_dir
            next_z = z+z_dir
            # if the adjacent point has been touched, skip it
            if (next_x, next_y, next_z) in seen:
                continue
            
            # check that the point is still a valid point within our search space
            if mins[0] <= next_x <= maxes[0] and mins[1] <= next_y <= maxes[1] and mins[2] <= next_z <= maxes[2]:
                # this is still valid point within our object
                if (next_x, next_y, next_z) in d:
                    # and its a physical point that we've hit
                    # so that means we came from some outside dirction that we've never touched before
                    # and hit the object, so this must be a surface area side
                    c += 1
                else:
                    # otherwise its not a physical point, so start searching from this new point
                    queue.append((next_x, next_y, next_z))
    
    return c


filename = "input.txt"
print(count_surface_area(parse_input(filename)))
print(count_exterior_surface_area(parse_input(filename)))