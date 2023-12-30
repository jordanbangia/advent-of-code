from collections import defaultdict, namedtuple
from dataclasses import dataclass
from typing import Literal, Union

H = "high"
L = "low"

Pulse = Literal["high", "low"]


@dataclass
class NodePulse:
    node: str
    pulse: Pulse
    sender: str

    def __str__(self) -> str:
        return f"{self.sender} -{self.pulse}-> {self.node}"


@dataclass
class Button:
    name: str
    dests: list[str]

    def push(self) -> list[NodePulse]:
        return [NodePulse(d, L, self.name) for d in self.dests]


@dataclass
class Broadcaster:
    name: str
    dests: list[str]

    def pulse(self, p: Pulse, node: str) -> list[NodePulse]:
        return [NodePulse(d, p, self.name) for d in self.dests]


@dataclass
class FlipFlop:
    name: str
    dests: list[str]

    state: Literal["on", "off"] = "off"

    def pulse(self, p: Pulse, node: str) -> list[NodePulse]:
        if p == H:
            return []

        self.state = "on" if self.state == "off" else "off"
        return [
            NodePulse(d, L if self.state == "off" else H, self.name) for d in self.dests
        ]


@dataclass
class Conjunction:
    name: str
    dests: list[str]

    memory: dict[str, Pulse] = None

    def set_input_nodes(self, ins: list[str]):
        self.memory = {i: L for i in ins}

    def pulse(self, p: Pulse, node: str) -> list[NodePulse]:
        self.memory[node] = p
        if all(v == H for v in self.memory.values()):
            return [NodePulse(d, L, self.name) for d in self.dests]
        return [NodePulse(d, H, self.name) for d in self.dests]


@dataclass
class Output:
    def pulse(self, p: Pulse, node: str) -> list[NodePulse]:
        return []


Node = Union[Broadcaster, Button, FlipFlop, Conjunction, Output]


def parse_input(i) -> dict[str, Node]:
    nodes = {}

    conjunction_nodes = []

    all_node_names = set()

    with open(f"{i}.txt", "r") as f:
        for line in f.readlines():
            line = line.strip()

            parts = line.split(" -> ")
            node_name = parts[0].strip()
            destinations = [d.strip() for d in parts[1].strip().split(",")]
            all_node_names.update(set(destinations))

            if node_name == "broadcaster":
                nodes["broadcaster"] = Broadcaster("broadcaster", destinations)
                all_node_names.add("broadcaster")
            elif "%" in node_name:
                node_name = node_name.replace("%", "").strip()
                all_node_names.add(node_name)
                nodes[node_name] = FlipFlop(node_name, destinations)
            elif "&" in node_name:
                node_name = node_name.replace("&", "").strip()
                all_node_names.add(node_name)
                nodes[node_name] = Conjunction(node_name, destinations)
                conjunction_nodes.append(node_name)
            else:
                raise ValueError(f"what is this {node_name}")

    for conjunction_node in conjunction_nodes:
        nodes[conjunction_node].set_input_nodes(
            [node_name for node_name, n in nodes.items() if conjunction_node in n.dests]
        )

    nodes["button"] = Button("button", ["broadcaster"])
    all_node_names.add("button")

    for node_name in all_node_names:
        if node_name not in nodes:
            print("Node was not defined, defaulting to an output node")
            nodes[node_name] = Output()

    print(nodes)
    return nodes


def part_1(i, its=1000, debug=False):
    nodes = parse_input(i)
    pulse_count = defaultdict(int)

    for i in range(its):
        button = nodes["button"]

        broadcasts = button.push()

        for broadcast in broadcasts:
            pulse_count[broadcast.pulse] += 1
            if debug:
                print(broadcast)

        while len(broadcasts) > 0:
            next_broadcasts = []

            for broadcast in broadcasts:
                node = broadcast.node
                pulse = broadcast.pulse
                sender = broadcast.sender

                for node_pulse in nodes[node].pulse(pulse, sender):
                    next_broadcasts.append(node_pulse)
                    pulse_count[node_pulse.pulse] += 1

                    if debug:
                        print(node_pulse)

            broadcasts = next_broadcasts

    print(f"{pulse_count[L]} low pulses")
    print(f"{pulse_count[H]} high pulses")

    return pulse_count[L] * pulse_count[H]


def part_2(i):
    """
    not really different from part 2, but instead of
    running it for only a set number of iterations
    we run it continously.

    From trying just once, we can tell that we need to find a
    pattern in the button press results to find when the values
    all line up properly.

    We notice from the input that rx receives a low
    when all of the inputs to rg are high.
    The inputs to rg are kd, zf, gs, vg.
    So we are tracking for when all those inputs line as
    being high at the same time, cause thats the only time
    a collective output to rx of low will come through.

    kd, zf, gs, and vg are also all conjunction nodes.
    So we want to find the # of button presses required to get
    each individual node to output high.  Each node will output
    high when it:
    A) gets a button press to reach it, and
    B) not all of its previously seen inputs are high.
    That will happen only ever n_i cylces.

    We find the lcm of the n_i to find when it will all line up for our rx node.

    """
    nodes = parse_input(i)

    finished = False
    button_count = 0
    while not finished:
        node_pulses = nodes["button"].push()
        button_count += 1

        while len(node_pulses) > 0:
            node_pulse = node_pulses.pop(0)

            if node_pulse.node == "rx" and node_pulse.pulse == L:
                finished = True
                break

            if node_pulse.node not in nodes:
                continue

            if node_pulse.node == "broadcaster":
                node_pulses.extend(
                    nodes[node_pulse.node].pulse(node_pulse.pulse, node_pulse.sender)
                )
            elif isinstance(nodes[node_pulse.node], FlipFlop):
                node_pulses.extend(
                    nodes[node_pulse.node].pulse(node_pulse.pulse, node_pulse.sender)
                )
            elif isinstance(nodes[node_pulse.node], Conjunction):
                node_pulses.extend(
                    nodes[node_pulse.node].pulse(node_pulse.pulse, node_pulse.sender)
                )
                if L in nodes[node_pulse.node].memory.values():
                    new_level = H
                else:
                    new_level = L

                if node_pulse.node in ["kd", "zf", "vg", "gs"] and new_level == H:
                    print("Debug", button_count, "/", node_pulse.node, "/", new_level)


# print(part_1("input"))
print(part_2("input"))

# kd -> 3767
# zf -> 3779
# gs -> 3889
# vg -> 4057
