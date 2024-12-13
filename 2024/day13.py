import math
import os
import time
import sys

from collections import defaultdict
from collections import Counter
sys.setrecursionlimit(5000)


def parseInput(file_path):
    with open(file_path, "r") as file:
        return [line.strip() for line in file]


def solve(aDelta, bDelta, prizeC, steps, aCount, bCount, memo):
    if prizeC in memo:
        return memo[prizeC]

    pX, pY = prizeC[0], prizeC[1]
    aX, aY = aDelta[0], aDelta[1]
    bX, bY = bDelta[0], bDelta[1]

    if pX == 0 and pY == 0:
        return steps

    if pX < 0 or pY < 0 or aCount > 100 or bCount > 100:
        return -1

    minSteps = -1
    bCoords = (pX - bX, pY - bY)
    bSolution = solve(aDelta, bDelta, bCoords, steps + 1, aCount, bCount + 1, memo)
    if bSolution != -1:
        minSteps = bSolution

    aCoords = (pX - aX, pY - aY)
    aSolution = solve(aDelta, bDelta, aCoords, steps + 3, aCount + 1, bCount, memo)
    if aSolution != -1:
        minSteps = min(aSolution, bSolution) if bSolution != -1 else aSolution

    memo[prizeC] = minSteps
    return minSteps


def part1(data):
    totalTokens = 0
    i = 0
    while i < len(data):
        buttonAStr, buttonBStr, prizeStr = data[i], data[i + 1], data[i + 2]
        aDelta = [int(buttonAStr.split(" ")[2].split("+")[1].replace(",", "")), int(buttonAStr.split(" ")[3].split("+")[1])]
        bDelta = [int(buttonBStr.split(" ")[2].split("+")[1].replace(",", "")), int(buttonBStr.split(" ")[3].split("+")[1])]
        prizeCoords = (int(prizeStr.split(" ")[1].split("=")[1].replace(",", "")), int(prizeStr.split(" ")[2].split("=")[1]))

        minTokens = solve(aDelta, bDelta, prizeCoords, 0, 0, 0, {})
        totalTokens += minTokens if minTokens != -1 else 0

        i += 4

    return totalTokens


def part2(data):
    totalTokens = 0
    i = 0
    while i < len(data):
        buttonAStr, buttonBStr, prizeStr = data[i], data[i + 1], data[i + 2]
        aDelta = (int(buttonAStr.split(" ")[2].split("+")[1].replace(",", "")), int(buttonAStr.split(" ")[3].split("+")[1]))
        bDelta = (int(buttonBStr.split(" ")[2].split("+")[1].replace(",", "")), int(buttonBStr.split(" ")[3].split("+")[1]))
        prizeC = (
            int(prizeStr.split(" ")[1].split("=")[1].replace(",", "")) + 10000000000000,
            int(prizeStr.split(" ")[2].split("=")[1]) + 10000000000000,
        )

        pX, pY = prizeC[0], prizeC[1]
        aX, aY = aDelta[0], aDelta[1]
        bX, bY = bDelta[0], bDelta[1]
        aPush = -1
        bPush = -1

        # aX * aPush + bX * bPush = pX
        # aY * aPush + bY * bPush = pY
        # let lcmA = lcm(aX, aY), mX * aX = lcmA, mY * aY = lcmA
        # mX * aX * aPush + mX * bX * bPush = mX * pX
        # mY * aY * aPush + mY * bY * bPush = mY * pY
        # (mX * bX - mY * bY) * bPush = mX * pX -  mY * pY
        # bPush = (mX * pX -  mY * pY)/(mX * bX - mY * bY)

        lcmA, lcmB = math.lcm(aX, aY), math.lcm(bX, bY)
        maX, maY = lcmA // aX, lcmA // aY
        mbX, mbY = lcmB // bX, lcmB // bY

        bPush = (maX * pX - maY * pY) / (maX * bX - maY * bY)
        aPush = (mbX * pX - mbY * pY) / (mbX * aX - mbY * aY)

        # why does this work? there's other non-positive non-integer solutions to these, but we always find the integer one if it exists, why?
        if float.is_integer(aPush) and float.is_integer(bPush) and aPush >= 0 and bPush >= 0:
            totalTokens += 3 * int(aPush) + int(bPush)

        i += 4

    return totalTokens


day = os.path.basename(__file__).split(".")[0].replace("day", "")
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
