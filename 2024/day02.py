import os
import time


def parse_input(file_path):
    with open(file_path, 'r') as file:
        return [line.strip() for line in file]

def parse_int_input(file_path):
    with open(file_path, 'r') as file:
        return [[int(item) for item in line.strip().split()] for line in file]

def isSafeReport(report):
    isSafe = True
    isIncreasing = report[0] < report[1]

    for val1, val2 in zip(report,report[1:]):
        if abs(val1 - val2) > 3:
            isSafe = False
            break
        if isIncreasing and val2 <= val1:
            isSafe = False
            break
        elif not isIncreasing and val1 <= val2:
            isSafe = False
            break

    return isSafe

def isSafeReportWithDampener(report):
    isSafe = True
    isIncreasing = report[0] < report[1]

    for i, vals in enumerate(zip(report,report[1:])):
        val1, val2 = vals[0], vals[1]
        if abs(val1 - val2) > 3:
            isSafe = False
        if isIncreasing and val2 <= val1:
            isSafe = False
        elif not isIncreasing and val1 <= val2:
            isSafe = False

        if not isSafe:
            isSafe = isSafeReport(report[:i] + report[i+1:]) or isSafeReport(report[:i+1] + report[i+2:]) or (i > 0 and isSafeReport(report[:i-1] + report[i:]))
            break

    return isSafe   

def part1(data):
    return sum([1 if isSafeReport(report) else 0 for report in data])

def part2(data):
    return sum([1 if isSafeReportWithDampener(report) else 0 for report in data])

day = os.path.basename(__file__).split('.')[0].replace('day', '')
input_path = f"./inputs/day{day}.txt"

start_time = time.time()
#input_data = parse_input(input_path)
input_data = parse_int_input(input_path)
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