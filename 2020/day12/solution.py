def get_input():
    data = []
    with open("input.txt", "r") as f:
        for line in f:
            line = line.strip()
            data += [(line[0], int(line[1:]))]
    return data


order = ["N", "E", "S", "W"]


def solve_first():
    print("solution 1:")
    data = get_input()

    class heading:
        direction = "E"

        # assume north, east are positive directions, south, west negative
        east_west = 0
        north_south = 0

    def update_position(cardinal, amount):
        if cardinal == "N":
            heading.north_south += val
        elif cardinal == "S":
            heading.north_south -= val
        elif cardinal == "E":
            heading.east_west += val
        elif cardinal == "W":
            heading.east_west -= val

    def apply_direction(angle):
        steps = int(angle / 90)
        start = -1
        for i, e in enumerate(order):
            if e == heading.direction:
                start = i
                break
        heading.direction = order[(start + steps) % 4]

    for action, val in data:
        if action in "NSEW":
            update_position(action, val)
        elif action == "F":
            update_position(heading.direction, val)
        elif action == "R":
            apply_direction(val)
        elif action == "L":
            apply_direction(-1 * val)

    print(heading.direction, heading.east_west, heading.north_south)
    print(abs(heading.east_west) + abs(heading.north_south))


def solve_second():
    print("solution 2:")

    data = get_input()

    class waypoint:
        east_west = 10
        north_south = 1

        def update_position(self, cardinal, amount):
            if cardinal == "N":
                self.north_south += val
            elif cardinal == "S":
                self.north_south -= val
            elif cardinal == "E":
                self.east_west += val
            elif cardinal == "W":
                self.east_west -= val

        def rotate(self, direction, rotation):
            mult = 1 if direction == "R" else -1

            steps = int(rotation / 90)
            for _ in range(steps):
                new_east_west = self.north_south * mult
                new_north_south = self.east_west * mult * -1
                self.east_west = new_east_west
                self.north_south = new_north_south

        def __repr__(self):
            return f"E/W:{self.east_west}, N/S:{self.north_south}"

    class ship:
        east_west = 0
        north_south = 0

        def move(self, w: waypoint, times):
            self.east_west += w.east_west * times
            self.north_south += w.north_south * times

        def __repr__(self):
            return f"E/W:{self.east_west}, N/S:{self.north_south}"

    s = ship()
    w = waypoint()
    for action, val in data:
        if action in "NSEW":
            w.update_position(action, val)
        elif action in "RL":
            w.rotate(action, val)
        elif action == "F":
            s.move(w, val)

    print(w)
    print(s)
    print(abs(s.east_west) + abs(s.north_south))


solve_first()
print("\n")
solve_second()