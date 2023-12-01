from pipe import where, map

def part1():
    lines = []
    digits = []
    with open("input.txt") as f: 
        lines = f.readlines()
    
    for line in lines:
        line = list(line)
        print(line)
        first_digit = list(line | where(lambda x: x.isdigit()))[0]
        last_digit = list(reversed(line) | where(lambda x: x.isdigit()))[0]
        digits.append(first_digit + last_digit)
        
    digits = digits | map(lambda x: int(x)) 
    total = sum(digits)
    print(total)
        
def part2(): 
    lines = []
    digits = []
    with open("input.txt") as f: 
        lines = f.readlines()
    
    digitMap = {
        "one": 1,
        "two": 2,
        "three": 3,
        "four": 4,
        "five": 5,
        "six": 6,
        "seven": 7,
        "eight": 8,
        "nine": 9,
        "zero": 0,
        "0": 0,
        "1": 1,
        "2": 2,
        "3": 3,
        "4": 4,
        "5" : 5,
        "6": 6,
        "7": 7,
        "8": 8,
        "9": 9
    }
    for line in lines:
      first_index = -1
      last_index = -1
      first_num = ""
      last_num = ""

      for digit in digitMap.keys():
        
        first_index_temp = line.find(digit)
        last_index_temp = line.rfind(digit)
       
        if first_index_temp != -1 and (first_index == -1 or first_index_temp < first_index):
            first_index = first_index_temp
            first_num = digitMap[digit]
        
        if last_index_temp != -1 and (last_index == -1 or last_index_temp > last_index):
            last_index = last_index_temp
            last_num = digitMap[digit]
        
      digits.append(str(first_num) + str(last_num))
          
    
    print(list(digits))
    digits = list(digits | map(lambda x: int(x)))
    total = sum(digits)
    print("total",total)


    


if __name__ == "__main__":
    part2()