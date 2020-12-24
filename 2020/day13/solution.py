bus_ids_str = "13,x,x,41,x,x,x,37,x,x,x,x,x,419,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,19,x,x,x,23,x,x,x,x,x,29,x,421,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,17"


def find_next_leaving(start_time, bus_interval) -> int:
    mod = start_time % bus_interval
    if mod == 0:
        return 0
    return bus_interval - mod


def solve_one():
    earliest_departure_time = 1002394

    bus_ids = []
    for bus_id in bus_ids_str.split(","):
        if bus_id != "x":
            bus_ids += [int(bus_id)]

    next_arrivals = [
        (bus_id, find_next_leaving(earliest_departure_time, bus_id))
        for bus_id in bus_ids
    ]

    next_arrivals.sort(key=lambda entry: entry[1])

    best_entry = next_arrivals[0]

    print(
        "next bus arrives in: ",
        best_entry[1],
        "with bus id: ",
        best_entry[0],
    )
    print(
        "output: ",
        best_entry[0] * best_entry[1],
    )


def solve_two():
    bus_ids = [int(bus_id) if bus_id != "x" else 0 for bus_id in bus_ids_str.split(",")]

    bus_iter = filter(lambda b: b[1], enumerate(bus_ids))
    # print(bus_iter)

    t = 0
    prod = 1
    for depart, bus_id in bus_iter:
        while (t + depart) % bus_id != 0:
            t += prod
        # consider the first loop
        # at this point, we've found a start time t that satisfies
        # bus[0] departs at t and bus[1] departs at t+1

        # we then take that start and multiply it by the bus ID to move it far into the future
        # multiplying by the bus id would still satisfy the relation that we've already found
        prod *= bus_id
    print(t)


# solve_one()
solve_two()