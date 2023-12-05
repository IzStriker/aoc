import re
import pipe

def part1():
    lines = []

    with open('input.txt', 'r') as f:
        lines = f.readlines()

    seeds = list(lines[0].strip().split(":")[1].strip().split(" ") | pipe.map(lambda x: int(x)))

    maps = {}
    guide = {}
    current_map = ""
    for i in range(2, len(lines)):
        if (":" in lines[i]):
            current_map = lines[i].strip().split(":")[0].split(" ")[0]
            source = lines[i].strip().split(":")[0].split("-")[0]
            dest = lines[i].strip().split(":")[0].split("-")[-1]
            guide[source] = dest.split(" ")[0]
            maps[current_map] = {}
        elif(lines[i].strip() != ""):
            line = lines[i].strip().split(" ")
            source_start = int(line[0])
            dest_start = int(line[1])
            length = int(line[2])
            for i in range(length):
                maps[current_map][dest_start + i] =  source_start + i
    
    values = []
    for seed in seeds:
        map = "seed"
        current_value = seed
        while map in guide:
            current_map = maps[f"{map}-to-{guide[map]}"] 
            if (current_value in current_map):
                current_value = current_map[current_value]
            map = guide[map]
        values.append(current_value)
    
    print(values)


    
    

if __name__ == "__main__":
    part1()