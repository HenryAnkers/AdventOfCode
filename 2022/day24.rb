class Day24
  def initialize
    input = File.read('input/day24.txt').split("\n")

    puts solve(input)
  end

  def updateBlizzardMap(currentMap, maxX, maxY)
    newMap = Hash.new{ |h,k| h[k] = []}

    currentMap.each do |k, blizzards|
      x, y = k
      blizzards.each do |b|
        xd, yd = b
        newX, newY = x+xd, y+yd
        if newX >= maxX
          newMap[[1, y+yd]] << b
        elsif newX <= 0
          newMap[[maxX-1, y+yd]] << b

        elsif newY >= maxY
          newMap[[x+xd, 1]] << b
        elsif newY <= 0
          newMap[[x+xd, maxY-1]] << b

        else
          newMap[[x+xd, y+yd]] << b
        end
      end
    end

    newMap
  end

  def parseInput(input)


    blizzardMap = Hash.new([])
    input.each_with_index do |line, y|
      line.split("").each_with_index do |char, x|
        if char == "^"
          blizzardMap[[x,y]] = [[0,-1]]
        elsif char == ">"
          blizzardMap[[x,y]] = [[1,0]]
        elsif char == "v"
          blizzardMap[[x,y]] = [[0,1]]
        elsif char == "<"
          blizzardMap[[x,y]] = [[-1,0]]
        end
      end
    end

    return [1,0], [input[0].length-2, input.length-1], blizzardMap
  end

  def DFS(location, finishLocation, time, blizzardMaps, visited, maxX, maxY, bestTime) # BFS is obvious answer but I've done it a bunch of times so wanted to try DFS, awkward question for it though
    if time >= bestTime[0] || time > 1000 # issue with DFS was dealing with the stupid cycles where you want to go back 1 unit or pause, easy to get stuck in loops. without this bestTime hack it won't work (prune any stack that goes over 1k in length, sufficient for my answer)
      return -1
    end

    x,y = location

    if visited[[x,y,time]]
      return -1 # if we've already visited it in previously at the same time, dont do it again
    end
    visited[[x,y,time]] = true

    if !blizzardMaps.key?(time)
      blizzardMaps[time] = updateBlizzardMap(blizzardMaps[time-1], maxX, maxY) # assume we always have it
    end

    blizzardMap = blizzardMaps[time] # blizzard doesnt even look like a word anymore

    if blizzardMap.key?(location) || x <= 0 || y < 0 || (y == 0 && x != 1) || (y == maxY && x != maxX-1) || x >= maxX || y > maxY
      return -1
    end

    deltas = []
    if finishLocation[0] >= location[0] || finishLocation[1] >= location[0]
      deltas = [[0,1],[1,0],[0,-1],[-1,0],[0,0]]  # prefer traveling in a certain direction.
    else
      deltas = [[-1,0],[0,-1],[1,0],[0,1],[0,0]]
    end

    potentialAnswers = []
    deltas.each do |d|
      xd, yd = d
      newX, newY = x+xd, y+yd

      if [newX,newY] == finishLocation
        if time < bestTime[0]
          bestTime[0] = time + 1
        end

        return time + 1
      end

      potentialAns =  DFS([newX,newY],finishLocation,time+1,blizzardMaps,visited,maxX,maxY,bestTime)
      if potentialAns != -1
        potentialAnswers << potentialAns
      end
    end

    return potentialAnswers.length > 0 ? potentialAnswers.min : -1

  end

  def solve(input)
    # at the start of the round move every blizzard
    # find the locations you can move to that are valid (including staying still), if there are none return -1
    # if you can move to the end location (y = maxY-1) then return the number of steps you took to get there
    # otherwise recursively run the function on all possible locations
    # easiest solution is to just do BFS at all locations at time T, keep track of all set of possible locations, once it contains the finish location youve found the optimal
    # is it possible with greedy DFS? probably, lets try it

    blizzardMaps = Hash.new { |h,k| h[k] = Hash.new([])} #in the form blizzardMaps[t][x,y] = [deltas] where delta = [xd, yd]
    startLocation, finishLocation, blizzardMap0 = parseInput(input)
    blizzardMaps[0] = blizzardMap0
    visited = Hash.new(false)
    bestTimeFound = Hash.new(100000000000000) # stupid hack for DFS

    part1 = DFS(startLocation, finishLocation, 0, blizzardMaps, visited, finishLocation[0]+1, finishLocation[1], bestTimeFound)
    puts part1

    part2A = DFS(finishLocation, startLocation, part1, blizzardMaps, Hash.new(false), finishLocation[0]+1, finishLocation[1], Hash.new(100000000000000))
    part2 = DFS(startLocation, finishLocation, part2A, blizzardMaps, Hash.new(false), finishLocation[0]+1, finishLocation[1], Hash.new(100000000000000))
    puts part2
  end
end


Day24.new