import re
from collections import namedtuple
import time

State = namedtuple('State', [
    't',
    'r',
    'b',
    'p',
    'ps',
])

Resources = namedtuple('Resources', [
    'ore',
    'clay',
    'obs',
    'geo'
])

Bots = namedtuple('Bots', [
    'ore',
    'clay',
    'obs',
    'geo'
])

Buildable = namedtuple('PreviouslyBuildable', [
    'ore',
    'clay',
    'obs',
    'geo',
])

class Blueprint:
    def __init__(self, id, ore_robot_ore_cost, clay_r_ore_cost, obs_r_ore_cost, obs_r_clay_cost, geo_r_ore_cost, geo_r_obs_cost):
        self.id = id

        self.ore_r_ore_cost = ore_robot_ore_cost
        self.clay_r_ore_cost = clay_r_ore_cost
        self.obs_r_ore_cost = obs_r_ore_cost
        self.obs_r_clay_cost = obs_r_clay_cost
        self.geo_r_ore_cost = geo_r_ore_cost
        self.geo_r_obs_cost = geo_r_obs_cost

    def __str__(self):
        return ','.join([str(i) for i in [
            self.id,
            self.ore_r_ore_cost,
            self.clay_r_ore_cost,
            self.obs_r_ore_cost,
            self.obs_r_clay_cost,
            self.geo_r_ore_cost,
            self.geo_r_obs_cost
        ]])

    def gen_resources(self, s: State) -> Resources:
        return Resources(
            s.b.ore,
            s.b.clay,
            s.b.obs,
            s.b.geo
        )

    def gen_buildable(self, s: State) -> Buildable:
        return Buildable(
            ore=s.r.ore >= self.ore_r_ore_cost,
            clay=s.r.ore >= self.clay_r_ore_cost,
            obs=s.r.ore >= self.obs_r_ore_cost and s.r.clay >= self.obs_r_clay_cost,
            geo=s.r.ore >= self.geo_r_ore_cost and s.r.obs >= self.geo_r_obs_cost
        )

    def buy_ore_bot(self, state: State):
        return (
            Resources(
                state.r.ore - self.ore_r_ore_cost,
                state.r.clay,
                state.r.obs,
                state.r.geo,
            ),
            Bots(
                state.b.ore + 1,
                state.b.clay,
                state.b.obs,
                state.b.geo,
            )
        )
    
    def buy_clay_bot(self, state: State):
        return (
            Resources(
                state.r.ore - self.clay_r_ore_cost,
                state.r.clay,
                state.r.obs,
                state.r.geo,
            ),
            Bots(
                state.b.ore,
                state.b.clay + 1,
                state.b.obs,
                state.b.geo,
            )
        )

    def buy_obs_bot(self, state: State):
        return (
            Resources(
                state.r.ore - self.obs_r_ore_cost,
                state.r.clay - self.obs_r_clay_cost,
                state.r.obs,
                state.r.geo,
            ),
            Bots(
                state.b.ore,
                state.b.clay,
                state.b.obs + 1,
                state.b.geo,
            )
        )
                    
    def buy_geo_bot(self, state: State):
        return (
            Resources(
                state.r.ore - self.geo_r_ore_cost,
                state.r.clay,
                state.r.obs - self.geo_r_obs_cost,
                state.r.geo,
            ),
            Bots(
                state.b.ore,
                state.b.clay,
                state.b.obs,
                state.b.geo + 1,
            )
        )
    
    def combine_resources(self, r1, r2):
        return Resources(
            ore=r1.ore + r2.ore,
            clay=r1.clay + r2.clay,
            obs=r1.obs + r2.obs,
            geo=r1.geo + r2.geo,
        )

    def sim(self, time_limit) -> int:
        max_geodes = -1

        initial_state = State(0, Resources(0, 0, 0, 0), Bots(1, 0, 0, 0), Buildable(False, False, False, False), None)

        queue = [initial_state]
        while len(queue) > 0:
            next_queue_items = []
            while len(queue) > 0:
                state = queue.pop(0)
                
                max_geodes = max(state.r.geo, max_geodes)
                if state.r.geo < max_geodes:
                    continue
                if state.t == time_limit:
                    continue

                # we are going to calculate all possible next states
                # given our current state and one tick of the clock
                # note that we can only build one robot at a time

                # the order we check is important, we want to prioritize getting
                # to geo bots as that maximizes our chance of finding a good path

                buildable = self.gen_buildable(state)

                next_states = []
                if buildable.geo:
                    # its always considered worth while buying a geode producing robot
                    # as it gets us closer to the goal
                    # so if can do it then build the bot, not worth considering the other options
                    next_states.append(self.buy_geo_bot(state))
                else:
                    if buildable.obs and not state.p.obs:
                        # could buy an obsidian producing robot
                        
                        # its only worthwhile to make up to geo_r_obs_cost obs_bots
                        # any more means that we're producing more obs than we can possibly use
                        # any less and we have to wait multiple iterations to get enough to build a geo bot
                        if state.b.obs < self.geo_r_obs_cost:
                            next_states.append(self.buy_obs_bot(state))
                        
                    if buildable.clay and not state.p.clay:
                        # could buy a clay producing robot

                        # same argument, its only worthwhile to buy up to obs_r_clay_cost clay bots
                        # any more and we are over producing
                        if state.b.clay < self.obs_r_clay_cost:
                            next_states.append(self.buy_clay_bot(state))
                    if buildable.ore and not state.p.ore:
                        # could buy another ore producing robot

                        # here it only makes sense to buy up to the max ore cost of any bot
                        # any more is unecessary, as we can't use it
                        # any less and for some bot we won't be able to produce it at every tick
                        # realistically, this could be stopped at self.geo_r_ore_cost
                        if state.b.ore < max(self.geo_r_ore_cost, self.clay_r_ore_cost, self.ore_r_ore_cost):
                            next_states.append(self.buy_ore_bot(state))
                    
                    # include the option of doing nothing and just collecting
                    # but in doing so, also track what we could have purchased
                    # there's no point buying something in the next iteration
                    # that could have been bought in this iteration.
                    next_queue_items.append(
                        State(
                            state.t + 1, 
                            self.combine_resources(state.r, self.gen_resources(state)),
                            Bots(state.b.ore, state.b.clay, state.b.obs, state.b.geo),
                            buildable,
                            state,
                        ),
                    )

                # compute how many resources we'll get at the end of the minute
                next_resources = self.gen_resources(state)
                for nr, nb in next_states[::-1]:
                    # we're gonna BFS it
                    s = State(
                        state.t + 1,
                        self.combine_resources(nr, next_resources),
                        nb, 
                        Buildable(False, False, False, False),
                        state,
                    )
                    next_queue_items.append(s)
            queue.extend(next_queue_items)
        
        return max_geodes


def parse_input(filename: str):
    blueprints = []
    with open(filename, 'r') as f:
        for line in f.readlines():
            line = line.strip()
            matches = re.search(
                "Blueprint ([0-9]*): Each ore robot costs ([0-9]*) ore. Each clay robot costs ([0-9]*) ore. Each obsidian robot costs ([0-9]*) ore and ([0-9]*) clay. Each geode robot costs ([0-9]*) ore and ([0-9]*) obsidian.",
                line
            )

            blueprints.append(
                Blueprint(
                    int(matches.group(1)),
                    int(matches.group(2)),
                    int(matches.group(3)),
                    int(matches.group(4)),
                    int(matches.group(5)),
                    int(matches.group(6)),
                    int(matches.group(7)),
                )
            )
    return blueprints

def part_1(filename):
    bps = parse_input(filename)
    quality_level = 0
    for blueprint in bps:
        start = time.time()
        max_geos = blueprint.sim(24)
        print("Blueprint", blueprint.id, "produced", max_geos, ". That took", time.time() - start)
        quality_level += blueprint.id * max_geos
    print(quality_level)

def part_2(filename):
    bps = parse_input(filename)[:3]

    max_geo_prod = 1
    for blueprint in bps:
        start = time.time()
        max_geos = blueprint.sim(32)
        print("Blueprint", blueprint.id, "produced", max_geos, ". That took", time.time() - start)
        max_geo_prod *= max_geos
    print(max_geo_prod)

filename = "input.txt"
part_1(filename)
part_2(filename)