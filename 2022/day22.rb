class Day22
  def initialize
    input = File.read('input/day22.txt').split("\n")

    # puts part1(input[0..-1])
    puts part2(input[0..-1])
  end


  def parseInput(input)
    directions = input.pop().split("")
    directionList = []
    current = ""
    directions.each do |token|
      if token == "L"
        directionList << current + "L"
        current = ""
      elsif token == "R"
        directionList << current + "R"
        current = ""
      else
        current += token
      end
    end
    if current != ""
      directionList << current
    end

    startingLocation = ""
    map = Hash.new{ |h, k| h[k] = Hash.new(-1) }
    input.each_with_index do |line, y|
      line.split("").each_with_index do |token, x|
        if startingLocation == "" && token != " "
          startingLocation = [x,y]
        end
        if token == "#"
          map[x][y] = 0
        elsif token == "."
          map[x][y] = 1
        end
      end
    end

    [map, directionList, startingLocation]
  end


  def wrapAroundX(map, startLocation, direction)
    maxX = map.keys.max
    x,y = startLocation
    if direction == "R"
      x = 0
    else
      x = maxX
    end

    delta = x == 0 ? 1 : -1

    while true
      if map[x][y] == 1
        return [x,y]
      elsif map[x][y] == 0
        return startLocation
      end

      x += delta
    end
  end


  def getNextPosition1(map, location, facing) #moves you forward one unit whatever direction you're facing.
    x,y = location
    if facing == 3      # move up
      if map[x][y-1] == 0
        return location
      elsif map[x][y-1] == 1
        return [x,y-1]
      else
        newY =  map[x].keys.max

        if map[x][newY] == 0
          return location
        end

        return [x, newY]
      end

    elsif facing == 1   # move down
      if map[x][y+1] == 0
        return location
      elsif map[x][y+1] == 1
        return [x,y+1]
      else
        newY =  map[x].keys.min

        if map[x][newY] == 0
          return location
        end

        return [x, newY]
      end

    elsif facing == 0   # move right
      if map[x+1][y] == 0
        return location
      elsif map[x+1][y] == 1
        return [x+1,y]
      else
        return wrapAroundX(map, location, "R")
      end

    else                # move left
      if map[x-1][y] == 0
        return location
      elsif map[x-1][y] == 1
        return [x-1,y]
      else
        return wrapAroundX(map, location, "L")
      end

    end
  end

  def move(map, direction, location, facing)
    directionToTurn = direction[-1]
    numberOfMoves = -1
    if directionToTurn != "L" and directionToTurn != "R"
      numberOfMoves = direction[0..-1].to_i
    else
      numberOfMoves = direction[0..-2].to_i
    end

    (1..numberOfMoves).each do |x|
      newLocation = getNextPosition1(map, location, facing)
      if location == newLocation
        break
      end
      location = newLocation
    end

    if directionToTurn == "R"
      facing = (facing + 1) % 4
    elsif directionToTurn == "L"
      facing = (facing - 1) % 4
    end

    return location, facing
  end


  def part1(input)
    facing = 0  #0 1 2 3 right down left up
    map, directions, location = parseInput(input)
    directions.each_with_index do |direction, i|
      location, facing = move(map, direction, location, facing)
      puts "#{location} #{facing} #{direction}"
      puts ""
    end

    return facing + ((location[0] + 1) * 4) + ((location[1] + 1) * 1000)
  end

  ###################################

  def moveCubeFace(map, startLocation, direction)
    x,y = startLocation
    cubeCoords = [y.div(50),x.div(50)]
    localX = x%50
    localY = y%50

    if cubeCoords === [0,1]
      if direction == 2 # left
        return 0, [0, 150 - localY]
      elsif direction == 3 #up
        return 0, [0, 150 + localX]
      end


    elsif cubeCoords == [0,2]
      if direction == 0 # right
        return 2, [99, 149 - localY]
      elsif direction == 1 #down
        return 2, [99, 50 + localX]
      elsif direction == 3 #up
        return 3, [localX, 199]
      end


    elsif cubeCoords == [1,1]
      if direction == 0 # right
        return 3, [100 + localY, 49]
      elsif direction == 2 #left
        return 1, [localY, 100]
      end


    elsif cubeCoords == [2,0]
      if direction == 3 # up
        return 0, [50, 50+localX]
      elsif direction == 2 #left
        return 0, [50, 49-localY]
      end

    elsif cubeCoords == [2,1]
      if direction == 0 # right
        return 2, [149, 49-localY]
      elsif direction == 1 # down
        return 2, [49, 150+localX]
      end


    elsif cubeCoords == [3,0]
      if direction == 2 # left
        return 1, [50 + localY, 0]
      elsif direction == 1 #down
        return 1, [100 + localX, 0]
      elsif direction == 0 #right
        return 3, [50 + localY, 149]
      end
      end

  end

  def getNextPosition2(map, location, facing) #moves you forward one unit whatever direction you're facing.
    x,y = location

    if facing == 3      # move up
      if map[x][y-1] == 0
        return facing, location
      elsif map[x][y-1] == 1
        return facing, [x,y-1]
      else
        newFacing, newLocation = moveCubeFace(map, location, facing)
        if map[newLocation[0]][newLocation[1]] != 1
          return facing, location
        else
          return newFacing, newLocation
        end
      end

    elsif facing == 1   # move down
      if map[x][y+1] == 0
        return facing, location
      elsif map[x][y+1] == 1
        return facing, [x,y+1]
      else
        newFacing, newLocation = moveCubeFace(map, location, facing)
        if map[newLocation[0]][newLocation[1]] != 1
          return facing, location
        else
          return newFacing, newLocation
        end
      end

    elsif facing == 0   # move right
      if map[x+1][y] == 0
        return facing, location
      elsif map[x+1][y] == 1
        return facing, [x+1,y]
      else
        newFacing, newLocation = moveCubeFace(map, location, facing)
        if map[newLocation[0]][newLocation[1]] != 1
          return facing, location
        else
          return newFacing, newLocation
        end
      end

    else                # move left
      if map[x-1][y] == 0
        return facing, location
      elsif map[x-1][y] == 1
        return facing, [x-1,y]
      else
        newFacing, newLocation = moveCubeFace(map, location, facing)
        if map[newLocation[0]][newLocation[1]] != 1
          return facing, location
        else
          return newFacing, newLocation
        end
      end
    end
  end


  def move2(map, direction, location, facing) #what a mess.. will clean up
    directionToTurn = direction[-1]
    numberOfMoves = -1
    if directionToTurn != "L" and directionToTurn != "R"
      numberOfMoves = direction[0..-1].to_i
    else
      numberOfMoves = direction[0..-2].to_i
    end

    (1..numberOfMoves).each do |x|
      facing, newLocation = getNextPosition2(map, location, facing)
      if location == newLocation
        break
      end
      location = newLocation
    end

    if directionToTurn == "R"
      facing = (facing + 1) % 4
    elsif directionToTurn == "L"
      facing = (facing - 1) % 4
    end

    return facing, location
  end


  def part2(input)
    # ?????????

    # my input looks like
    # #AB
    # #C#
    # DE#
    # F##
    # we can say A is (0,1), b is (0,2)...
    #   then our current cube is our location.div(50)
    #   our current position in the cube is location %50

    # mapping this out you can see the transitions (up from (0,1) goes to left of (3,0)...)
    # we know the coordinates of our new cube
    # we know the direction it will be facing after the transition and the edge we'll be on
    # bit annyoing...

    facing = 0  #0 1 2 3 right down left up
    map, directions, location = parseInput(input)
    directions.each_with_index do |direction, i|
      facing, location = move2(map, direction, location, facing)
      puts "#{location} #{facing} #{direction}"
      puts ""
    end

    return facing + ((location[0] + 1) * 4) + ((location[1] + 1) * 1000)
  end
end


Day22.new
