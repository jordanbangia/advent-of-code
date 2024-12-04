from dataclasses import dataclass
import uuid


@dataclass
class Coord:
    x: int = -1
    y: int = -1
    z: int = -1

    @staticmethod
    def from_tuple(t: tuple[int, int, int]) -> "Coord":
        return Coord(t[0], t[1], t[2])


@dataclass
class Brick:
    name: str
    low: Coord
    high: Coord

    moved: bool = False

    moved_low: Coord = None
    moved_high: Coord = None

    def __post_init__(self):
        # orient the brick s.t the lowest z position of the brick is at low
        if self.high.z < self.low.z:
            self.low, self.high = self.high, self.low

    # def overlaps_brick


def parse_input(i):
    bricks = []
    with open(f"{i}.txt", "r") as f:
        for line in f.readlines():
            bricks.append(
                Brick(
                    str(uuid.uuid4()),
                    Coord.from_tuple(
                        [int(c) for c in line.strip().split("~")[0].split(",")]
                    ),
                    Coord.from_tuple(
                        [int(c) for c in line.strip().split("~")[1].split(",")]
                    ),
                )
            )
    return bricks


def part_1(i):
    bricks = parse_input(i)

    bricks = sorted(bricks, key=lambda b: b.low.z)

    b_map = {b.name: b for b in bricks}

    print(bricks)

    grid = {}

    def fill(b: Brick):
        for x in range(b.low.x, b.high.x + 1):
            for y in range(b.low.y, b.high.y + 1):
                for z in range(b.low.z, b.high.z + 1):
                    grid[(x, y, z)] = b.name

    for brick in bricks:
        if brick.low.z == 1:
            fill(brick)
        else:
            


print(part_1("test"))
