def part1():
    lines = []
    possibleGames = []
    with open('input.txt', 'r') as f:
        lines = f.readlines()
    max = {
        'red': 12,
        'green': 13,
        'blue': 14
    }

    for line in lines: 
        possible = True
        game = line.split(":")[0].split(" ")[1]
        
        rounds = line.split(":")[1].split(";")
        
        for round in rounds:
            colors = round.split(",")
            for color in colors:
                num = color.strip().split(" ")[0]
                colorId = color.strip().split(" ")[1]
                if max[colorId] < int(num):
                    possible = False
        if possible:
            possibleGames.append(int(game))
    print(sum(possibleGames))
    
def part2():
    lines = []
    gamePowers = []
    with open('input.txt', 'r') as f:
        lines = f.readlines()
  

    for line in lines: 
        game = line.split(":")[0].split(" ")[1]
        rounds = line.split(":")[1].split(";")

        highestColor= {}

        for round in rounds:
            colors = round.split(",")
            for color in colors:
                num = color.strip().split(" ")[0]
                colorId = color.strip().split(" ")[1]
                if colorId in highestColor:
                    if highestColor[colorId] < int(num):
                        highestColor[colorId] = int(num)
                else:
                    highestColor[colorId] = int(num)
        gamePower = 1
        for color in highestColor:
            gamePower *= highestColor[color]
        gamePowers.append(gamePower)

    print(sum(gamePowers))

if __name__ == '__main__':
    part2()