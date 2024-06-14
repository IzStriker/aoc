from pipe import map, filter
from math import lcm

def parse_input():
    instructions = ""
    nodes: dict[str, (str, str)] = {}
    
    with open("input.txt") as file:
        instructions = file.readline().strip()
        file.readline() # skip empty line
        for node in list(file.readlines() | map(lambda l: l.strip())): 
            data = node.split(" = ")
            key = data[0]
            left, right = data[1][1:9].split(", ")
            nodes[key] = (left, right)
    
    return instructions, nodes

def part1():
    instructions, nodes = parse_input()
    
    current = "AAA"
    i = 0
    while current != "ZZZ":
        isRight = instructions[i % len(instructions)] == "R"
        current = nodes[current][int(isRight)]
        
        i += 1
    print(i)

def part2():
    instructions, nodes = parse_input()
    
    start_nodes = list(list(nodes.keys()) | filter(lambda k: k[2] == "A"))
    path_lengths: list[int] = []
    for start in start_nodes:
        i = 0
        current = start 
        while current[2] != "Z":
            isRight = instructions[i % len(instructions)] == "R"
            current = nodes[current][int(isRight)]
            i += 1
        path_lengths.append(i)
    
    print(lcm(*path_lengths))

        
            

if __name__ == '__main__':
    part2()