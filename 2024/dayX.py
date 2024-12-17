import os
import time
from collections import defaultdict


def parseInput(file_path):
    with open(file_path, "r") as file:
        return [line.strip() for line in file]


def parseIntInput(file_path):
    with open(file_path, "r") as file:
        return [[int(item) for item in line.strip().split()] for line in file]


def parseInputAsMap(file_path):
    with open(file_path, "r") as file:
        data = [line.strip() for line in file]
        dataMap = defaultdict(bool)
        for y, line in enumerate(data):
            for x, char in enumerate(line):
                dataMap[(x,y)] = char
        return dataMap


def part1(data):
    return None


def part2(data):
    return None


day = os.path.basename(__file__).split(".")[0].replace("day", "")
input_path = f"./inputs/day{day}.txt"

start_time = time.time()
input_data = parseInput(input_path)
# input_data = parseInputAsMap(input_path)
# input_data = parseIntInput(input_path)
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
