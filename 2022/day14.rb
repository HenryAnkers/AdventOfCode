class Day14
  def initialize
    input = File.read('input/day14.txt')

    puts part1(input)
    puts part2(input)
  end


  def parseInputMap(walls)
    map = Hash.new{ |h, k| h[k] = Hash.new }
    lowestWall = -1

    walls.each do |wall|
      coords = wall.split(" -> ")

      prevX = coords[0].split(",")[0].to_i
      prevY = coords[0].split(",")[1].to_i

      (1...coords.length).each do |i|
        newX = coords[i].split(",")[0].to_i
        newY = coords[i].split(",")[1].to_i

        minX = [prevX, newX].min
        maxX = [prevX, newX].max
        minY = [prevY, newY].min
        maxY = [prevY, newY].max

        (minX...maxX + 1).each do |x|
          (minY...maxY + 1).each do |y|
            map[x][y] = true
            lowestWall = [lowestWall, y].max
          end
        end

        prevX = newX
        prevY = newY
      end
    end

    [map, lowestWall]
  end

  def processSand(sand, map, maxY)
    x = sand[0]
    y = sand[1]

    if y > maxY  # falling into the abyss :pensive:
      return -1
    end

    if !map[x][y+1]  # we can go down, can't settle
      return processSand([x, y+1], map, maxY)
    end

    if !map[x-1][y+1]  # we can go down and left, can't settle
      return processSand([x-1, y+1], map, maxY)
    end

    if !map[x+1][y+1]  # we can go down and right, can't settle
      return processSand([x+1, y+1], map, maxY)
    end

    map[x][y] = true  # we have settled
    return 1
  end

  def part1(input)
    walls = input.split("\n")
    map, lowestWall = parseInputMap(walls)
    settled = 0

    while true
      processNext = processSand([500,0], map, lowestWall)

      if processNext == -1
        return settled
      end

      settled += 1
    end
  end

  def part2(input)
    walls = input.split("\n")
    map, lowestWall = parseInputMap(walls)
    settled = 0
    floorDepth = lowestWall + 2

    (map.keys.min - 10000...map.keys.max + 10000).each do |x| # bit surprised this worked lol
      map[x][floorDepth] = true
    end

    while true
       processSand([500,0], map, floorDepth)

      if map[500][0] == true
        return settled + 1
      end

      settled += 1
    end
  end
end


Day14.new