require 'set'

class Day15
  def initialize
    input = File.read('input/day15.txt').split("\n")

    startTime1 = Time.now
    puts "part 1: #{part1(input)} - calculated in #{Time.now - startTime1}s"

    startTime2 = Time.now
    puts "part 2: #{part2(input)} - calculated in #{Time.now - startTime2}s"
  end


  def part1(input)
    map = Hash.new{ |h, k| h[k] = Hash.new }

    input.each do |line|
      sensorX, sensorY = line.split(":")[0].split(",").map { |x| x.split("=")[1].to_i }
      beaconX, beaconY = line.split(":")[1].split(",").map { |x| x.split("=")[1].to_i }

      distanceX = (sensorX - beaconX).abs
      distanceY = (sensorY - beaconY).abs
      taxiDistance = distanceX + distanceY

      (0...taxiDistance + 1).each do |y|
          x = (taxiDistance - y)

          map[sensorY+y][sensorX+x] = map[sensorY+y][sensorX+x] || false
          map[sensorY+y][sensorX-x] = map[sensorY+y][sensorX-x] || false
          map[sensorY-y][sensorX+x] = map[sensorY-y][sensorX+x] || false
          map[sensorY-y][sensorX-x] = map[sensorY-y][sensorX-x] || false
      end
    end

    map[2000000].keys.max - map[2000000].keys.min
  end



  def part2(input) #any area not covered by the beacon has to be at least 1 block away from every beacon, which is basically the radius we calculated part 1 + 1
    #so can we find the diamond JUST containing those coordinates for each beacon...
    # add them to a set, remove them if we see them subsequently covered by another beacon (run the first beacon again to remove any leftovers...)
    # it will be 1 away from 4 beacons, and any that is 4 away from 1 beacon must be a beacon with radius 1 (otherwise it would overlap another one...)
    map = Hash.new(0)

    input.each do |line|
      sensorX, sensorY = line.split(":")[0].split(",").map { |x| x.split("=")[1].to_i }
      beaconX, beaconY = line.split(":")[1].split(",").map { |x| x.split("=")[1].to_i }

      distanceX = (sensorX - beaconX).abs
      distanceY = (sensorY - beaconY).abs
      taxiDistance = distanceX + distanceY + 1

      (0...taxiDistance + 1).each do |y|
        x = (taxiDistance - y)

        map["#{sensorY+y},#{sensorX+x}"] += 1
        map["#{sensorY+y},#{sensorX-x}"] += 1
        map["#{sensorY-y},#{sensorX+x}"] += 1
        map["#{sensorY-y},#{sensorX-x}"] += 1
      end
    end

    map.each do |k,v|
      if v == 4
        return (4000000 * k.split(",")[1].to_i) + k.split(",")[0].to_i # imagine if I got the x and y coords mixed up and had to spend 2 minutes rerunning this piece of crap :^)
      end
    end
  end




  ### the graveyard of bad ideas ###




  # let x,y be the coords of the hidden beacon
  # we know that for every sensor, S, the beacon it is paired with B is such that the distance between S and B is less than the distance between S and (x,y)
  # so (2,18) is closer to (-2,15) than (x,y)
  # (9,16) is closer to (10,16) than (x,y)...
  # can we use this information to solve the question at all

  def addRange(arr, newRange) #not as simple as intersection, if they're touching we want to merge them...
    newRanges = []
    merged = false

    arr.each do |range|
      if !merged && (arr.size + newRange.size >= ([range.end, newRange.end].max - [range.first, newRange.first].min)) #they either touch or overlap?
        newRanges << ([range.first, newRange.first].min..[range.end, newRange.end].max)
        merged = true
      else
        newRanges << range
      end
    end

    if not merged
      newRanges << newRange
    end

    newRanges
  end

  def part2b(input)  #first way doesn't really let me calculate this... need to be smarter?
    map = Hash.new{ |h, k| h[k] = [] }

    input.each do |line|
      sensorX, sensorY = line.split(":")[0].split(",").map { |x| x.split("=")[1].to_i }
      beaconX, beaconY = line.split(":")[1].split(",").map { |x| x.split("=")[1].to_i }


      distanceX = (sensorX - beaconX).abs
      distanceY = (sensorY - beaconY).abs
      taxiDistance = distanceX + distanceY


      (0...taxiDistance + 1).each do |y|
        x = (taxiDistance - y)

        map[sensorY+y] = addRange(map[sensorY+y], ([-x,x].min..[-x,x].max))
        map[sensorY+y] = addRange(map[sensorY-y], ([-x,x].min..[-x,x].max))
      end
    end

    puts (map[2000000]) # this isnt even close to being right
  end
end


Day15.new