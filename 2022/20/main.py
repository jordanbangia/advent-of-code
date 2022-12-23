

def parse_input(filename: str) -> list[int]:
    with open(filename, 'r') as f:
        return [
            int(a.strip()) for a in f.readlines()
        ]
    

def mix(filename: str, multiplier: int = 1, mixes: int = 1) -> int:
    nums = parse_input(filename)
    n = len(nums)

    orig = [(i, a*multiplier) for i, a in enumerate(nums)]
    nums = [a for a in orig]
    assert len(set(nums)) == n
    

    for _ in range(mixes):
        # for each number 
        for i in range(n):
            for j in range(n):
                # find the current index of that number in the list nums
                if nums[j][0] == i:
                    # once found, extract that (indx, number) pair
                    num = nums[j]

                    # remove it from the list
                    nums.pop(j)

                    # if the value == -1 * the index, that effectively puts the value at 0
                    # or for this question, puts it at the end of the list
                    if num[1] == -j:
                        nums.append(num)
                    else:
                        # otherwise, we can insert it at (index + value) % (n-1)
                        nums.insert((j+num[1]) % (n-1), num)
                    break

    zi = -1
    for i in range(n):
        if nums[i][1] == 0:
            zi = i
            break

    a = nums[(zi + 1000) % n][1]
    b = nums[(zi + 2000) % n][1]
    c = nums[(zi + 3000) % n][1]
    return a + b + c



print(mix('input.txt'))
print(mix('input.txt', multiplier=811589153, mixes=10))
