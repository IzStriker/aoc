from os import path


def main():
    input = []
    with open("input.txt", "r") as file:
        input = file.readlines()

    # part1(input)
    part2(input)


def part1(input):
    currentPath = ""
    sizes = {}

    for line in input:
        if "$" in line:
            if "ls" in line:
                continue
            dir_op = line.replace("$ cd", "").strip()
            if "/" in dir_op:
                currentPath = "\\"
            elif ".." in dir_op:
                oldPath = currentPath
                currentPath = path.dirname(currentPath)

                if oldPath not in sizes.keys():
                    sizes[oldPath] = 0

                sizes[currentPath] += sizes[oldPath]
            else:
                currentPath = path.join(currentPath, dir_op)
        else:
            if not currentPath in sizes.keys():
                sizes[currentPath] = 0

            if not "dir" in line:
                sizes[currentPath] += int(line.split(" ")[0])

    total = 0
    for key, value in sizes.items():
        if value < 100000:
            total += value
    # print(sizes)
    print(total)


def part2(input):
    currentPath = ""
    sizes = {"\\": 0}
    total = 0
    for i, line in enumerate(input):
        if "$" in line:
            if "ls" in line:
                continue
            dir_op = line.replace("$ cd", "").strip()
            if "/" in dir_op:
                currentPath = "\\"
            elif ".." in dir_op:
                oldPath = currentPath
                currentPath = path.dirname(currentPath)
                sizes[currentPath] += sizes[oldPath]
            else:
                currentPath = path.join(currentPath, dir_op)
        else:
            if not currentPath in sizes.keys():
                sizes[currentPath] = 0

            if "dir" not in line:
                sizes[currentPath] += int(line.split(" ")[0])
                total += int(line.split(" ")[0])

        print(i)
        if i == len(input) - 1:
            while currentPath != "\\":
                oldPath = currentPath
                currentPath = path.dirname(currentPath)
                sizes[currentPath] += sizes[oldPath]

    totalDiskSpace = 70_000_000
    spaceRequired = 30_000_000
    totalFree = (totalDiskSpace - sizes["\\"])
    needed = spaceRequired - totalFree
    min = sizes["\\"] * 2

    print(sizes)
    print(total)
    print(len(input))
    for key, value in sizes.items():
        if value >= needed and value < min:
            min = value

    print(min)


if __name__ == "__main__":
    main()
