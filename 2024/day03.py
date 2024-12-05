import os
import time

def parse_input(file_path):
    with open(file_path, 'r') as file:
        return [line.strip() for line in file]

def parse_int_input(file_path):
    with open(file_path, 'r') as file:
        return [[int(item) for item in line.strip().split()] for line in file]


def processMul(line, i):
    num1 = ""
    foundComma = False
    while not foundComma and i < len(line):
        if line[i] == ",":
            i += 1
            break
        elif not str.isdigit(line[i]):
            return 0
        else:
            num1 += line[i]
        i += 1

    num2 = ""
    foundEnd = False
    while not foundEnd and i < len(line):
        if line[i] == ")":
            break
        elif not str.isdigit(line[i]):
            return 0
        else:
            num2 += line[i]
        i += 1
  
    if num1 == "" or num2 == "":
        return 0

    return int(num1) * int(num2)

def part1(data):
    totalVal = 0
    lineToParse = "".join(data) + "----"
    i = 0
    while i < len(lineToParse):
        if lineToParse[i:i+4] == "mul(":
            i += 4
            tempVal = processMul(lineToParse, i)
            if tempVal != False:
                totalVal += tempVal
        else:
            i += 1

    return totalVal

def part2(data):
    totalVal = 0
    lineToParse = "".join(data) + "-------"
    i = 0
    shouldAdd = True
    while i < len(lineToParse):
        if lineToParse[i:i+4] == "do()":
            shouldAdd = True
            i += 4
        elif lineToParse[i:i+7] == "don't()":
            shouldAdd = False
            i += 7
        elif lineToParse[i:i+4] == "mul(":
            i += 4
            tempVal = processMul(lineToParse, i)
            if tempVal != False and shouldAdd:
                totalVal += tempVal
        else:
            i += 1

    return totalVal


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