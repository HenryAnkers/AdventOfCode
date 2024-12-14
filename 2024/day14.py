import os
import time
from collections import defaultdict


def parseInput(file_path):
    with open(file_path, "r") as file:
        return [line.strip() for line in file]


def part1(data):
    answer = 0
    robotMap = defaultdict(list)

    for line in data:
        iX, iY = map(int, line.split()[0].split("=")[1].split(","))
        vX, vY = map(int, line.split()[1].split("=")[1].split(","))
        # print(f"in {iX} - {iY} - {vX} - {vY}")

        robotMap[(iX, iY)] += [(vX, vY)]

    sizeX = 101
    sizeY = 103
    for _ in range(100):
        nextMap = defaultdict(list)

        for coord in list(robotMap.keys()):
            for robot in robotMap[coord]:
                vX, vY = robot[0], robot[1]
                newX = (coord[0] + vX) % sizeX
                newY = (coord[1] + vY) % sizeY
                nextMap[(newX, newY)] += [(vX, vY)]
                # print(f"{coord} - {newX} - {newY}")

        robotMap = nextMap

    q1, q2, q3, q4 = 0, 0, 0, 0
    for coord in robotMap.keys():
        x, y = coord[0], coord[1]
        num = len(robotMap[coord])
        if x < sizeX // 2 and y < sizeY // 2:
            q1 += num
        elif x > sizeX // 2 and y < sizeY // 2:
            q2 += num
        elif x < sizeX // 2 and y > sizeY // 2:
            q3 += num
        elif x > sizeX // 2 and y > sizeY // 2:
            q4 += num

    # print(f"{q1} {q2} {q3} {q4}")
    answer = q1 * q2 * q3 * q4
    return answer


def part2(data):
    answers = []
    robotMap = defaultdict(list)

    for line in data:
        iX, iY = map(int, line.split()[0].split("=")[1].split(","))
        vX, vY = map(int, line.split()[1].split("=")[1].split(","))

        robotMap[(iX, iY)] += [(vX, vY)]

    sizeX = 101
    sizeY = 103
    for _ in range(10000):
        nextMap = defaultdict(list)

        for coord in list(robotMap.keys()):
            for robot in robotMap[coord]:
                vX, vY = robot[0], robot[1]
                newX = (coord[0] + vX) % sizeX
                newY = (coord[1] + vY) % sizeY
                nextMap[(newX, newY)] += [(vX, vY)]

        robotMap = nextMap

        q1, q2, q3, q4 = 0, 0, 0, 0
        for coord in robotMap.keys():
            x, y = coord[0], coord[1]
            num = len(robotMap[coord])
            if x < sizeX // 2 and y < sizeY // 2:
                q1 += num
            elif x > sizeX // 2 and y < sizeY // 2:
                q2 += num
            elif x < sizeX // 2 and y > sizeY // 2:
                q3 += num
            elif x > sizeX // 2 and y > sizeY // 2:
                q4 += num

        answer = q1 * q2 * q3 * q4
        answers.append(answer)
        # if answer < 100000000:
        #     print(answer)
        #     for y in range(sizeY):
        #         str = ""
        #         for x in range(sizeX):
        #             if len(robotMap[(x, y)]) != 0:
        #                 str += "O"
        #             else:
        #                 str += "-"
        #         print(str)

    return answers.index(min(answers))


day = os.path.basename(__file__).split(".")[0].replace("day", "")
input_path = f"./inputs/day{day}.txt"

start_time = time.time()
input_data = parseInput(input_path)

end_time = time.time()
print(f"Day {day} - Parsing took {end_time - start_time:.6f} seconds")
# print(input_data)

start_time = time.time()
solution1 = part1(input_data)
end_time = time.time()
print(f"Day {day} - Part 1: {solution1} (took {end_time - start_time:.6f} seconds)")

start_time = time.time()
solution2 = part2(input_data)
end_time = time.time()
print(f"Day {day} - Part 2: {solution2} (took {end_time - start_time:.6f} seconds)")
