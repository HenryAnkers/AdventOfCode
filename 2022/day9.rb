require 'set'

class Day9
  def initialize
    input = File.read('input/day9.txt')

    puts solve(input, 1)
    puts solve(input, 9)
  end

  def sign(int)
    return 1 if int > 0
    return -1 if int < 0
    return 0
  end

  def moveTail (headLocation, tailLocation)
    # If the head is ever two steps directly up, down, left, or right from the tail, the tail must also move one step in that direction so it remains close enough:
    # Otherwise, if the head and tail aren't touching and aren't in the same row or column, the tail always moves one step diagonally to keep up:
    headDiff = headLocation[0] - tailLocation[0]
    tailDiff = headLocation[1] - tailLocation[1]

    if tailDiff.abs == 2 && headDiff == 0
      return [tailLocation[0], tailLocation[1] + sign(tailDiff)]
    end

    if headDiff.abs == 2 && tailDiff == 0
      return [tailLocation[0] + sign(headDiff), tailLocation[1]]
    end

    if (headDiff.abs + tailDiff.abs) > 2
      return [tailLocation[0] + sign(headDiff), tailLocation[1] + sign(tailDiff)]
    end

    return tailLocation
  end

  def solve(input, numTails)
    headLocation = [0,0]
    tailLocations = []
    (0...numTails).each do
      tailLocations << [0,0]
    end
    visitedLocations = Set.new()

    for instruction in input.split("\n")
      direction = instruction.split()[0]
      magnitude = instruction.split()[1].to_i  #I did input[1] by accident and spent 30 minutes trying to see why it worked for the example input and not the main one ;-;

      (0...magnitude).each do
        case direction
        when "U"
          headLocation[1] += 1
        when "D"
          headLocation[1] -= 1
        when "L"
          headLocation[0] -= 1
        when "R"
          headLocation[0] += 1
        end

        target = headLocation
        tailLocations = tailLocations.map do |location|
          newLocation = moveTail(target, location)
          target = newLocation
          newLocation
        end

        visitedLocations.add("#{target[0]}:#{target[1]}")
      end
    end

    visitedLocations.length
  end
end

Day9.new