def main():
    file = []
    with open("input.txt", "r") as f:
        file = f.readlines()

    part1(file, gen_stacks())
    part2(file, gen_stacks())


def gen_stacks():

    file = []
    stacks = {}
    with open("input.txt", "r") as f:
        file = f.readlines()

    for line in file:
        if any(char.isdigit() for char in line):
            for i in range(len(list(filter(lambda x: x.isnumeric(), line.split(" "))))):
                stacks[str(i+1)] = []
            break

    rawStack = []
    for line in file:
        if any(char.isdigit() for char in line):
            break
        current = []
        line_iter = iter(line)
        for char in line_iter:
            if char == " ":
                current.append("")
                next(line_iter, None)
                next(line_iter, None)
                next(line_iter, None)
                continue
            if char == "[":
                continue

            current.append(char)
            next(line_iter, None)
            next(line_iter, None)
            next(line_iter, None)

        rawStack.append(current)
        current = []

    for i in range(len(rawStack), 0, -1):
        for j in range(len(rawStack[i - 1])):
            if not rawStack[i - 1][j] == "":
                stacks[str(j+1)].append(rawStack[i - 1][j])
    return stacks


def part1(file, stacks):

    for line in filter(lambda x: "move" in x, file):
        orders = line.split(" ")
        for i in range(int(orders[1])):
            temp = stacks[orders[3]].pop()
            stacks[orders[5].strip()].append(temp)

    ans = ""
    for i in range(len(stacks)):
        ans += stacks[str(i + 1)].pop()
    print(ans)


def part2(file, stacks):
    for line in filter(lambda x: "move" in x, file):
        orders = line.split(" ")
        temp = stacks[orders[3]][len(stacks[orders[3]]) - int(orders[1]):]
        del stacks[orders[3]][len(stacks[orders[3]]) - int(orders[1]):]
        stacks[orders[5].strip()].extend(temp)

    ans = ""
    for i in range(len(stacks)):
        ans += stacks[str(i + 1)].pop()
    print(ans)


if __name__ == "__main__":
    main()
