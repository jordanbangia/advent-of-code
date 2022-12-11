
with open("input.txt", "r") as f:
    commands = [r.strip() for r in f.readlines()]

def is_interesting_cycle(cycle)-> bool:
    return cycle == 20 or (cycle-20) % 40 == 0

x = 1
cycle = 1
strength = 0

sprite_line = [c for c in "###....................................."]

current_row = ""

print_output = []

def update_print_output():
    global current_row
    sprite_indx = (cycle-1) % 40
    if sprite_line[sprite_indx] == "#":
        current_row += "#"
    else:
        current_row += "."
    if cycle % 40 == 0:
        print_output.append(current_row)
        current_row = ""

def update_sprite_line(clear: bool):
    print(x)
    if x-1 < len(sprite_line)-1:
        sprite_line[x-1] = "." if clear else "#"
    if x < len(sprite_line)-1:
        sprite_line[x] = "." if clear else "#"
    if x+1 < len(sprite_line)-1:
        sprite_line[x+1] = "." if clear else "#"


for command_line in commands:
    parts = command_line.split(" ")
    print(command_line, cycle, x, "drawing at", )
    
    if parts[0] == "noop":
        update_print_output()
        cycle += 1
        if is_interesting_cycle(cycle):
            added_str = cycle * x
            strength += added_str
        
    elif parts[0] == "addx":
        update_print_output()
        
        # tick first cycle
        cycle += 1
        if is_interesting_cycle(cycle):
            added_str = cycle * x
            strength += added_str

        # draw next character
        update_print_output()

        # tick second cycle
        cycle += 1

        # clear out the old sprite location
        update_sprite_line(True)

        # increment sprite midpoint
        x += int(parts[1])

        update_sprite_line(False)

        if is_interesting_cycle(cycle+2):
            added_str = cycle * x
            strength += added_str
        

print(strength)

for r in print_output:
    print(r)