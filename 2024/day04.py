import os
import time

def parse_input(file_path):
    with open(file_path, 'r') as file:
        return [line.strip() for line in file]

def countXmasOccurances(x,y,data, tData):
    count = 0

    if data[y][x:x+4] == "XMAS" or data[y][x:x+4] == "SAMX":
        count += 1
    if tData[x][y:y+4] == "XMAS" or tData[x][y:y+4] == "SAMX":
        count += 1

    if y + 3 < len(data):
        if x+3 < len(data[0]):
            diag = "".join([data[y+d][x+d] for d in range(0,4)])
            if diag == "XMAS" or diag == "SAMX":
                count += 1
        if x-3 >= 0:
            diag = "".join([data[y+d][x-d] for d in range(0,4)])
            if diag == "XMAS" or diag == "SAMX":
                count += 1               

    return count

def part1(data):
    xmasCount = 0
    transposedData = ["".join(list(row)) for row in zip(*data)]
    for y, line in enumerate(data):
        for x, char in enumerate(line):
            if char == "X" or char == "S":
                xmasCount += countXmasOccurances(x,y,data, transposedData)

    return xmasCount


def countMasOccurances(x,y,data, tData):
    if y + 2 < len(data):
        hasLeft = False
        if x + 2 < len(data[0]):
            diag = "".join([data[y+d][x+d] for d in range(0,3)])
            if diag == "MAS" or diag == "SAM":
                hasLeft = True


        if hasLeft: 
            diag = "".join([data[y+d][x+2-d] for d in range(0,3)])
            if diag == "MAS" or diag == "SAM":
                return 1              

    return 0

def part2(data):
    masCount = 0
    transposedData = ["".join(list(row)) for row in zip(*data)]
    for y, line in enumerate(data):
        for x, char in enumerate(line):
            if char == "M" or char == "S":
                masCount += countMasOccurances(x,y,data, transposedData)

    return masCount


day = os.path.basename(__file__).split('.')[0].replace('day', '')
input_path = f"./inputs/day{day}.txt"

start_time = time.time()
input_data = parse_input(input_path)
end_time = time.time()
print(f"Day {day} - Parsing took {end_time - start_time:.6f} seconds")

start_time = time.time()
solution1 = part1(input_data)
end_time = time.time()
print(f"Day {day} - Part 1: {solution1} (took {end_time - start_time:.6f} seconds)")

start_time = time.time()
solution2 = part2(input_data)
end_time = time.time()
print(f"Day {day} - Part 2: {solution2} (took {end_time - start_time:.6f} seconds)")