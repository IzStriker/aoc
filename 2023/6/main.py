import pipe
def part1():
    lines = []

    with open('input.txt', 'r') as f:
        lines = f.readlines()

    times = list(lines[0].strip().split(":")[1].split(" ") | pipe.where(lambda x: x != "") | pipe.map(lambda x: int(x)))
    distances = list(lines[1].strip().split(":")[1].split(" ") | pipe.where(lambda x: x != "") | pipe.map(lambda x: int(x)))

    waysToWin = []
    for i in range(len(times)): 
        speed = 1
        last_distance = 0
        wins = 0
        while True: 
            time = times[i] - speed 
            distance = speed * time
            if distance > distances[i]: 
                wins += 1
            if last_distance > distance and distance < distances[i]: 
                break
            last_distance = distance
            speed += 1
        waysToWin.append(wins)
    
    product = 1
    for i in waysToWin: 
        product *= i
    print(product)
            
def part2():
    lines = []

    with open('input.txt', 'r') as f:
        lines = f.readlines()

    times = int("".join(list(lines[0].split(":")[1].strip()) | pipe.where(lambda x: x.strip() != '')))
    distances = int("".join(list(lines[1].split(":")[1].strip()) | pipe.where(lambda x: x.strip() != '')))

    speed = 1
    last_distance = 0
    wins = 0
    while True: 
        time = times - speed 
        distance = speed * time
        if distance > distances: 
            wins += 1
        if last_distance > distance and distance < distances: 
            break
        last_distance = distance
        speed += 1

    print(wins)
if __name__ == "__main__":
    part2()