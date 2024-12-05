import os
import time
import collections as c

def parse_input(file_path):
    with open(file_path, 'r') as file:
        return [[int(item) for item in line.strip().split()] for line in file]

def part1(data):
    totalDistance = 0
    leftNumbers = sorted([line[0] for line in data])
    rightNumbers = sorted([line[1] for line in data])
    for a,b in zip(leftNumbers, rightNumbers):
        totalDistance += abs(a-b)

    return totalDistance

def part2(data):
    totalSim = 0
    leftNumbers = [line[0] for line in data]
    rightNumbers = [line[1] for line in data]
    rightCount = c.Counter(rightNumbers)
    for num in leftNumbers:
        totalSim += num * rightCount[num]

    return totalSim


day = os.path.basename(__file__).split('.')[0].replace('day', '')
input_path = f"./inputs/day{day}.txt"

start_time = time.time()
input_data = parse_input(input_path)
end_time = time.time()
print(f"Day {day} - Parsing took {end_time - start_time:.6f} seconds")
#print(input_data)

start_time = time.time()
solution1 = part1(input_data)
end_time = time.time()
print(f"Day {day} - Part 1: {solution1} (took {end_time - start_time:.6f} seconds)")

start_time = time.time()
solution2 = part2(input_data)
end_time = time.time()
print(f"Day {day} - Part 2: {solution2} (took {end_time - start_time:.6f} seconds)")