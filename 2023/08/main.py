
from dataclasses import dataclass
import math


@dataclass
class Node:
    name: str

    left: str
    right: str

    def go(self, dir: str) -> str:
        if dir.upper() == 'L':
            return self.left
        else:
            return self.right

def parse(i) -> tuple[str, dict[str, Node]]:
    nodes = {}

    with open(f"{i}.txt", "r") as f:
        lines = f.readlines()

        instructions = lines[0].strip()

        for l in lines[2:]:
            parts = l.strip().split(" = ")
            node_name = parts[0].strip()
            left_right_parts = parts[1].strip().replace("(", "").replace(")", "").replace(",", "").split(" ")
            left_name = left_right_parts[0].strip()
            right_name = left_right_parts[1].strip()

            nodes[node_name] = Node(node_name, left_name, right_name)
    
        return instructions, nodes

def part_1(i):
    instructions, nodes = parse(i)

    i = 0
    current_node = "AAA"

    while current_node != "ZZZ":
        og = current_node
        instruction = instructions[i % len(instructions)]
        current_node = nodes[current_node].go(instruction)
        i += 1
        # break
        
    
    return i

def part_2(i):
    instructions, nodes = parse(i)

    a_nodes = {n for n in nodes.keys() if n.endswith('A')}
    z_nodes = {n for n in nodes.keys() if n.endswith('Z')}

    print(len(instructions))

    z_to_a = {}
    a_steps = {}


    for n in a_nodes:
        a_node = n
        print(f"Searching {n}")
        steps = 0
        hits = []
        while steps < 100_000:
            instruction = instructions[steps % len(instructions)]
            n = nodes[n].go(instruction)
            steps += 1
            if n in z_nodes:
                if n not in z_to_a:
                    z_to_a[n] = a_node
                    a_steps[a_node] = steps
                else:
                    assert z_to_a[n] == a_node
        
    return math.lcm(*[v for v in a_steps.values()])
        
    

print("part 1:", part_1("input"))
print("part 2:", part_2("input"))


'''
how about we track each node -> and how many steps it takes to get from a given node to 

'''