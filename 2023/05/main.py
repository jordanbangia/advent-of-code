


from dataclasses import dataclass

@dataclass
class Map:
    source: int
    dest: int
    range: int

    def s_to_d(self, source_num):
        if self.source <= source_num < self.source + self.range:
            return self.dest + (source_num - self.source)
        return -1

@dataclass
class MapList:
    maps: list[Map]

    cache: dict[int, int]

    def s_to_d(self, source_num):
        if source_num in self.cache:
            print("cache hit!")
            return self.cache[source_num]
        
        for m in self.maps:
            a = m.s_to_d(source_num)
            if a != -1:
                self.cache[source_num] = a
                return a
        self.cache[source_num] = source_num
        return source_num



seeds = []

seed_to_soil = MapList([], {})
soil_to_fertilizer = MapList([], {})
fertilizer_to_water = MapList([], {})
water_to_light = MapList([], {})
light_to_temperature = MapList([], {})
temperature_to_humidity = MapList([], {})
humidity_to_location = MapList([], {})

def parse_map_line(l):
    parts = l.strip().split()
    destination_start = int(parts[0])
    source_start = int(parts[1])
    range_len = int(parts[2])
    return Map(source_start, destination_start, range_len)
    

with open('input.txt', 'r') as f:
    mode = ''

    for line in f.readlines():
        if "seeds:" in line:
            seeds = [int(s) for s in line.strip().replace("seeds:", "").strip().split(' ')]

        elif "seed-to-soil" in line:
            mode = 'seed-to-soil'
        elif "soil-to-fertilizer" in line:
            mode = 'soil-to-fertilizer'
        elif "fertilizer-to-water" in line:
            mode = 'fertilizer-to-water'
        elif "water-to-light" in line:
            mode = 'water-to-light'
        elif "light-to-temperature" in line:
            mode = "light-to-temperature"
        elif "temperature-to-humidity" in line:
            mode = "temperature-to-humidity"
        elif "humidity-to-location" in line:
            mode = "humidity-to-location"
        elif line.strip() == '':
            mode = ''
        
        elif mode == 'seed-to-soil':
            seed_to_soil.maps.append(parse_map_line(line))
        elif mode == 'soil-to-fertilizer':
            soil_to_fertilizer.maps.append(parse_map_line(line))
        elif mode == 'fertilizer-to-water':
            fertilizer_to_water.maps.append(parse_map_line(line))
        elif mode == 'water-to-light':
            water_to_light.maps.append(parse_map_line(line))
        elif mode == 'light-to-temperature':
            light_to_temperature.maps.append(parse_map_line(line))
        elif mode == 'temperature-to-humidity':
            temperature_to_humidity.maps.append(parse_map_line(line))
        elif mode == 'humidity-to-location':
            humidity_to_location.maps.append(parse_map_line(line))


def _convert_seed_to_location(seed):
    return humidity_to_location.s_to_d(
            temperature_to_humidity.s_to_d(
                light_to_temperature.s_to_d(
                    water_to_light.s_to_d(
                        fertilizer_to_water.s_to_d(
                            soil_to_fertilizer.s_to_d(
                                seed_to_soil.s_to_d(seed)
                            )
                        )
                    )
                )
            )
        )



print("part 1")
print(min(_convert_seed_to_location(s) for s in seeds))

'''
For part 2 what we want to do is operate on a whole range of values at a time.

Given seed_s, seed_e, we want to convert all the values to soil
To do this, we check the each soil range we have.  For any intersection
we can do the immedaite mapping, ie.
(seed_a, seed_b) -> (soil_a, soil_b)
For any other area, we would need to extract the remaining region.

Maybe something like:
as a region, what can you output, and whats left over?
it returns the part it converted and everything that it couldn't
try the remaining "couldn't" parts in the remaining mappers
anyything left over is considered a direct map


'''


def _c(seed_range: list[tuple[int, int]], mappers: MapList):
    rs = seed_range
    o = []
    for m in mappers.maps:
        n = []
        for r in rs:
            start, end = r[0], r[1]
            ms, me = m.source, m.source + m.range - 1
            
            # |1 -- 1| |2 -- 2| or |2 -- 2| |1 -- 1|
            if end < ms or me < start:
                n.append(r)
                print('0', n)

            # |1 -- |2 -- 2| -- 1|
            elif start <= ms and me <= end:
                o.append((m.s_to_d(ms), m.s_to_d(me)))
                print('1', o, n)

                if start != ms:
                    n.append((start, ms-1))
                if end != me:
                    n.append((me+1, end))

            # |1 -- |2 -- 1| -- 2|
            elif start <= ms and end <= me:
                o.append(
                    (m.s_to_d(ms), m.s_to_d(end))
                )
                if start != ms:
                    n.append((start, ms-1))
                print('2', o, n)
            

            # |2 -- |1 -- 1| -- 2|
            elif ms <= start and end <= me:
                o.append((m.s_to_d(start), m.s_to_d(end)))
                print('3', o)
            
             # case 5 |2 -- |1 -- 2| -- 1|:
            elif ms <= start and me <= end:
                o.append(
                    (m.s_to_d(start), m.s_to_d(me))
                )
                n.append((me+1, end))
                print('4', o, n)
        rs = n
    
    # anything left over just maps out compeltely
    o.extend(rs)
    if any(i[0] > i[1] for i in o):
        raise ValueError(o)
    return o


print("part 2")

min_loc = float('inf')

for i in range(0, len(seeds), 2):
    seed = seeds[i]
    offset = seeds[i+1]

    soils = _c([(seed, seed+offset-1)], seed_to_soil)
    print('s', soils)
    
    ferts = _c(soils, soil_to_fertilizer)
    print('f', ferts)
    waters = _c(ferts, fertilizer_to_water)
    print('w', waters)
    lights = _c(waters, water_to_light)
    print('l', lights)
    temps = _c(lights, light_to_temperature)
    print('t', temps)
    humids = _c(temps, temperature_to_humidity)
    print('h', humids)
    locs = _c(humids, humidity_to_location)
    print('l', locs)

    min_loc = min(min(t[0] for t in locs), min_loc)


print(min_loc)