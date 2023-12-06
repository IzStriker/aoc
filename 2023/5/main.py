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
            maps[current_map] = []
        elif(lines[i].strip() != ""):
            line = lines[i].strip().split(" ")
            source_start = int(line[0])
            dest_start = int(line[1])
            length = int(line[2])
            maps[current_map].append([source_start, dest_start, length])
    
    values = []
    for seed in seeds:
        map = "seed"
        current_value = seed
        while map in guide:
            current_map = maps[f"{map}-to-{guide[map]}"] 
            
            for value_range in current_map:
                dest_start = value_range[0]
                source_start = value_range[1]
                length = value_range[2]

                if current_value >= source_start and current_value <= source_start + length:
                    current_value = dest_start + (current_value - source_start)
                    break
            map = guide[map]
        values.append(current_value)
    
    print(min(values))


def part2():
    lines = []

    with open('input.txt', 'r') as f:
        lines = f.readlines()

    seeds = list(lines[0].strip().split(":")[1].strip().split(" ") | pipe.map(lambda x: int(x)))
    seed_ranges = []
    i = 0
    while i < len(seeds):
        seed_ranges.append((seeds[i], seeds[i] + seeds[i+1]))
        i += 2

    maps = {}
    guide = {}
    current_map = ""
    for i in range(2, len(lines)):
        if (":" in lines[i]):
            current_map = lines[i].strip().split(":")[0].split(" ")[0]
            source = lines[i].strip().split(":")[0].split("-")[0]
            dest = lines[i].strip().split(":")[0].split("-")[-1]
            guide[source] = dest.split(" ")[0]
            maps[current_map] = []
        elif(lines[i].strip() != ""):
            line = lines[i].strip().split(" ")
            source_start = int(line[0])
            dest_start = int(line[1])
            length = int(line[2])
            maps[current_map].append((source_start, dest_start, length))
    
    values = []

    map = "seed"
    ranges = seed_ranges
    while map in guide:
        current_map = maps[f"{map}-to-{guide[map]}"]
        new_ranges = []
        for value_range in current_map:
            source_start, dest_start, length = value_range
            f


if __name__ == "__main__":
    part2()