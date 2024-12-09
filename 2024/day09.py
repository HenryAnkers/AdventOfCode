import os
import time
from collections import defaultdict

def parseInput(file_path):
    with open(file_path, 'r') as file:
        return [line.strip() for line in file]

def part1(data):
    answer = 0
    currentId = 0
    currentIndex = 0
    isFreeSpace = False
    dataMap = defaultdict(int)
    freeSpaceIndices = []

    for digit in data[0]:
        if not isFreeSpace:
            for i in range(currentIndex, currentIndex + int(digit)):        
                dataMap[i] = currentId
            currentId += 1
        else:
            for i in range(currentIndex, currentIndex + int(digit)):        
                freeSpaceIndices.append(i)
        isFreeSpace = not isFreeSpace
        currentIndex += int(digit)

    for dataIndex in reversed(sorted(dataMap.keys())):
        if len(freeSpaceIndices) > 0:
            nextFreeSpace = freeSpaceIndices.pop(0)
            if nextFreeSpace < dataIndex:
                dataMap[nextFreeSpace] = dataMap[dataIndex]
                del dataMap[dataIndex]
            else:
                break
        
    for index in dataMap.keys():
        answer += index * dataMap[index]

    return answer

def part2(data):
    answer = 0
    currentId = 0
    currentIndex = 0
    isFreeSpace = False
    dataBlocks = []
    freeSpaceBlocks = []

    for digit in data[0]:
        if not isFreeSpace:
            dataBlocks.append((currentIndex, currentIndex + int(digit) - 1, currentId))
            currentId += 1
        else:
            freeSpaceBlocks.append((currentIndex, currentIndex + int(digit) - 1))
        isFreeSpace = not isFreeSpace
        currentIndex += int(digit)

    newDataBlocks = []
    for data in reversed(dataBlocks):
        dataLength = data[1] - data[0] + 1
        dataId = data[2]

        suitableFreeSpaceIndex = -1
        for i, freeSpace in enumerate(freeSpaceBlocks):
            if freeSpace[0] < data[0] and (freeSpace[1] - freeSpace[0] + 1) >= dataLength:
                suitableFreeSpaceIndex = i
                break
        
        if suitableFreeSpaceIndex != -1:
            freeSpace = freeSpaceBlocks[suitableFreeSpaceIndex]
            newDataBlocks.append((freeSpace[0], freeSpace[0] + dataLength - 1, dataId))
            if freeSpace[1] >= freeSpace[0] + dataLength:
                freeSpaceBlocks[suitableFreeSpaceIndex] = ((freeSpace[0] + dataLength, freeSpace[1]))
            else:
                del freeSpaceBlocks[suitableFreeSpaceIndex]
        else:
             newDataBlocks.append(data)
        
    for data in newDataBlocks:
        for i in range(data[0], data[1] + 1):
            answer += i * data[2]
    return answer


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