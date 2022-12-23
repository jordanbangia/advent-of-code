from collections import namedtuple

def add(x, y):
    if isinstance(x, (int, float)) and isinstance(y, (int, float)):
        return x + y
    
    elif isinstance(x, list):
        return x + [(add, y)]
    else:
        return y + [(add, x)]
    
def mult(x, y):
    if isinstance(x, (int, float)) and isinstance(y, (int, float)):
        return x * y
    elif isinstance(x, list):
        return x + [(mult, y)]
    else:
        return y + [(mult, x)]

def sub(x, y):
    if isinstance(x, (int, float)) and isinstance(y, (int, float)):
        return x - y
    elif isinstance(x, list):
        return x + [(sub, y)]
    else:
        return y + [(mult, -1), (add, x)]

def div(x, y):
    if isinstance(x, (int, float)) and isinstance(y, (int, float)):
        return x / y
    elif isinstance(x, list):
        return x + [(div, y)]
    else:
        raise Exception("can't handle this")

def opposite_op(op):
    if op[0] == add:
        return sub
    elif op[0] == sub:
        return add
    elif op[0] == mult:
        return div
    elif op[0] == div:
        return mult

def parse_input(filename: str):
    d = dict()
    with open(filename, 'r') as f:
        for line in f.readlines():
            line = line.strip()

            key, equation = line.split(':')

            if any(i in equation for i in ['+', '-', '*', '/']):
                # its an equation line
                if '+' in equation:
                    parts = equation.strip().split('+')
                    d[key.strip()] = [parts[0].strip(), add, parts[1].strip()]
                elif '-' in equation:
                    parts = equation.strip().split('-')
                    d[key.strip()] = [parts[0].strip(), sub, parts[1].strip()]
                elif '*' in equation:
                    parts = equation.strip().split('*')
                    d[key.strip()] = [parts[0].strip(), mult, parts[1].strip()]
                elif '/' in equation:
                    parts = equation.strip().split('/')
                    d[key.strip()] = [parts[0].strip(), div, parts[1].strip()]
            else:
                d[key.strip()] = int(equation.strip())
    return d

def part_1(filename: str):
    d = parse_input(filename)

    keys = list(d.keys())

    while not isinstance(d['root'], (int, float)):
        for k in keys:
            if isinstance(d[k], (int, float)):
                continue
            else:
                num1 = d[k][0]
                num2 = d[k][2]
                
                if isinstance(d[num1], (int, float)) and isinstance(d[num2], (int, float)):
                    d[k] = d[k][1](d[num1], d[num2])
    
    return d['root']


class Node:
    def __init__(self, name):
        self.name = name
        self.left = None
        self.right = None
        self.op = None
        self.val = None

    def eval(self):
        if self.val is not None:
            return self.val

        self.val = self.op(self.left.eval(), self.right.eval())
        return self.val

def part_2(filename: str):
    # I want to create a tree based on this equation
    # each node in the tree is either a pass through with one child
    # or an operation with 2 children

    d = parse_input(filename)

    monkeys = dict()
    for monkey_name in d.keys():
        monkeys[monkey_name] = Node(monkey_name)
    
    for monkey_name, eq in d.items():
        if isinstance(eq, (int, float)):
            monkeys[monkey_name].val = eq
        else:
            monkeys[monkey_name].left = monkeys[eq[0]]
            monkeys[monkey_name].right = monkeys[eq[2]]
            monkeys[monkey_name].op = eq[1]

    root = "root"
    me = "humn"

    relies_on_me = [m for m in monkeys.values() if (m.left and m.left.name == me) or (m.right and m.right.name == me)]
    assert len(relies_on_me) == 1

    monkeys[me].val = ['x']

    left_val = monkeys[root].left.eval()
    right_val = monkeys[root].right.eval()

    lhs = left_val if isinstance(left_val, list) else right_val
    rhs = left_val if not isinstance(left_val, list) else right_val

    print(lhs)
    while True:
        a = lhs.pop()
        if a == 'x':
            return rhs
        print(a, opposite_op(a), a[1])
        rhs = opposite_op(a)(rhs, a[1])
        


    print(lhs)
    print(rhs)
    


# print(part_1('input.txt'))
print(part_2('input.txt'))
        

    