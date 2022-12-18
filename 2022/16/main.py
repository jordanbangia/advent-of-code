
import re


class Valve():
    def __init__(self, name, flow_rate, n):
        self.name = name
        self.flow_rate = flow_rate
        self.n = n

def parse_input(filename: str):
    prog = re.compile(r'Valve ([A-Z]{2}) has flow rate=(\d+); '
                    r'tunnels? leads? to valves? ((([A-Z]{2}), )*([A-Z]{2}))')

    d = {}

    with open(filename, "r") as f:
        for line in f.readlines():
            match = prog.match(line)
            name = match.group(1)
            flow_rate = int(match.group(2))
            n = set(match.group(3).split(', '))
            valve = Valve(name, flow_rate, n)
            d[name] = valve

    steps = {x:{y:1 if y in d[x].n else float('inf') for y in d} for x in d}

    # Floyd-Warshall algo
    for k in steps:
        for i in steps:
            for j in steps:
                steps[i][j] = min(steps[i][j], steps[i][k] + steps[k][j])
    
    return d, steps


def travel(valves, distances, state_machine, last_valve, time_remaining, state, flow, answer):
    '''
    valves: constant, tracking all the nodes that we could possibly want to visit
    distance: constant, tracking the distances between all nodes
    state_machine: constant, tracking the bit masking for each node

    last_valve: the last valve that we openned
    time_remaining: how much time we have to keep opening valves
    state: the current setup of open / closed valves
    flow: the current amount of flow released
    answer: tracking, mapping states -> how much flow they produce

    This is like the travelling salesman problem.
    For the given start node, we try all possible other nodes that have some value.
    We check for each of these steps if we've already opened this valve (in which case no point going back)
    and we check if we have enough time.
    We also keep track of the best we've seen for a given state (opening valve1 valve2 doesn't produce 
    the same amount as valve2 valve1, but we only care to know that one).
    This could result in going through all permutations but we're likely to run out of time before we
    reach the end of the permutations.

    '''
    answer[state] = max(answer.get(state, 0), flow)
    for valve in valves:
        # travel + open
        minutes = time_remaining - distances[last_valve][valve] - 1
        if (state_machine[valve] & state) or (minutes <= 0):
            # if we've already opened this valve or we've run out of time
            continue
        
        # otherweise, open this valve
        nextState = state | state_machine[valve]
        # track the flow
        newFlow = flow + minutes * valves[valve].flow_rate

        # assuming we just did this node, try the next set of nodes
        travel(valves, distances, state_machine, valve, minutes, nextState, newFlow, answer)
    return answer

def part1(filename:str) -> int:
    all_valves, distances = parse_input(filename)

    minutes = 30

    valves = {name: valve for (name, valve) in all_valves.items() if valve.flow_rate > 0}
    state = {v: 1<<i for i, v in enumerate(valves)}

    last_valve = 'AA'
    starting_state = 0
    starting_flow = 0

    paths =  travel(valves, distances, state, last_valve, minutes, starting_state, starting_flow, {})
    
    max_flow = max(paths.values())
    return max_flow

def part2(filename: str) -> int:
    all_valves, distances = parse_input(filename)

    minutes = 26

    valves = {name: valve for (name, valve) in all_valves.items() if valve.flow_rate > 0}
    state = {v: 1<<i for i, v in enumerate(valves)}

    last_valve = 'AA'
    starting_state = 0
    starting_flow = 0

    # calculate the same set of paths
    paths =  travel(valves, distances, state, last_valve, minutes, starting_state, starting_flow, {})

    # and the best we can do is the paths where neither of us overlap
    # any case where we overlap would technically be double counted
    total_flow = max(
        my_val + elephants_val for k1, my_val in paths.items() for k2, elephants_val in paths.items() if not (k1 & k2)
    )
    return total_flow


if __name__ == "__main__":
    print(part1("input.txt"))
    print(part2("input.txt"))