
from functools import cache


def parse_input(i):
    all_patterns = []
    with open(f"{i}.txt", "r") as f:
        pattern = []
        for line in f.readlines():
            line = line.strip()
            if line == "":
                all_patterns.append(pattern)
                pattern = []
            else:
                pattern.append(line)
        all_patterns.append(pattern)
    return all_patterns

def find_reflection_in_pattern(pattern):
    num_rows = len(pattern)
    num_cols = len(pattern[0])

    def _compare_vertical_reflection(i: int):
        for row in pattern:
            side_1 = row[:i][::-1]
            side_2 = row[i:]
            comparison_length = min(len(side_1), len(side_2))
            if not all(side_1[i] == side_2[i] for i in range(comparison_length)):
                return False
        return True

    def _check_for_vertical_reflection():
        # for each possible length of reflectin
        for i in range(1, num_cols):
            if _compare_vertical_reflection(i):
                return i
        return None

    def _check_for_horizontal_reflection():
        for i in range(1, num_rows):
            side_1 = pattern[:i][::-1]
            side_2 = pattern[i:]
            comparison_length = min(len(side_1), len(side_2))
            if all(side_1[i] == side_2[i] for i in range(comparison_length)):
                return i
        return None


    p = _check_for_vertical_reflection()
    if p is not None:
        print(f"Vertical reflection @ {p}")
        return p
    
    q = _check_for_horizontal_reflection()
    if q is not None:
        print(f"Horizontal reflection @ {q}")
        return q*100

    return -1
    


def find_reflection_with_smudge(pattern):
    num_rows = len(pattern)
    num_cols = len(pattern[0])

    @cache
    def _compare_h(i, comparison_length, k, smudge_used):
        if k >= comparison_length:
            return smudge_used

        side_1 = pattern[:i][::-1]
        side_2 = pattern[i:]

        s1 = side_1[k]
        s2 = side_2[k]

        mismatches = [True for i in range(len(s1)) if s1[i] != s2[i]]
        if len(mismatches) == 0:
            return _compare_h(i, comparison_length, k+1, smudge_used)
        elif len(mismatches) == 1 and not smudge_used:
            # use the smudge to flip the one and keep going
            return _compare_h(i, comparison_length, k+1, True)
        else:
            return False
            # there's 1 mismatch and we've already tried the smudge
            # or there's more than 2 mismatches

    def _check_for_horizontal_reflection():
        for i in range(1, num_rows):
            side_1 = pattern[:i]
            side_2 = pattern[i:]

            comparison_length = min(len(side_1), len(side_2))
            if _compare_h(i, comparison_length, 0, False):
                return i
        return None
    
    @cache
    def _compare_v(i, k, smudge_used):
        if k >= num_rows:
            return smudge_used
        
        side_1 = pattern[k][:i][::-1]
        side_2 = pattern[k][i:]

        mismatches = [True for j in range(min(len(side_1), len(side_2))) if side_1[j] != side_2[j]]

        if len(mismatches) == 0:
            return _compare_v(i, k+1, smudge_used)
        elif len(mismatches) == 1 and not smudge_used:
            # try it with the smudge
            return _compare_v(i, k+1, True)
        else:
            return False

        

    def _check_for_vertical_reflection():
        for i in range(1, num_cols):
            if _compare_v(i, 0, False):
                return i
        return None
    
    p = _check_for_vertical_reflection()
    if p is not None:
        return p

    q = _check_for_horizontal_reflection()
    if q is not None:
        print(f"Horizontal reflection @ {q}")
        return q * 100
            

def part_2(i):
    patterns = parse_input(i)
    t = 0
    for p in patterns:
        t += find_reflection_with_smudge(p)
    return t

def part_1(i):
    patterns = parse_input(i)
    t = 0
    for p in patterns:
        t += find_reflection_in_pattern(p)
    return t

# print(part_1("input"))
print(part_2("input"))
    
            
