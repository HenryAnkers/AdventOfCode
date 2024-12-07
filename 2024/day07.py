import os
import time

def parseInput(file_path):
    with open(file_path, 'r') as file:
        return [line.strip() for line in file]
    

def checkEquation(target, currentResult, remainingNumbers, isPart1):
    if not len(remainingNumbers):
        return target == currentResult
    
    concatValue = currentResult * (10 ** len(str(remainingNumbers[0]))) + remainingNumbers[0]
    part1Ans = checkEquation(target, currentResult * remainingNumbers[0], remainingNumbers[1:], isPart1) or checkEquation(target, currentResult + remainingNumbers[0], remainingNumbers[1:], isPart1)

    return part1Ans if (isPart1 or part1Ans) else checkEquation(target, concatValue, remainingNumbers[1:], isPart1)

def part1(data):
    result = 0
    for equation in data:
        target = int(equation.split(":")[0])
        numbers = [int(item) for item in equation.split(":")[1].strip().split()]
        initialValue = numbers.pop(0)

        if checkEquation(target, initialValue, numbers, True):
            result += target

    return result


def part2(data):
    result = 0
    for equation in data:
        target = int(equation.split(":")[0])
        numbers = [int(item) for item in equation.split(":")[1].strip().split()]
        initialValue = numbers.pop(0)

        if checkEquation(target, initialValue, numbers, False):
            result += target

    return result


day = os.path.basename(__file__).split('.')[0].replace('day', '')
input_path = f"./inputs/day{day}.txt"

start_time = time.time()
input_data = parseInput(input_path)
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