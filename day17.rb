class Day17
  def initialize
    input = File.read('input/day17.txt').split()[0]

    puts part1(input)
    puts part2(input)
  end

  def visualize(map, maxY, shape)  # finally something simple enough for me to visualize
    (0..maxY+2).each do |y|
      y = (maxY + 2 - y)
      string = ""
      (1..8).each do |x|
        if map[x][y] == true || shape.include?([x,y])
          string << '#'
        else
          string << ' '
        end
      end
      puts string
    end
  end

  # left edge needs to be two units away from the left wall (x=1), bottom edge needs to be at y
  def getNextShape(number, y)
    if number == 0
      return [[3,y],[4,y],[5,y],[6,y]]
    end

    if number == 1
      return [[4,y],[3,y+1],[4,y+1],[5,y+1],[4,y+2]]
    end

    if number == 2
      return [[3,y],[4,y],[5,y],[5,y+1],[5,y+2]]
    end

    if number == 3
      return [[3,y],[3,y+1],[3,y+2],[3,y+3]]
    end

    if number == 4
      return [[3,y],[4,y], [3,y+1], [4,y+1]]
    end
  end

  def updateMap(map, shape, theWind, windIndex)
    while true
      # attempt to move the shape according to the wind
      #   - check every brick. if any brick would occupy an area with rock (map[x+i][y] is true) then break
      #   - if we do not break out, we will update all the x coordinates according to the wind direction
      d = theWind[windIndex] == "<" ? -1 : 1

      collision = false
      shape.each do |coord|
        x, y = coord
        if x+d > 7 || x+d < 1 || map[x+d][y]
          collision = true
        end
      end

      if !collision
        shape.each do |coord|
          coord[0] += d
        end
      end

      # now attempt to move the shape down 1
      #   - check every brick. if any brick would occupy an area with rock (map[x][y-1] is true)
      #     -  add each current coordinate of the shape to the map
      #     -  return the maximum Y value from the new shape
      #   - if it's not true, translate the shape down by 1 and continue
      landed = false
      shape.each do |coord|
        x, y = coord
        if map[x][y-1]
          landed = true
        end
      end

      if landed
        yValues = []
        shape.each do |coord|
          x, y = coord
          yValues << y
          map[x][y] = true
        end
        return windIndex+1, yValues.max
      else
        shape.each do |coord|
          coord[1] -= 1
        end
      end

      windIndex += 1
      windIndex = windIndex % theWind.length
    end
  end

  def part1(input) # normal one today... jfc
    map = Hash.new{ |h, k| h[k] = Hash.new(false) }
    (1..7).each do |x|
      map[x][0] = true   # model so the y=0 is the floor, cos if we reach y=1 we have 1 unit of rock, y=2 we have 2 units of rock.. etc, so maxY becomes our answer
    end

    maxY = 0
    windIndex = 0

    (0..2021).each do |i|
      nextShape = getNextShape(i%5, maxY+4)
      windIndex, potentialMaxY = updateMap(map, nextShape, input, windIndex)
      maxY = [potentialMaxY, maxY].max
    end

    maxY
  end


  def getMemoKey(map, y, windIndex, shapeIndex)
    distanceStrings = []
    key =

    (1..7).each do |x|
      yDistance = 1

      while true
        if map[x][y-yDistance] == true
          distanceStrings << "#{x},#{yDistance}"
          break
        end

        yDistance += 1
      end
    end

    [distanceStrings.join("-"), windIndex, shapeIndex].join(":")
  end

  def part2(input)
    # again IDK how to solve this
    # I guess it's too high for us to simulate but the shapes repeat, the wind direction repeats, so there must be a loop, how do we find it?
    # Maybe it's like the top y = 100 repeats? maybe its the top 1000 lines? maybe its the top 10000 lines?

    # Consider a shape being droppeds Y coordinates to be 0 (from the left edge or whatever how we create them, so some are technically -1)
    # if for every X value 1 to 7, the distance between y = 0 and the first brick, and the shape is the same, and the wind direction is the same, then the program must repeat
    # so memo over [FancyYCoordinatesMadeIntoAKey,ShapeNumber,WindDirectionIndex] might work
    # WindDirectionIndex is 10k, ShapeNumber is 5, IDK how many values are possible for the first one, this might mess it up or might not (its n^7, so its risky, im hooing its all like <30 in practice?) [In reality it was always <15 or so]
    # then if we've seen it before, we take the height then, look at the height now, thats the difference...
    #   calculate the height difference between the current cycle and the cycle that would get us nearest to 1000000000000, call this D. Suppose this number leaves us C cycles away from 1000000000000...
    #   run the current cycle C more times, get the maxY height, return maxY + D, done

    # this worked for the example but not the actual input
    # realised that the second time we saw a cycle, the height delta was slightly different than the first time we found the cycle, this gave a correct answer
    # I think this is because the first cycle hit the floor, so the height delta between 1 and 2 would be slightly different to 2 and 3, 3 and 4.. etc

    map = Hash.new{ |h, k| h[k] = Hash.new(false) }
    (1..7).each do |x|
      map[x][0] = true
    end

    memo = Hash.new()
    maxY = 0
    windIndex = 0
    totalHeightDelta = -1
    remainingRocks = -1 # dodgy hack: we initialize this value so we can break out once we find a cycle.

    (0..100000).each do |i|
      nextShape = getNextShape(i%5, maxY+4)
      memoKey = getMemoKey(map, maxY+4, windIndex, i%5)

      if memo[memoKey] && totalHeightDelta == -1 # once we have found the height delta we will add the offset, so we must only run it once
        if memo[memoKey][2] == 0
          memo[memoKey] = [i, maxY, 1]  # see above for why we can't use the first cycle values
        else
          prevIndex, prevMaxY = memo[memoKey]
          indexDelta = i - prevIndex
          heightDelta = maxY - prevMaxY

          numCyclesToRun = (1000000000000 - i).div(indexDelta) # what is the maximum number of cycles (which are size indexDelta) we can fit into the remaining number of cycles (~1000000000000)
          totalHeightDelta = numCyclesToRun * heightDelta
          remainingRocks = 1000000000000 - (i + numCyclesToRun * indexDelta) # we skipped the cycles, we now need to run the simulation this many more times and then add the totalHeightDelta to whatever our maxY is at the end
        end
      else
        memo[memoKey] = [i, maxY, 0] # this is the first time we've found the cycle, make note and continue
      end


      windIndex, potentialMaxY = updateMap(map, nextShape, input, windIndex)
      maxY = [potentialMaxY, maxY].max

      if remainingRocks != -1
        remainingRocks -= 1
        if remainingRocks == 0
          return totalHeightDelta + maxY
        end
      end
      end
    maxY
  end
end
Day17.new