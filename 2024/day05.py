import os
import time
from collections import defaultdict

def parseInput(file_path):
    with open(file_path, 'r') as file:
        orderings, pages = file.read().split("\n\n")
        orderDict = defaultdict(lambda: defaultdict(lambda: False))
        for line in orderings.split("\n"):
            x = line.split("|")[0].strip()
            y = line.split("|")[1].strip()
            orderDict[int(x)][int(y)] = True # if orderDict[x][y] is true => y must come after x

        return [orderDict, [[int(item) for item in line.strip().split(",")] for line in pages.split("\n")]]


def part1(data):
    orderings, pages = data
    total = 0
    for page in pages:
        isValidPage = True
        for i, num1 in enumerate(page[:-1]):
            for _, num2 in enumerate(page[i+1:]):
                if orderings[num2][num1]: 
                    isValidPage = False
                    break
        if isValidPage:
            total += page[len(page) // 2]

    return total

def part2(data):
    orderings, pages = data
    total = 0
    for page in pages:
        i = 0
        pageHasChanged = False
        while i < len(page[:-1]):
            hasChanged = False
            num1 = page[i]
            j = i + 1
            while j < len(page):
                num2 = page[j]
                if orderings[num2][num1]: 
                    page[i] = num2
                    page[j] = num1
                    hasChanged = True
                    pageHasChanged = True
                    break
                j += 1
            if not hasChanged: 
                i += 1

        total += page[len(page) // 2] if pageHasChanged else 0

    return total


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