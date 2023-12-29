from dataclasses import dataclass, asdict
from typing import Optional

ACCEPTED = "A"
REJECTED = "R"
NEXT_RULE = "NEXT_RULE"


@dataclass
class Part:
    x: int
    m: int
    a: int
    s: int

    def tot(self):
        return self.x + self.m + self.a + self.s


Range = tuple[int, int]


@dataclass
class PartRange:
    x: Range
    m: Range
    a: Range
    s: Range

    def distinct_combos(self):
        return (
            (self.x[1] - self.x[0] + 1)
            * (self.m[1] - self.m[0] + 1)
            * (self.a[1] - self.a[0] + 1)
            * (self.s[1] - self.s[0] + 1)
        )


@dataclass
class Rule:
    target: Optional[str]
    condition: Optional[str]
    value: Optional[int]
    next_workflow: str

    def check_part(self, p: Part) -> str:
        if self.target is None:
            return self.next_workflow

        part_val = asdict(p)[self.target]

        if self.condition == ">" and part_val > self.value:
            return self.next_workflow
        elif self.condition == "<" and part_val < self.value:
            return self.next_workflow
        else:
            return None

    def apply_rule_to_range(self, p: PartRange) -> list[tuple[PartRange, str]]:
        if self.target is None:
            return [(PartRange(p.x, p.m, p.a, p.s), self.next_workflow)]

        part_val_range = asdict(p)[self.target]

        r = []
        if self.condition == ">" and part_val_range[1] > self.value:
            # we need the section of the range > self.value
            d = asdict(p)
            d[self.target] = (max(self.value + 1, part_val_range[0]), part_val_range[1])
            r.append((PartRange(**d), self.next_workflow))
            if part_val_range[0] <= self.value:
                d = asdict(p)
                d[self.target] = (part_val_range[0], self.value)
                r.append((PartRange(**d), NEXT_RULE))
        elif self.condition == "<" and part_val_range[0] < self.value:
            d = asdict(p)
            d[self.target] = (part_val_range[0], min(self.value - 1, part_val_range[1]))
            r.append((PartRange(**d), self.next_workflow))
            if part_val_range[1] >= self.value:
                d = asdict(p)
                d[self.target] = (self.value, part_val_range[1])
                r.append((PartRange(**d), NEXT_RULE))

        return r


@dataclass
class Workflow:
    name: str
    rules: list[Rule]


def parse_input(input_file_name: str) -> tuple[dict[str, Workflow], list[Part]]:
    workflows = {}
    parts = []

    process_parts = False
    with open(f"{input_file_name}.txt", "r") as f:
        for line in f.readlines():
            if line == "\n":
                process_parts = True
                continue

            if process_parts:
                line = line.strip().removeprefix("{").removesuffix("}")
                p = line.split(",")
                assert len(p) == 4
                x, m, a, s = 0, 0, 0, 0
                for s in p:
                    if "x" in s:
                        x = int(s.replace("x=", ""))
                    elif "m" in s:
                        m = int(s.replace("m=", ""))
                    elif "a" in s:
                        a = int(s.replace("a=", ""))
                    elif "s" in s:
                        s = int(s.replace("s=", ""))
                parts.append(Part(x, m, a, s))
            else:
                line = line.strip()
                p = line.split("{")
                name = p[0].strip()

                instructions = p[1].replace("}", "").split(",")

                rules = []
                for instruction in instructions:
                    q = instruction.strip().split(":")
                    if len(q) == 1:
                        rules.append(Rule(None, None, None, instruction))
                        continue

                    next_workflow = q[1].strip()
                    instruction = q[0]
                    target = instruction[0]
                    operation = instruction[1]
                    value = int(instruction[2:])
                    rules.append(Rule(target, operation, value, next_workflow))

                workflows[name] = Workflow(name, rules)

    return workflows, parts


def part_1(i):
    workflows, parts = parse_input(i)

    def apply_rules(workflow: Workflow, part: Part) -> str:
        for rule in workflow.rules:
            n = rule.check_part(part)
            if n == ACCEPTED or n == REJECTED:
                return n
            elif n in workflows:
                return n
            else:
                # try the next rule
                continue
        raise ValueError("Should not hit this")

    part_sum = 0

    for part in parts:
        next_workflow = "in"
        while next_workflow not in (ACCEPTED, REJECTED):
            next_workflow = apply_rules(workflows[next_workflow], part)

        if next_workflow == ACCEPTED:
            part_sum += part.tot()

    return part_sum


def part_2(i):
    workflows, _ = parse_input(i)
    print(workflows)
    print("\n\n")

    part_ranges = [
        (PartRange(x=(1, 4000), m=(1, 4000), a=(1, 4000), s=(1, 4000)), "in")
    ]

    accepted_ranges = []

    def apply_rules(workflow, part: PartRange) -> list[tuple[PartRange, str]]:
        nextes = []

        parts = [part]

        for rule in workflow.rules:
            next_parts = [o for p in parts for o in rule.apply_rule_to_range(p)]

            nextes.extend([p for p in next_parts if p[1] != NEXT_RULE])
            parts = [p[0] for p in next_parts if p[1] == NEXT_RULE]

        return nextes

    i = 0
    while len(part_ranges) > 0:
        p = part_ranges.pop(0)
        next_workflow = p[1]
        part = p[0]
        print(part, next_workflow)

        if next_workflow == REJECTED:
            continue
        elif next_workflow == ACCEPTED:
            accepted_ranges.append(part)
            print("Accepted!")
        else:
            split_parts = apply_rules(workflows[next_workflow], part)
            print(split_parts)
            part_ranges.extend(split_parts)

        i += 1
    """
    we can do range work again
    so assume a part with all the available ranges
    we try each workflow starting from in again
    and we try each rule,  we calculate for each rule conditions that pass it and considitons that fail it
    end conditions ==>
        - ACCEPT means add it to our list of things
        - REJECT = [] no ranges
    in the end, count up all the ranges and the product of the xmas values

    167409079868000
    167409079868000
    """

    return sum(p.distinct_combos() for p in accepted_ranges)


# print(part_1("input"))
print(part_2("input"))
