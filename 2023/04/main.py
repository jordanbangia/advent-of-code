
total_score = 0
with open('input.txt', 'r') as f:

    lines = f.readlines()

    next_num_q = [1] * (len(lines) + 1)
    next_num_q[0] = 0

    for i, line in enumerate(lines):
        i += 1
        extra_cards = next_num_q[i]

        line = line.strip()
        # print(line.split(':')[0])
        numbers = line.split(':')[1].strip().split('|')
        winners = {int(a) for a in numbers[0].split(' ') if len(a) > 0}
        my_nums = {int(a) for a in numbers[1].split(' ') if len(a) > 0}

        my_winning_numbers = my_nums.intersection(winners)
        # print(my_winning_numbers)
        if len(my_winning_numbers) != 0:
            card_score = 2**(len(my_winning_numbers) - 1)
            # print(card_score)
            total_score += card_score
        
        if len(my_winning_numbers) == 0:
            continue
        else:
            for k in range(len(my_winning_numbers)):
                next_num_q[i+1+k] += extra_cards
                print(f"added {extra_cards} to {i+1+k}")


print(total_score)
print(sum(next_num_q))