
import math


class Monkey:

    def __init__(self):
        self.items = []
        self.operation = ""
        self.test = -1
        self.true = -1
        self.false = -1
        self.inspects = 0
        self.mod = 1
    
    def perform_op(self, item):
        v = eval(self.operation.replace("old", str(item)))
        return v % self.mod
    
    def test_throw(self, worry_level) -> bool:
        return worry_level % self.test == 0

    def target(self, test_result) -> int:
        if test_result:
            return self.true
        return self.false

    @staticmethod
    def from_strings(lines: list[str]) -> "Monkey":
        assert "Monkey" in lines[0]

        parts = lines[1].strip().split(":")
        items_s = parts[1].strip().split(",")
        items = [int(s.strip()) for s in items_s]

        parts = lines[2].strip().split("=")
        operation = parts[1].strip()

        parts = lines[3].strip().split(" ")
        divis_test = int(parts[-1])

        true_target = int(lines[4].strip().split(" ")[-1])
        false_target = int(lines[5].strip().split(" ")[-1])

        m = Monkey()
        m.items = items
        m.operation = operation
        m.test = divis_test
        m.true = true_target
        m.false = false_target

        return m



def part_1():
    monkeys = []

    with open("input.txt", "r") as f:
        curr = []
        for line in f.readlines():
            if line == "\n":
                monkeys.append(Monkey.from_strings(curr))
                curr = []
            else:
                curr.append(line)
        monkeys.append(Monkey.from_strings(curr))

    def count_items():
        items = 0
        for m in monkeys:
            items += len(m.items)
        print("items: ", items)
    
    for i in range(20):
        # # print("on turn", i)
        # count_items()
        for monkey in monkeys:
            for item in monkey.items:
                worry_level = math.floor(monkey.perform_op(item) // 3)
                test_result = monkey.test_throw(worry_level)
                next_target = monkey.target(test_result)
                monkeys[next_target].items.append(worry_level)
                # print("passed", worry_level, "to", next_target, test_result)
                monkey.inspects += 1
            monkey.items = []
    
    inspections = sorted([m.inspects for m in monkeys])
    print(inspections[-1], "*", inspections[-2], "=", inspections[-1] * inspections[-2])

def part_2():
    # for this one, we have to use modulo math to keep the numbers sane
    monkeys = []

    with open("input.txt", "r") as f:
        curr = []
        for line in f.readlines():
            if line == "\n":
                monkeys.append(Monkey.from_strings(curr))
                curr = []
            else:
                curr.append(line)
        monkeys.append(Monkey.from_strings(curr))

    def count_items():
        items = 0
        for m in monkeys:
            items += len(m.items)
        print("items: ", items)

    # calculate the common mod so that we can keep the numbers sane
    divs = [m.test for m in monkeys]
    mod = 1
    for d in divs:
        mod *= d

    for m in monkeys:
        m.mod = mod
    
    for i in range(10000):
        print("on turn", i)
        # count_items()
        for monkey in monkeys:
            for item in monkey.items:
                worry_level = monkey.perform_op(item)
                test_result = monkey.test_throw(worry_level)
                next_target = monkey.target(test_result)
                monkeys[next_target].items.append(worry_level)
                # print("passed", worry_level, "to", next_target, test_result)
                monkey.inspects += 1
            monkey.items = []
    
    inspections = sorted([m.inspects for m in monkeys])
    print(inspections[-1], "*", inspections[-2], "=", inspections[-1] * inspections[-2])


if __name__ == "__main__":
    print("Part 1:")
    part_1()
    print("\nPart 2:")
    part_2()