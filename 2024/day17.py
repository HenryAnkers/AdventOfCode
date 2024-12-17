import os
import time
from collections import defaultdict


def parseInput(file_path):
    with open(file_path, "r") as file:
        return [line.strip() for line in file]


def solve(a, b, c, program, target):
    def getComboOp(int, a, b, c):
        if int == 7:
            raise Exception("uh oh")
        if int <= 3:
            return int
        if int == 4:
            return a
        if int == 5:
            return b
        if int == 6:
            return c
        return -1

    output = []
    targetArr = target.split(",")

    i = 0
    while i < len(program):
        instruction = program[i]
        operand = program[i + 1]
        comboOp = getComboOp(operand, a, b, c)

        if instruction == 0:
            a = int(a // (2**comboOp))
        elif instruction == 6:
            b = int(a // (2**comboOp))
        elif instruction == 7:
            c = int(a // (2**comboOp))
        elif instruction == 1:
            b = b ^ operand
        elif instruction == 2:
            b = comboOp % 8
        elif instruction == 3:
            if a != 0:
                i = operand
                continue
        elif instruction == 4:
            b = b ^ c
        elif instruction == 5:
            output.append(str(comboOp % 8))
            if target != "" and (len(output) > len(targetArr) or targetArr[len(output) - 1] != output[-1]):
                return -1
        i += 2

    return ",".join(output)


def part1(data):
    a = int(data[0].split(": ")[1])
    # a = 136904920099226
    b = 0
    c = 0
    program = [int(i) for i in data[4].split(": ")[1].split(",")]

    return solve(a, b, c, program, "")


def part2(data):
    # what values of A will produce 2 as the first output?
    # 3 - 0
    # 2 - 1
    # 1 - 2
    # 0 - 3
    # 5 - 4
    # 3 - 5
    # 5 - 6
    # 5 - 7
    # 3,2 - 8
    # 2,2 - 9

    # always sets B to a%8 at the start
    # sets B to (a%8)**5
    # sets C to A // (2**B)
    # B = B ^ 6
    # A = A // 8
    # B = B ^ C
    # (((A%8) ** 5) ^ 6) ^ (A // 2 ** ((A%8)**5)) = X where X % 8 == 2 ?

    # A = A // 8 => 1 outputs 1 number, 8 outputs 2, 16 outputs 3... so our number must be between 8**(proglength) and 8 ** (proglength + 1)... this doesn't narrow it down much

    # what happens after we do A // 8? we're effectively removing the last 3 bits from A.
    # B = a % 8 -> we only care about these last 3 bits when calculating the number?
    # well c = A // 2**[0-7] so c can be huge still and it will depend on the whole number

    # trying to bruteforce and look for patterns:
    # 2 - 1
    # 2,4,1,5 - 2203
    # 2,4,1,5,7,5,1,6,0,3 - 690055578

    # maybe I should look the other way? what outputs 3 -> 3,0 -> 5,3,0, because the first is outputted by A, the second by A // 8.. so we should try building it in reverse
    # 5,5,3,0 - 1538
    # 6,5,5,3,0 - 12691
    # 4,6,5,5,3,0 - 101659

    # looking at a different approach, if I multiply 690055578 by 8 I get 5,2,4,1,5,7,5,1,6,0,3. Add 1 to it and I get 7,2,4,1,5,7,5,1,6,0,3... this is obvious significant but I have no idea what it means
    # can we just bruteforce this by starting with 3, then (3*8) + D1, then (3*8 + D1) * 8 + D2... this looks like base8 so D < 8?
    # ok maybe we can't because we get to (3 * 8) * 8) * 8 + 2) and then adding (0,7) doesn't get you 6,5,5,3,0
    # because not all values work, e.g. 1586 works to transform into 12691 but 1538 doesnt
    # maybe it's actually in the form n_i * 8 + d_i for every step? and we need to try every step in this form to find the lowest that can get us to the next?
    # then once we have all the steps to make "3,0", we can find every step used to make "5,3,0" in the same way... etc
    # this works, I guess the program is some sort of base 8 encoding of the number, but still don't really understand how/why

    program = [int(i) for i in data[4].split(": ")[1].split(",")]
    valuesToFind = program[:-2]
    target = "3,0"
    validInputs = [3]  # this assumes your program ends with 3,0. 3*8 is the only way to get the output 0.
    while len(valuesToFind) >= 0:
        newValidInputs = []

        for input in validInputs:
            for i in range(0, 8):
                for j in range(0, 8):
                    if solve(i * 8 * input + j, 0, 0, program, "") == target:
                        newValidInputs.append(i * 8 * input + j)

        validInputs = newValidInputs
        if len(valuesToFind) == 0:
            return min(validInputs)

        target = str(valuesToFind.pop(-1)) + "," + target


day = os.path.basename(__file__).split(".")[0].replace("day", "")
input_path = f"./inputs/day{day}.txt"

start_time = time.time()
input_data = parseInput(input_path)
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
