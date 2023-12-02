with open('input.txt', 'r') as f:
    lines = [s.strip() for s in f.readlines()]


games = []

red_limit = 12
green_limit = 13
blue_limit = 14

ids = []


g = []

for line in lines:
    parts = line.split(":")

    game_part = parts[0]
    game_num = int(game_part.replace("Game ", ""))

    draws_part = parts[1]
    draws = draws_part.split(";")

    d = []
    for draw in draws:
        pulls = draw.split(",")
        p = {}
        for pull in pulls:
            if "blue" in pull:
                p["blue"] = int(pull.strip().replace(" blue", ""))
            elif "red" in pull:
                p["red"] = int(pull.strip().replace(" red", ""))
            elif "green" in pull:
                p["green"] = int(pull.strip().replace(" green", ""))
            else:
                print("WHOOOOOPS")
        d.append(p)
    g.append((game_num, d))



# part 1
game_id_sum = 0
for game in g:
    if any(d.get("red", 0) > red_limit for d in game[1]):
        continue
    if any(d.get("blue", 0) > blue_limit for d in game[1]):
        continue
    if any(d.get("green", 0) > green_limit for d in game[1]):
        continue
    game_id_sum += game[0]

print(game_id_sum)

# part 2
power_sum = 0
for game in g:
    red_needed = max(d.get("red", 0) for d in game[1])
    blue_needed = max(d.get("blue", 0) for d in game[1])
    green_needed = max(d.get("green", 0) for d in game[1])

    power_sum += red_needed*blue_needed*green_needed

print(power_sum)
    