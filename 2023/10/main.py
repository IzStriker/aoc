from pipe import map, filter, sort

opposite = {"n": "s", "e": "w", "s": "n", "w": "e"}
pipe_types = {
    "|": "ns",
    "-": "ew",
    "L": "ne",
    "J": "nw",
    "7": "sw",
    "F": "se",
    "S": "nesw",
    ".": "",
}


def part1():
    maze = []
    with open("input.txt") as f:
        maze = list(f.readlines() | map(lambda x: x.strip()))

    distances = find_distances(maze, find_start(maze))
    max_distance = -1
    for _, v in distances.items():
        if v > max_distance:
            max_distance = v

    print(max_distance)


def part2():
    maze = []
    with open("2023/10/input.txt") as f:
        maze = list(f.readlines() | map(lambda x: x.strip()))

    start = find_start(maze)
    distances = find_distances(maze, start)
    replace_start(maze, start)
    print(internal_cells(maze, distances))


def find_start(maze: list[str]) -> tuple[int, int]:
    for y in range(len(maze)):
        for x in range(len(maze[y])):
            if maze[y][x] == "S":
                return x, y
    return -1, -1


def replace_start(maze: list[list[str]], start: tuple[int, int]):
    directions = ""
    for transform in "nesw":
        if is_connected(maze, start, transform_coords(start, transform)):
            directions += transform

    assert len(directions) == 2

    val = list(
        list(pipe_types.items())
        | filter(
            lambda x: directions[0] in x[1] and directions[1] in x[1] and x[0] != "S"
        )
    )[0][0]

    maze[start[1]] = maze[start[1]].replace("S", val)


def find_distances(
    maze: list[str], start_position: tuple[int, int]
) -> dict[tuple[int, int], int]:
    distances: dict[tuple[int, int], int] = {}

    distances[start_position] = 0
    # ((x,y), distance)
    working_set = [(start_position, 0)]
    while len(working_set) > 0:
        (current_x, current_y), current_distance = working_set.pop()

        directions = pipe_types[maze[current_y][current_x]]

        for d in directions:
            new_coord = transform_coords((current_x, current_y), d)
            new_distance = current_distance + 1
            if is_connected(
                maze, (current_x, current_y), new_coord
            ) and update_distance(distances, new_coord, new_distance):
                working_set.append((new_coord, new_distance))

    return distances


def update_distance(
    distances: dict[tuple[int, int], int], coord: tuple[int, int], distance: int
) -> bool:
    if (not coord in distances) or distances[coord] > distance:
        distances[coord] = distance
        return True
    return False


def is_connected(
    maze: list[str], coord1: tuple[int, int], coord2: tuple[int, int]
) -> bool:
    x1, y1 = coord1
    x2, y2 = coord2

    direction = ""
    if y1 > y2:
        direction = "n"
    elif y1 < y2:
        direction = "s"
    elif x1 > x2:
        direction = "w"
    elif x1 < x2:
        direction = "e"

    assert direction != ""

    return (
        direction in pipe_types[maze[y1][x1]]
        and opposite[direction] in pipe_types[maze[y2][x2]]
    )


def transform_coords(coords: tuple[int, int], transformation: str) -> tuple[int, int]:
    x, y = coords
    if transformation == "n":
        return x, y - 1
    if transformation == "s":
        return x, y + 1
    if transformation == "w":
        return x - 1, y
    if transformation == "e":
        return x + 1, y


def internal_cells(maze: list[list[str]], perimeter: dict[tuple[int, int], int]) -> int:
    cells = []

    for y in range(len(maze)):
        inside = False
        last_corner_type = ""
        for x in range(len(maze[y])):
            if y == 5:
                pass
            if (x, y) in perimeter:

                if (
                    (maze[y][x] == "|")
                    or (maze[y][x] == "7" and last_corner_type == "up")
                    or (maze[y][x] == "J" and last_corner_type == "down")
                ):
                    inside = not inside
                    last_corner_type = ""
                elif maze[y][x] == "F" or maze[y][x] == "L":
                    last_corner_type = "up" if maze[y][x] == "L" else "down"

            elif inside:
                cells.append((x, y))

    return len(cells)


if __name__ == "__main__":
    part2()
