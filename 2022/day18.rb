require 'set'

class Day18
  def initialize
    input = File.read('input/day18.txt').split("\n").map{|x| x.split(",").map{|y| y.to_i}}

    puts part1(input)
    puts part2(input)
  end


  def calculateUncoveredFaces(input, map)
    uncoveredFaces = 0
    input.each do |coords|
      x,y,z = coords
      faces = 6

      faces -= map["#{x+1},#{y},#{z}"]
      faces -= map["#{x-1},#{y},#{z}"]
      faces -= map["#{x},#{y+1},#{z}"]
      faces -= map["#{x},#{y-1},#{z}"]
      faces -= map["#{x},#{y},#{z+1}"]
      faces -= map["#{x},#{y},#{z-1}"]

      uncoveredFaces += faces
    end

    uncoveredFaces
  end

  def part1(input)
    # we can add each value to a map representing the 3d coords
    # for each value we can then check how many neighbours it has. this is the number of uncovered faces for the drop

    map = Hash.new(0)
    input.each do |coords|
      x,y,z = coords
      map["#{x},#{y},#{z}"] = 1
    end

    calculateUncoveredFaces(input, map)
  end


  def part2(input)
    # suppose we make a map X Y Z that contains all the points possible for a slice of 3d space (the biggest coord see is 16 or so, this is about 32^3 points max)
    # we initially mark all the values as 1
    # we then mark all the location of the rocks as 0, leaving us with a map of just air nodes
    # we then do a BFS from a air node we know isn't in a rock ((maxX+1,maxY+1,maxZ+1) is safe as no lava drops go that far), these are open air nodes and we mark these with a value of 0 in our new map as well
    # therefore all open air nodes and all lava has been removed from our map... leaving only the closed air nodes
    # we can then find the surface area of these using the same method we used for part 1
    # then our answer is (part1 - surface area of trapped air)

    lavaMap = Hash.new(0)
    trappedAirNodes = Hash.new(1)

    maxX = 0 # should have just hardcoded this in as 20 or something
    maxY = 0
    maxZ = 0
    input.each do |coords|
      x,y,z = coords

      maxX = [maxX, x.abs].max
      maxY = [maxY, y.abs].max
      maxZ = [maxZ, z.abs].max

      lavaMap["#{x},#{y},#{z}"] = 1
    end

    totalCount = 0
    (-maxX-1..maxX+1).each do |x|
      (-maxY-1..maxY+1).each do |y|
        (-maxZ-1..maxZ+1).each do |z|
          totalCount +=1
          trappedAirNodes["#{x},#{y},#{z}"] = 1  # why am I still doing this as a hash if I've got to fill out the points anyway? because I'm an idiot
        end
      end
    end
    input.each do |coords|
      x,y,z = coords
      trappedAirNodes["#{x},#{y},#{z}"] = 0
    end


    openAirNodes = 1
    nodesToExplore = [[maxX+1,maxY+1,maxZ+1],[-(maxX+1),-(maxY+1),-(maxZ+1)]]
    explored = Hash.new(0)
    explored["#{maxX+1},#{maxY+1},#{maxZ+1}"] = 1
    trappedAirNodes["#{maxX+1},#{maxY+1},#{maxZ+1}"] = 0

    while nodesToExplore.length > 0
      nextNode = nodesToExplore.shift()
      x,y,z = nextNode

      [[x+1,y,z],[x-1,y,z],[x,y+1,z],[x,y-1,z],[x,y,z+1],[x,y,z-1]].each do |coord|
        x_i,y_i,z_i = coord

        if (x_i <= maxX + 1 && x_i >= -(maxX+1)) && (y_i <= maxY + 1 && y_i >= -(maxY+1)) && (z_i <= maxZ + 1 && z_i >= - (maxZ+1)) && lavaMap["#{x_i},#{y_i},#{z_i}"] == 0 && explored["#{x_i},#{y_i},#{z_i}"] == 0 # it's in the grid and not in the rock, so it's an open air node
          openAirNodes += 1
          explored["#{x_i},#{y_i},#{z_i}"] = 1
          trappedAirNodes["#{x_i},#{y_i},#{z_i}"] = 0
          nodesToExplore << [x_i,y_i,z_i]
        end
      end
    end

    part1(input) - calculateUncoveredFaces(trappedAirNodes.keys.select{|x| trappedAirNodes[x] == 1}.map{|x| x.split(",").map{|x| x.to_i}}, trappedAirNodes)
  end
end


Day18.new