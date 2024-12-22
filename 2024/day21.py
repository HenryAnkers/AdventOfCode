import os
import time
from collections import defaultdict
from collections import Counter


def parseInput(file_path):
    with open(file_path, "r") as file:
        return [line.strip() for line in file]


def part1(data):
    # this is hard
    # figure out how the robot on the numeric keypad needs to move -> first robot -> second robot -> you
    # we don't always know the optimal move to make, but it's either X then Y then Y then X (moving one unit in X then one unit in Y then one in X never makes sense)
    # sometimes one possibility isn't possible because it would take us to forbidden square, but other times both are possible but only one is optimal, and it's hard to tell which?
    # so... we try each possibility... this could grow up to 2^(path_length_n-1) but a lot of the time we will only have one choice so maybe this is OK?
    # I really wish I wrote this recursively.....

    numericKeypadList = [["7", "8", "9"], ["4", "5", "6"], ["1", "2", "3"], [False, "0", "A"]]
    directionalKeypadList = [[False, "^", "A"], ["<", "v", ">"]]
    numericKeypad = {val: (x, y) for y, row in enumerate(numericKeypadList) for x, val in enumerate(row) if val is not False}
    directionalKeypad = {val: (x, y) for y, row in enumerate(directionalKeypadList) for x, val in enumerate(row) if val is not False}
    paths = []
    answer = 0

    robotPositions = ["A" for _ in range(0, 3)]
    for code in data:
        shortestPath = ""
        for char in code:
            pathsToFind = []

            for i, robotLocation in enumerate(robotPositions):
                nextPathsToFind = []
                if i == 0:
                    currentPos = numericKeypad[robotLocation]
                    newPos = numericKeypad[char]
                    dX, dY = newPos[0] - currentPos[0], newPos[1] - currentPos[1]

                    xFirstPossible = True
                    yFirstPossible = True
                    if currentPos[1] == 3 and newPos[0] == 0:
                        xFirstPossible = False
                    elif currentPos[0] == 0 and newPos[1] == 3:
                        yFirstPossible = False

                    if xFirstPossible and abs(dX) > 0:
                        path = ""
                        for _ in range(0, abs(dX)):
                            path += "<" if dX < 0 else ">"
                        for _ in range(0, abs(dY)):
                            path += "^" if dY < 0 else "v"
                        nextPathsToFind.append(path + "A")
                    if yFirstPossible and abs(dY) > 0:
                        path = ""
                        for _ in range(0, abs(dY)):
                            path += "^" if dY < 0 else "v"
                        for _ in range(0, abs(dX)):
                            path += "<" if dX < 0 else ">"
                        nextPathsToFind.append(path + "A")
                    if dX == 0 and dY == 0:
                        newNextPaths += [p + "A" for p in nextPaths]
                    robotPositions[i] = char

                else:
                    # print(F"{i} - {pathsToFind}")
                    for pathToFind in pathsToFind:
                        currentPos = directionalKeypad[robotLocation]  # this will always be A at the start and end of a path
                        nextPaths = [""]

                        for char in pathToFind:
                            newNextPaths = []
                            newPos = directionalKeypad[char]
                            dX, dY = newPos[0] - currentPos[0], newPos[1] - currentPos[1]

                            xFirstPossible = True
                            yFirstPossible = True
                            if currentPos[1] == 0 and newPos[0] == 0:
                                xFirstPossible = False
                            elif currentPos[0] == 0 and newPos[1] == 0:
                                yFirstPossible = False

                            if xFirstPossible and abs(dX) > 0:
                                path = ""
                                for _ in range(0, abs(dX)):
                                    path += "<" if dX < 0 else ">"
                                for _ in range(0, abs(dY)):
                                    path += "^" if dY < 0 else "v"
                                newNextPaths += [p + path + "A" for p in nextPaths]
                            if yFirstPossible and abs(dY) > 0:
                                path = ""
                                for _ in range(0, abs(dY)):
                                    path += "^" if dY < 0 else "v"
                                for _ in range(0, abs(dX)):
                                    path += "<" if dX < 0 else ">"
                                newNextPaths += [p + path + "A" for p in nextPaths]
                            if dX == 0 and dY == 0:
                                newNextPaths += [p + "A" for p in nextPaths]

                            currentPos = newPos
                            nextPaths = newNextPaths
                        nextPathsToFind += nextPaths

                shortestPathLength = min([len(p) for p in nextPathsToFind])
                pathsToFind = [p for p in nextPathsToFind if len(p) == shortestPathLength]
                if i == len(robotPositions) - 1:
                    shortestPath += min(pathsToFind)

        paths.append(shortestPath)
        # print(len(shortestPath))
        answer += int(code[0] + code[1] + code[2]) * len(shortestPath)

        # <vA<A>>^Av<<A>>^AAvAA<^A>Av<A^>AAv<<A>^A>AvA^Av<<A>A^>AvA<^A>Av<<A>>^AvA^A
        # <vA<A>>^Av<<A>>^AAvAA<^A>Av<A^>AAv<<A>^A>AvA^Av<<A>A>^A<A>vA^Av<<A>>^AvA^A

    return answer


def part2(data):
    # approach just clearly doesn't work for part 2, we go 1 -> 1 -> 4 -> 64 -> 65536. can't even look for patterns with current approach
    # don't really know how to narrow the number of paths we check down more than we already are?
    # we basically need a way of knowing if the X or Y move first is better without trying all possibilites, which is what we initially tried to do for part 1 and failed

    # there are only three ambiguous moves (and their opposites, so 6), (A <-> V), (> <-> ^) and (A <--> <) where we can go either X or Y first
    # I realise now that going < V < for A to < might be a valid move, so how can my part one solution work when I will only consider V < < !?!?!

    # A -> V
    # we can go either < V or V <
    # on the next keypad, starting at A, this doesn't seem to make a difference as we have to pass both < and V anyway
    # (V)   --> (<) V A   ---> (V < < A) (> A) (> ^ A)
    # (V)   --> (V) < A   ---> (V < A) (< A) (> > ^ A)

    # V -> A
    # > ^ or ^ >
    # on the next...
    # (>) ^ A  -->  (V A) (< ^ / ^ <)A  (> A)
    # we don't want to do < ^ because in our current system we won't try the shortest path if we move left first... wtf... can we just try using always ^ <
    # (^) > A  -->  (< A) (V > / > V)A  (^ A)
    # these two are inconsequential because you need to perform the same moves no matter what, so just pick V >

    #  ^ to >
    # (V >) or (> V) - this doesnt matter because you visit both in the path in same number of steps, same as A --> V

    # > to ^
    # same as above, < ^ or ^ <, we said ^ <

    # A to < / < to A
    # < V < is obviously terrible for the next keypad... so of course it wont get picked

    # so it looks like pick the Y axis move always if the Y and X axis moves are available?
    # this still works and is significantly faster but now hangs at 20. there must be something we can cache?

    # we always make some set of moves and then go back to A... we can split the current path up by "A" and figure out given this path what the next one looks like
    # then memo[path+"A"] = result, move on to the next part... eventually we hope to cache a lot of the list
    # this still isn't perfect but might be fast enough to work? no lol its still slow as fuck.

    numericKeypadList = [["7", "8", "9"], ["4", "5", "6"], ["1", "2", "3"], [False, "0", "A"]]
    directionalKeypadList = [[False, "^", "A"], ["<", "v", ">"]]
    numericKeypad = {val: (x, y) for y, row in enumerate(numericKeypadList) for x, val in enumerate(row) if val is not False}
    directionalKeypad = {val: (x, y) for y, row in enumerate(directionalKeypadList) for x, val in enumerate(row) if val is not False}
    paths = []
    answer = 0
    pathLengths = []
    memo = {"A": "A"}

    robotPositions = ["A" for _ in range(0, 3 + 23)]
    for code in data:
        shortestPath = ""
        for char in code:
            pathsToFind = []

            for i, robotLocation in enumerate(robotPositions):
                nextPathsToFind = []
                if i == 0:
                    currentPos = numericKeypad[robotLocation]
                    newPos = numericKeypad[char]
                    dX, dY = newPos[0] - currentPos[0], newPos[1] - currentPos[1]

                    xFirstPossible = True
                    yFirstPossible = True
                    if currentPos[1] == 3 and newPos[0] == 0:
                        xFirstPossible = False
                    elif currentPos[0] == 0 and newPos[1] == 3:
                        yFirstPossible = False

                    if xFirstPossible and abs(dX) > 0:
                        path = ""
                        for _ in range(0, abs(dX)):
                            path += "<" if dX < 0 else ">"
                        for _ in range(0, abs(dY)):
                            path += "^" if dY < 0 else "v"
                        nextPathsToFind.append(path + "A")
                    if yFirstPossible and abs(dY) > 0:
                        path = ""
                        for _ in range(0, abs(dY)):
                            path += "^" if dY < 0 else "v"
                        for _ in range(0, abs(dX)):
                            path += "<" if dX < 0 else ">"
                        nextPathsToFind.append(path + "A")
                    if dX == 0 and dY == 0:
                        nextPathsToFind.append("A")
                    robotPositions[0] = char

                else:
                    # print()
                    for pathToFind in pathsToFind:
                        print(f"{i}")
                        currentPos = directionalKeypad[robotLocation]  # this will always be A at the start and end of a path
                        nextPath = ""

                        moves = pathToFind.split("A")
                        for q, move in enumerate(moves):
                            if q != len(moves) - 1:
                                move += "A"
                            movePath = ""
                            if move in memo:
                                nextPath += memo[move]
                                continue
                            else:
                                for char in move:
                                    nextPathPart = ""
                                    newPos = directionalKeypad[char]
                                    dX, dY = newPos[0] - currentPos[0], newPos[1] - currentPos[1]

                                    yFirstPossible = True
                                    if currentPos[0] == 0 and newPos[1] == 0:
                                        yFirstPossible = False

                                    if yFirstPossible and abs(dY) > 0:
                                        for _ in range(0, abs(dY)):
                                            nextPathPart += "^" if dY < 0 else "v"
                                        for _ in range(0, abs(dX)):
                                            nextPathPart += "<" if dX < 0 else ">"
                                    else:
                                        for _ in range(0, abs(dX)):
                                            nextPathPart += "<" if dX < 0 else ">"
                                        for _ in range(0, abs(dY)):
                                            nextPathPart += "^" if dY < 0 else "v"
                                    nextPathPart += "A"
                                    currentPos = newPos
                                    movePath += nextPathPart

                            memo[move] = movePath

                            nextPath += movePath
                        nextPathsToFind.append(nextPath)
                        # print(nextPathsToFind)

                        # path = pathToFind.split("A")
                        # newPathToFind = "A".join(path)
                        # for char in newPathToFind:
                        #     nextPathPart = ""
                        #     newPos = directionalKeypad[char]
                        #     dX, dY = newPos[0] - currentPos[0], newPos[1] - currentPos[1]

                        #     yFirstPossible = True
                        #     if currentPos[0] == 0 and newPos[1] == 0:
                        #         yFirstPossible = False

                        #     if yFirstPossible and abs(dY) > 0:
                        #         for _ in range(0, abs(dY)):
                        #             nextPathPart += "^" if dY < 0 else "v"
                        #         for _ in range(0, abs(dX)):
                        #             nextPathPart += "<" if dX < 0 else ">"
                        #     else:
                        #         for _ in range(0, abs(dX)):
                        #             nextPathPart += "<" if dX < 0 else ">"
                        #         for _ in range(0, abs(dY)):
                        #             nextPathPart += "^" if dY < 0 else "v"
                        #     nextPathPart += "A"

                        #     currentPos = newPos
                        #     nextPath += nextPathPart
                        # nextPathsToFind.append(nextPath)

                shortestPathLength = min([len(p) for p in nextPathsToFind])
                pathsToFind = [p for p in nextPathsToFind if len(p) == shortestPathLength]
                pathLengths.append(len(pathsToFind[0]))
                # if len(pathLengths) > 12:
                #     print(F"{(pathLengths[-1] - pathLengths[1]) / (pathLengths[-2] - pathLengths[1])}")
                if i == len(robotPositions) - 1:
                    shortestPath += min(pathsToFind)
                    # print(f" len {len(pathsToFind)}")

        paths.append(shortestPath)
        print(len(shortestPath))
        answer += int(code[0] + code[1] + code[2]) * len(shortestPath)

    return answer


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
