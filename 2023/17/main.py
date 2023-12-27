# classic dynamic programming / search
# start at the given position
# try moving in any possible direction
# - move forward, tracking that we've now moved x units forward
# - move left, resetting the moved forward units
# - move right, resetting the move forward units
# keep track of heat loss over time

# how do we prune branches that aren't useful?
# - if we ever hit the end, remember that thats the lowest value we've seen
#   then we can start removing branches that are already at more heat loss
#   then we have
# - if we get to a certain position again with higher heat loss
#   than when we were there previously, then we could remove that branch
#   as well?  (maybe the combo of position + straight movement)
#   - we want to allow for a possible path that loops through and uses
#     a square again to produce an overall lower score
#     but we should avoid redundantly checking the same tile over and over again


# what is the state
# current position, direction of travel, # of steps moved forward, current total heat loss

# NOPE!

# It's Dijkstra's algorithm, trying to find the minimum cost path through the grid


from heapq import heappop, heappush
from itertools import count
from collections import defaultdict


def parse_input(i):
    with open(f"{i}.txt", "r") as f:
        return [[int(n) for n in line.strip()] for line in f.readlines()]


movement = {
    0: (0, -1),
    1: (1, 0),
    2: (0, 1),
    3: (-1, 0),
}


def find_min_heat_loss(
    i: str,  # input file name,
    max_consecutive_steps: int,
    min_consecutive_steps: int,
):
    heat_map = parse_input(i)

    pq = []
    entry_finder = {}
    # auto incrementing counter
    counter = count()

    def add_task(task, priority=0):
        if task in entry_finder:
            remove_task(task)
        count = next(counter)
        entry = [priority, count, task]
        entry_finder[task] = entry
        heappush(pq, entry)

    def remove_task(task):
        entry = entry_finder.pop(task)
        entry[-1] = "REMOVED"

    def pop_task():
        while pq:
            priority, count, task = heappop(pq)
            if task != "REMOVED":
                del entry_finder[task]
                return task

    # create an initial baseline of every state with an extremely high heat
    # we look to improve on this at each iteration
    for y in range(len(heat_map)):
        for x in range(len(heat_map[0])):
            for direction in movement.keys():
                for consecutive in range(1, max_consecutive_steps + 1):
                    add_task((x, y, direction, consecutive), 1000000)

    add_task((0, 0, 1, 0))
    add_task((0, 0, 2, 0))

    total_heat = defaultdict(lambda: 1000000)
    total_heat[(0, 0, 1, 0)] = 0
    total_heat[(0, 0, 2, 0)] = 0

    while True:
        # get a new task
        t = pop_task()
        x, y, direction, consecutive = t

        # if we've reached the end condition we can return
        # we know it shold be the min heat used
        if x == len(heat_map[0]) - 1 and y == len(heat_map) - 1:
            return total_heat[t]

        # find all neighbouring moves
        # these are a turn, and reset the consecutive direction
        neighbors = []
        if consecutive >= min_consecutive_steps:
            neighbors.extend([((direction + 1) % 4, 1), ((direction - 1) % 4, 1)])
        if consecutive < max_consecutive_steps:
            # include going straight if we haven't gone too far yet
            neighbors.append((direction, consecutive + 1))

        # for each neighbour, we calcualte its new position
        for neighbor in neighbors:
            new_direction, new_consecutive = neighbor
            new_x, new_y = (
                x + movement[new_direction][0],
                y + movement[new_direction][1],
            )

            # check if new_x / new_y are in bounds
            if 0 <= new_x < len(heat_map[0]) and 0 <= new_y < len(heat_map):
                new_t = (new_x, new_y, new_direction, new_consecutive)
                new_heat = total_heat[t] + heat_map[new_y][new_x]
                # check if this new heat is less than the total heat seen at this
                # state previously.  State consisting of locaiton + direction + steps
                if new_heat < total_heat[new_t]:
                    total_heat[new_t] = new_heat
                    # if its less, we set that as the new best one we've seen
                    # and enqueue the task
                    add_task(new_t, new_heat)


def part_1(i):
    return find_min_heat_loss(i, 3, 0)


def part_2(i):
    return find_min_heat_loss(i, 10, 4)


case = "input"
print(part_1(case))
print(part_2(case))
