import os
import time
from collections import defaultdict


def parseIntInput(file_path):
    with open(file_path, "r") as file:
        return [[int(item) for item in line.strip().split()][0] for line in file]


def generateSecretNumber(currentNum):
    numToMix = currentNum * 64
    currentNum = currentNum ^ numToMix
    currentNum = currentNum % 16777216

    numToMix = currentNum // 32
    currentNum = currentNum ^ numToMix
    currentNum = currentNum % 16777216

    numToMix = currentNum * 2048
    currentNum = currentNum ^ numToMix
    currentNum = currentNum % 16777216

    return currentNum


def part1(data):
    # this is suspiciously easy...
    ans = 0
    for num in data:
        currentNum = num
        for _ in range(2000):
            currentNum = generateSecretNumber(currentNum)
        ans += currentNum
    return ans


def part2(data):
    # so we have to find the sequence of four price changes that will give us the most profit across ALL secret numbers... ffs
    # we can have 0 in the changes so there are basically 1996 * 2173 potential sequences in the bruteforce scenario
    # if the sequence ends in a negative it's not going to be the one we want, because the price must be lower than the previous price... so we can only look for positive ending sequences?
    # lets try bruteforcing like this and see how long it will take

    # 3s just for 4 inputs... as expected takes way too long for real input
    # we can just cache 'given this sequence of deltas a,b,c,d -> what is the first price we return from the sequence of secret numbers starting at num' and this should speed it up?

    ans = 0
    allPriceDeltas = set()
    secretNumbers = {}
    for num in data:
        currentNum = num
        allSequenceValues = defaultdict(int)
        prices = [int(str(num)[-1])]
        priceDeltas = []

        for i in range(2000):
            currentNum = generateSecretNumber(currentNum)
            price = currentNum % 10
            prices.append(price)
            priceDelta = prices[-1] - prices[-2]
            priceDeltas.append(priceDelta)
            if i > 2:
                lastPriceDeltas = priceDeltas[-4:]
                priceDeltaString = ",".join([str(i) for i in lastPriceDeltas])
                if not priceDeltaString in allSequenceValues:
                    allSequenceValues[priceDeltaString] = price
                if lastPriceDeltas[-1] >= 0:
                    allPriceDeltas.add(priceDeltaString)

        secretNumbers[num] = allSequenceValues

    for i, priceDelta in enumerate(allPriceDeltas):
        tmpAns = 0
        for num in data:
            tmpAns += secretNumbers[num][priceDelta]

        ans = max(ans, tmpAns)

    return ans


day = os.path.basename(__file__).split(".")[0].replace("day", "")
input_path = f"./inputs/day{day}.txt"

start_time = time.time()
input_data = parseIntInput(input_path)
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
