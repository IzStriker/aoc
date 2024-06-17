from pipe import map, filter

def pairs(l):
    return [(l[i], l[i+1]) for i in range(len(l)-1)]

def get_differences(l):
    p = pairs(l)
    diffs = []
    for (a,b) in p:
        diffs.append(b-a)
    return diffs

def part1():
    lines = []
    with open("input.txt") as f:
        lines = f.readlines() | map(lambda x: list(x.strip().split(" ") | map(int)))
    
    predicted_values = []
    for line in lines: 
        differences = [line]
        current = line
        while any(list(current | map(lambda x: x != 0))):
            current = get_differences(current)
            differences.append(current)
        
        current_predicted = 0
        for i in reversed(range(len(differences))):
            current_predicted = differences[i][-1] + current_predicted
        
        predicted_values.append(current_predicted)
    print(sum(predicted_values))

def part2():
    lines = []
    with open("input.txt") as f:
        lines = f.readlines() | map(lambda x: list(x.strip().split(" ") | map(int)))
    
    predicted_values = []
    for line in lines: 
        differences = [line]
        current = line
        while any(list(current | map(lambda x: x != 0))):
            current = get_differences(current)
            differences.append(current)
        
        current_predicted = 0
        for i in reversed(range(len(differences))):
            current_predicted = differences[i][0] - current_predicted
        
        predicted_values.append(current_predicted)
    print(sum(predicted_values))

if __name__ == "__main__":
    part2()