from pipe import where

def part1():
    lines = []

    with open('input.txt', 'r') as f:
        lines = f.readlines()

    total = 0
    for line in lines:
        game = line.split(":")[0].strip().split(" ")[1]
        
        set1 = list(line.strip().split(":")[1].strip().split("|")[0].split(" ") | where(lambda x: x != ""))
        set2 = list(line.strip().split(":")[1].strip().split("|")[1].split(" ") | where(lambda x: x != ""))

        num = 0
        for i in set2:
            if i in set1:
                num += 1
        
        if num != 0:
            total +=  2 ** (num - 1)
    print(total)

        
def part2():
    lines = []

    with open('input.txt', 'r') as f:
        lines = f.readlines()

    games = {}
    for line in lines:
        game = int(list(line.strip().split(":")[0].strip().split(" ") | where(lambda x: x != ""))[1])
        games[game] = 1

    total = 0
    for line in lines:
        game = int(list(line.strip().split(":")[0].strip().split(" ") | where(lambda x: x != ""))[1])
        subTotal = 0
        set1 = list(line.strip().split(":")[1].strip().split("|")[0].split(" ") | where(lambda x: x != ""))
        set2 = list(line.strip().split(":")[1].strip().split("|")[1].split(" ") | where(lambda x: x != ""))

        for i in set2:
            if i in set1:
                subTotal += 1
        for i in range(games[game]):
            for j in range(1, subTotal + 1):
                games[game + j] +=  1
        total += games[game]
    
    print(total)
        

if __name__ == '__main__':
    part2()