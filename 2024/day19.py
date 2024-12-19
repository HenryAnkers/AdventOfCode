import os
import time


def parseInput(file_path):
    with open(file_path, "r") as file:
        return [line.strip() for line in file]


def canBeMade(pattern, substrings, memo, isPart1):
    if pattern in memo:
        return memo[pattern]

    if pattern == "":
        return 1

    ans = 0
    for substring in substrings:
        if pattern.startswith(substring):
            ans += canBeMade(pattern[len(substring) :], substrings, memo, isPart1)

            if isPart1 and ans != 0:
                return True

    memo[pattern] = ans
    return ans


def part1(data):
    substrings = data[0].split(", ")
    return sum([canBeMade(pattern, substrings, {}, True) for pattern in data[2:]])


def part2(data):
    substrings = data[0].split(", ")
    return sum([canBeMade(pattern, substrings, {}, False) for pattern in data[2:]])


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
