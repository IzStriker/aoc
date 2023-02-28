import re
import math


class Monkey:
    def __init__(self):
        self.number = -1
        self.startItems = []
        self.operation = ""
        self.test = -1
        self.throwTrue = -1
        self.throwFalse = -1


def main():
    monkeys = read()
    # part1(monkeys)
    part2(monkeys)


def read():
    lines = []
    with open("input.txt", "r") as file:
        lines = file.readlines()

    monkeys = []
    i = 0
    while i < len(lines):
        monkey = Monkey()
        # number
        num = re.findall(r'\d+', lines[i])
        monkey.number = num[0]

        # starting items
        i += 1
        monkey.startItems = list(map(lambda x: int(x.strip()), lines[i].replace(
            "  Starting items:", "").split(",")))

        # Operation
        i += 1
        monkey.operation = lines[i].replace("  Operation: new = ", "")

        # divisible test
        i += 1
        monkey.test = int(re.findall(r'\d+', lines[i])[0])

        # pass if true
        i += 1
        monkey.throwTrue = int(re.findall(r'\d+', lines[i])[0])

        # divisible test
        i += 1
        monkey.throwFalse = int(re.findall(r'\d+', lines[i])[0])
        # clear blank line
        i += 1

        monkeys.append(monkey)
        # next iteration of looop
        i += 1
    return monkeys


def part1(monkeys: list):
    inspect = {}
    for k in range(20):
        for i in range(len(monkeys)):
            for j in range(len(monkeys[i].startItems)):
                if i not in inspect.keys():
                    inspect[i] = 1
                else:
                    inspect[i] = inspect[i] + 1

                item = monkeys[i].startItems.pop(0)
                op = monkeys[i].operation
                op = op.replace("old", str(item))
                ans = eval(op)
                ans //= 3

                if ans % monkeys[i].test == 0:
                    monkeys[monkeys[i].throwTrue].startItems.append(ans)
                else:
                    monkeys[monkeys[i].throwFalse].startItems.append(ans)
    print(inspect)


def part2(monkeys: list):
    inspect = {}
    for k in range(10000):
        mod_lcm = math.lcm(*[monkey.test for monkey in monkeys])
        print(f"{k * 100 / 10000}%")
        for i in range(len(monkeys)):
            for j in range(len(monkeys[i].startItems)):
                if i not in inspect.keys():
                    inspect[i] = 1
                else:
                    inspect[i] = inspect[i] + 1

                item = monkeys[i].startItems.pop(0)
                op = monkeys[i].operation
                op = op.replace("old", str(item))
                ans = eval(op)

                # ans %= mod_lcm
                print(ans, monkeys[i].test)
                print()
                if ans % monkeys[i].test == 0:
                    monkeys[monkeys[i].throwTrue].startItems.append(ans)
                else:
                    monkeys[monkeys[i].throwFalse].startItems.append(ans)

    values = []
    for key, value in inspect.items():
        values.append(value)

    values.sort()
    print(values)


if __name__ == "__main__":
    main()
