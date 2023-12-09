

from collections import defaultdict
from dataclasses import dataclass
from functools import cmp_to_key

class HandTypes:
    Five = 7
    Four = 6
    FullHouse = 5
    Three = 4
    TwoPair = 3
    Pair = 2
    High = 1
    N = 0

@dataclass
class Hand:

    h: str
    bid: int

    normal_hand_type = HandTypes.N
    wild_hand_type = HandTypes.N


    def __post_init__(self):
        self.normal_hand_type = self._normal_hand_type()
        self.wild_hand_type = self._wild_hand_type()

    def _wild_hand_type(self):
        counts = defaultdict(int)
        wilds = 0
        for c in self.h:
            if c == 'J':
                wilds += 1
            else:
                counts[c] += 1
        
        if wilds == 0:
            return self._normal_hand_type()
        if wilds == 4 or wilds == 5:
            # if we have 4 wilds, we can easily convert this to 5 of a kind
            return HandTypes.Five
        elif wilds == 3:
            if len(counts) == 1:
                # we have a pair, convert to a 5
                return HandTypes.Five
            else:
                # we have 2 distinct cards, conert to 4+1
                return HandTypes.Four
        elif wilds == 2:
            if len(counts) == 1:
                # we have 3 of a kind, so convert to 5
                return HandTypes.Five
            elif len(counts) == 2:
                # we have a pair + a high card, so conver to 4+1
                return HandTypes.Four
            else:
                return HandTypes.Three
        elif wilds == 1:
            if len(counts) == 1:
                # 4 of a kind
                return HandTypes.Five
            elif len(counts) == 2:
                # we have a 2 pair or 3 + 1
                if any(c == 3 for c in counts.values()):
                    # currently its a three of a kind
                    return HandTypes.Four
                else:
                    # its a 2 pair
                    return HandTypes.FullHouse
            elif len(counts) == 3:
                # its a pair + 2 cards
                return HandTypes.Three
            elif len(counts) == 4:
                # its just a high card
                return HandTypes.Pair


    def _normal_hand_type(self):
        # parse hand:
        counts = defaultdict(int)
        for c in self.h:
            counts[c] += 1

        if len(counts) == 5:
            return HandTypes.High
        elif len(counts) == 4:
            # single pair
            return HandTypes.Pair
        elif len(counts) == 1:
            return HandTypes.Five
        elif len(counts) == 3:
            # could be 2 pair or three
            if any(c == 3 for c in counts.values()):
                return HandTypes.Three
            else:
                return HandTypes.TwoPair
        elif len(counts) == 2:
            # could be four of a kind or full house
            if any(c == 4 for c in counts.values()):
                return HandTypes.Four
            else:
                return HandTypes.FullHouse


def cmp_with_vals(card_val, hand_type_field):
    def cmp(h1: Hand, h2: Hand):
        h1_hand_type = getattr(h1, hand_type_field)
        h2_hand_type = getattr(h2, hand_type_field)
        if h1_hand_type < h2_hand_type:
            return -1
        elif h2_hand_type < h1_hand_type:
            return 1
        else:
            for i in range(len(h1.h)):
                h1c = card_val[h1.h[i]]
                h2c = card_val[h2.h[i]]
                if h1c < h2c:
                    return -1
                elif h2c < h1c:
                    return 1
        return 0
    return cmp
            

def parse_input(i):
    hands = []
    with open(f"{i}.txt", "r") as f:
        for l in f.readlines():
            parts = l.strip().split(' ')
            hands.append(Hand(parts[0].strip(), int(parts[1].strip())))

    return hands


def part_1(i):
    card_val = {str(p): p for p in range(2, 10, 1)}
    card_val['T'] = 10
    card_val['J'] = 11
    card_val['Q'] = 12
    card_val['K'] = 13
    card_val['A'] = 14


    hands = parse_input(i)
    cmp_func = cmp_with_vals(card_val, "normal_hand_type")
    sorted_hands = sorted(hands, key=cmp_to_key(cmp_func))

    winnings = 0
    for i, h in enumerate(sorted_hands):
        rank = i + 1
        winnings += rank * h.bid
    
    return winnings

def part_2(i):
    card_val = {str(p): p for p in range(2, 10, 1)}
    card_val['T'] = 10
    card_val['J'] = 0
    card_val['Q'] = 12
    card_val['K'] = 13
    card_val['A'] = 14

    hands = parse_input(i)

    cmp_func = cmp_with_vals(card_val, "wild_hand_type")
    sorted_hands = sorted(hands, key=cmp_to_key(cmp_func))

    winnings = 0
    for i, h in enumerate(sorted_hands):
        rank = i + 1
        winnings += rank * h.bid
    
    return winnings

print(part_1('input'))
print(part_2('input'))