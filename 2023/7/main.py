from pipe import map
from copy import deepcopy
from functools import cmp_to_key
        

def compare_sets(part1: bool) -> int:
    def F(set1: list, set2: list) -> int:
        card_rank = ["A", "K", "Q"]  + (["J"] if part1 else [])  + ["T", "9", "8", "7", "6", "5", "4", "3", "2"] + (["J"] if not part1 else [])
        
        set1_type = get_set_type(set1[0], part1)
        set2_type = get_set_type(set2[0], part1)
        
        if set1_type > set2_type:
            return -1
        elif set1_type < set2_type:
            return 1
        
        for i in range(5):
            if card_rank.index(set1[0][i]) > card_rank.index(set2[0][i]):
                return -1
            elif card_rank.index(set1[0][i]) < card_rank.index(set2[0][i]):
                return 1
            
        return 0
    return F


def get_set_type(set: str, part1: bool) -> int:
    type_rank = [[5], [4], [3,2], [3], [2,2], [2], [0]]
    cards_matching: dict[str, int] = {}
    for card in set:
        if card in cards_matching: 
            cards_matching[card] += 1
        else: 
            cards_matching[card] = 1
    
    if not part1 and set != "JJJJJ" and "J" in cards_matching:
        js = cards_matching["J"]
        del cards_matching["J"]
        most_common_key = ""
        most_common_value = -1
        for k, v in cards_matching.items():
            if v > most_common_value:
                most_common_key = k
                most_common_value = v
        if most_common_key in cards_matching: cards_matching[most_common_key] += js
    
    for rank, required_matches in enumerate(deepcopy(type_rank)):
        for v in list(cards_matching.values()):
            if v in required_matches:
                required_matches.remove(v)
            if len(required_matches) == 0:
                return rank
            
    return len(type_rank) - 1
    
            
def part1():
    cards: list[int] = []
    
    with open ("input.txt", "r") as f:
        cards = list(f.readlines() | map(lambda x: x.strip().split(" ")))
    
    cards.sort(key=cmp_to_key(compare_sets(True)))
    
    total = 0
    for i in range(len(cards)):
        [_,bid] = cards[i]
        total += (i + 1) * int(bid)
       
    print(total)
    
def part2():
    cards: list[int] = []
    
    with open ("input.txt", "r") as f:
        cards = list(f.readlines() | map(lambda x: x.strip().split(" ")))
    
    cards.sort(key=cmp_to_key(compare_sets(False)))
    
    total = 0
    for i in range(len(cards)):
        [card,bid] = cards[i]
        total += (i + 1) * int(bid)
        # print(card, bid)
       
    print(total)
if __name__ == "__main__":
    part2()