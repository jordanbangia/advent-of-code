
from collections import defaultdict
from itertools import product


class Sensor:
    def __init__(self, x, y, beacon) -> None:
        self.x = x
        self.y = y
        self.nearest_beacon = beacon
        self.dist = abs(beacon.x - x) + abs(beacon.y - y)

    @property
    def key(self):
        return (self.x, self.y)

    def coverage(self, y):
        # return all pairs of points for a given sensor at this y
        used_dist = abs(self.y - y)
        if used_dist > self.dist:
            return None
        return [self.x - (self.dist-used_dist), self.x + (self.dist - used_dist)]

    def in_range(self, coords):
        return dist(coords, (self.x, self.y)) <= self.dist

class Beacon:
    def __init__(self, x, y) -> None:
        self.x = x
        self.y = y
    
    @property
    def key(self):
        return (self.x, self.y)

def dist(c1, c2):
    return abs(c1[0] - c2[0]) + abs(c1[1] - c2[1])

def parse_input(file_name: str):
    sensors = []
    beacons = []

    d = defaultdict(lambda: None)

    min_x = float('inf')
    min_y = float('inf')
    max_y = -1
    max_x = -1

    with open(file_name, "r") as f:
        for line in f.readlines():
            sensor_part, beacon_part = line.strip().split(":")

            sensor_x, sensor_y = sensor_part.strip().replace("Sensor at ", '').split(",")
            sensor_x = int(sensor_x.replace('x=', ''))
            sensor_y = int(sensor_y.replace('y=', ''))

            beacon_x, beacon_y = beacon_part.strip().replace("closest beacon is at ", '').split(',')
            beacon_x = int(beacon_x.replace('x=', ''))
            beacon_y = int(beacon_y.replace('y=', ''))

            beacon = Beacon(beacon_x, beacon_y)
            beacons.append(beacon)
            sensor = Sensor(sensor_x, sensor_y, beacon)
            sensors.append(sensor)

            d[beacon.key] = beacon
            d[sensor.key] = sensor

            max_y = max(max_y, sensor_y, beacon_y)
            max_x = max(max_x, sensor_x, beacon_x)
            min_y = min(min_y, sensor_y, beacon_y)
            min_x = min(min_x, sensor_x, beacon_x)


    return d, sensors


def print_data(data, x_range, y_range):
    for i in range(y_range[0], y_range[1] + 1):
        row = []
        for j in range(x_range[0], x_range[1] + 1):
            v = data[(j, i)]
            if v is None:
                row.append('.')
            elif isinstance(v, Sensor):
                row.append('S')
            elif isinstance(v, Beacon):
                row.append('B')
            else:
                row.append(v)
        print(''.join(row))


def part_1():
    data, sensors = parse_input('sample.txt')
    # print_data(data, x_range, y_range)

    interest_y = 10

    c = 0
    for i, sensor in enumerate(sensors):
        print("sensor", i)
        top = sensor.y - sensor.dist
        bot = sensor.y + sensor.dist
        if not(top <= interest_y <= bot):
            # this sensor doesn't cross over the y of interest
            # no point in doing its processing
            continue

        # use a simplification, instead of going across all x and all y, we only
        # need to consider (x,y) coords that are on the y of interest
        for x in range(sensor.x - sensor.dist, sensor.x + sensor.dist + 1):
            if dist((x,interest_y), sensor.key) <= sensor.dist:
                    if data[(x,interest_y)] is None:
                        data[(x,interest_y)] = '#'
                        c += 1
        
    # print_data(data, x_range, y_range)
    print(c)

