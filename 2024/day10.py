import os
import time
from collections import defaultdict

def parseInputAsMap(file_path):
    with open(file_path, 'r') as file:
        data = [line.strip() for line in file]
        dataMap = defaultdict(lambda: defaultdict(lambda: -1))
        for y, line in enumerate(data):
            for x, char in enumerate(line):
                dataMap[x][y] = int(char) if char != "." else -1

        return dataMap

def findTrailheads(datamap, x, y, nextInt, currPath, part1):
    if datamap[x][y] != nextInt:
        return set()

    if datamap[x][y] == 9:
        key = (x,y) if part1 else (x,y,currPath)
        return set([key])
    
    nextInt += 1
    return findTrailheads(datamap, x+1, y, nextInt, currPath + "A", part1) | findTrailheads(datamap, x-1, y, nextInt, currPath + "B", part1) | findTrailheads(datamap, x, y+1, nextInt, currPath + "C", part1) | findTrailheads(datamap, x, y-1, nextInt, currPath + "D", part1)

        
def part1(data):
    answer = 0
    for x in list(data.keys()):
        for y in list(data[x].keys()):
            if data[x][y] == 0:
                answer += len(findTrailheads(data, x, y, 0, "", True))
    
    return answer

def part2(data):
    answer = 0
    for x in list(data.keys()):
        for y in list(data[x].keys()):
            if data[x][y] == 0:
                answer += len(findTrailheads(data, x, y, 0, "", False))
    
    return answer


day = os.path.basename(__file__).split('.')[0].replace('day', '')
input_path = f"./inputs/day{day}.txt"

start_time = time.time()
input_data = parseInputAsMap(input_path)
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