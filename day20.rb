class Day20
  def initialize
    input = File.read('input/day20.txt').split("\n").map {|x| x.to_i}

    puts part1(input)
    puts part2(input)
  end

  def mixArray(input, uniqueArray)
    input.each_with_index do |x, i|
      prevIndex = uniqueArray.find_index([x,i])
      uniqueArray.delete_at(prevIndex)
      newIndex = (prevIndex + x) % (input.length - 1)
      if newIndex == 0
        newIndex = -1 # dumb hack, cba fixing
      end

      uniqueArray.insert(newIndex, [x,i])
    end
  end

  def part1(input)
    uniqueArray = []

    input.each_with_index do |x, i|  # dont know if I need to do this tbh
      uniqueArray << [x,i]
    end

    mixArray(input, uniqueArray)

    newArray = uniqueArray.map{|x| x[0]}
    zeroIndex = newArray.find_index(0)
    return newArray[(zeroIndex+1000) % newArray.length] + newArray[(zeroIndex+2000) % newArray.length] + newArray[(zeroIndex+3000) % newArray.length]
  end

  def part2(input)
    uniqueArray = []
    input.each_with_index do |x, i|
      input[i] *= 811589153
      uniqueArray << [input[i],i]
    end

    (0..9).each do
      mixArray(input, uniqueArray)
    end

    newArray = uniqueArray.map{|x| x[0]}
    zeroIndex = newArray.find_index(0)
    return newArray[(zeroIndex+1000) % newArray.length] + newArray[(zeroIndex+2000) % newArray.length] + newArray[(zeroIndex+3000) % newArray.length]
  end
end

Day20.new