

with open('input.txt', 'r') as f:
    lines = [s.strip() for s in f.readlines()]

ns = {
    'one': '1',
    'two': '2',
    'three': '3',
    'four': '4',
    'five': '5',
    'six': '6',
    'seven': '7',
    'eight': '8',
    'nine': '9',
}


o = []

for l in lines:
    first_digit_o = -1
    last_digit_o = -1
    
    w = ''
    for c in l:
        if c.isdigit():
            if first_digit_o == -1:
                first_digit_o = c

            last_digit_o = c

        else:
            # is an alpha character
            w += c

            m = next((n for n in ns.keys() if n in w), None)
            if m:
                if first_digit_o == -1:
                    first_digit_o = ns[m]
                last_digit_o = ns[m]
                w = ''

    first_digit = -1

    w = ''
    for c in l:
        if c.isdigit():
            first_digit = c
            break
        else:
            # is an alpha character
            w += c

            m = next((n for n in ns.keys() if n in w), None)
            if m:
                first_digit = ns[m]
                break
    
    last_digit = -1

    w = ''
    for c in l[::-1]:
        if c.isdigit():
            last_digit = c
            break
        else:
            w = c + w
            m = next((n for n in ns.keys() if n in w), None)
            if m:
                last_digit = ns[m]
                break


    if int(first_digit + last_digit) != int(first_digit_o + last_digit_o):
        print(l, int(first_digit + last_digit), int(first_digit_o + last_digit_o))

    o.append(int(first_digit + last_digit))

print(sum(o)) 



    # w = ''
    # for c in l:
    #     if c.isdigit():
    #         if first_digit == -1:
    #             first_digit = c

    #         last_digit = c

    #     else:
    #         # is an alpha character
    #         w += c

    #         m = next((n for n in ns.keys() if n in w), None)
    #         if m:
    #             if first_digit == -1:
    #                 first_digit = ns[m]
    #             last_digit = ns[m]
    #             w = ''