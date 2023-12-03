def part1(): 
    lines = []

    with open("input.txt", "r") as f:
        lines = f.readlines()
    
    map = {}
    nums = []
    partOfNums = []
    x = 0
    y = 0
    for line in lines:
        for c in line:
            if not (c == "\n" or c == "."):
                map[(x,y)] = c
            x += 1
        y += 1
        x = 0
    
    for key in map:
        if key in partOfNums:
            continue
        if str(map[key]).isnumeric():
            start = key
            end = key
            include = False
            num = ""
            # move right until end of number
            tempX = key[0]
            while (tempX, key[1]) in map:
                if  str(map[(tempX, key[1])]).isnumeric():    
                    end = (tempX, key[1])
                    num += map[end]
                else: break
                tempX += 1
            
            # from start to end, check if surrounding blocks are symbols
            while (start[0] <= end[0]):
                if start in partOfNums:
                    start = (start[0] + 1, start[1])
                    continue
                # check surrounding blocks including diagonals
                if (start[0] - 1, start[1]) in map and not str(map[(start[0] - 1, start[1])]).isnumeric():
                    include = True
                if (start[0] + 1, start[1]) in map and not str(map[(start[0] + 1, start[1])]).isnumeric():
                    include = True
                if (start[0], start[1] - 1) in map and not str(map[(start[0], start[1] - 1)]).isnumeric():
                    include = True
                if (start[0], start[1] + 1) in map and not str(map[(start[0], start[1] + 1)]).isnumeric():
                    include = True
                if (start[0] - 1, start[1] - 1) in map and not str(map[(start[0] - 1, start[1] - 1)]).isnumeric():
                    include = True
                if (start[0] + 1, start[1] + 1) in map and not str(map[(start[0] + 1, start[1] + 1)]).isnumeric():
                    include = True
                if (start[0] - 1, start[1] + 1) in map and not str(map[(start[0] - 1, start[1] + 1)]).isnumeric():
                    include = True
                if (start[0] + 1, start[1] - 1) in map and not str(map[(start[0] + 1, start[1] - 1)]).isnumeric():
                    include = True
                if(include):
                    partOfNums.append(start)
                
                
                start = (start[0] + 1, start[1])
            if(include):
                nums.append(int(num))
    
                
    print(nums)
    print(sum(nums))

def part2(): 
    lines = []

    with open("input.txt", "r") as f:
        lines = f.readlines()
    
    map = {}
    ratios = []
    # partOfNums = []
    x = 0
    y = 0
    for line in lines:
        for c in line:
            if not (c == "\n" or c == "."):
                map[(x,y)] = c
            x += 1
        y += 1
        x = 0
    
    for key in map:
        if map[key] == "*":
            partNums = []
            includedInNums = []
            nums = []
            # check surrounding blocks
            if (key[0] - 1, key[1]) in map and str(map[(key[0] - 1, key[1])]).isnumeric():
                partNums.append((key[0] - 1, key[1]))
            if (key[0] + 1, key[1]) in map and str(map[(key[0] + 1, key[1])]).isnumeric():
                partNums.append((key[0] + 1, key[1]))
            if (key[0], key[1] - 1) in map and str(map[(key[0], key[1] - 1)]).isnumeric():
                partNums.append((key[0], key[1] - 1))
            if (key[0], key[1] + 1) in map and str(map[(key[0], key[1] + 1)]).isnumeric():
                partNums.append((key[0], key[1] + 1))
            if (key[0] - 1, key[1] - 1) in map and str(map[(key[0] - 1, key[1] - 1)]).isnumeric():
                partNums.append((key[0] - 1, key[1] - 1))
            if (key[0] + 1, key[1] + 1) in map and str(map[(key[0] + 1, key[1] + 1)]).isnumeric():
                partNums.append((key[0] + 1, key[1] + 1))
            if (key[0] - 1, key[1] + 1) in map and str(map[(key[0] - 1, key[1] + 1)]).isnumeric():
                partNums.append((key[0] - 1, key[1] + 1))
            if (key[0] + 1, key[1] - 1) in map and str(map[(key[0] + 1, key[1] - 1)]).isnumeric():
                partNums.append((key[0] + 1, key[1] - 1))
            
            # get whole number
            for part in partNums:
                if part in includedInNums: continue
                start = part
                num = ""

                # move left util start of num found
                while start in map:
                    if  (start[0] - 1, start[1]) in map and  str(map[(start[0] - 1, start[1])]).isnumeric():    
                        start = (start[0] - 1, start[1])
                    else: break
                
                # move right until end of number
                while (start[0], start[1]) in map:
                    if  str(map[(start[0], start[1])]).isnumeric():    
                        num += map[start]
                        includedInNums.append(start)
                    else: break
                    start = (start[0] + 1, start[1])
                nums.append(int(num))
            if len(nums) == 2:
                ratios.append(nums[0] * nums[1])
    print(sum(ratios))
                

if __name__ == "__main__":
    part2()