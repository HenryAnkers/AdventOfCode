require 'json'

class Day13
  def initialize
    input = File.read('input/day13.txt')

    puts part1(input)
    puts "#{part2(input)} or #{part2b(input)}"
  end

  def comparePair(left, right)
    left.each_with_index do |l, index|
      if index >= right.length
        return -1 # right has run out of items first...
      end

      comparison = compareElement(left[index], right[index])
      return comparison if comparison != 0
    end

    return left.length == right.length ? 0 : 1
  end

  def compareElement(left, right)
    if left.is_a?(Numeric) && right.is_a?(Numeric)
      if left == right
        return 0
      end
      return left < right ? 1 : -1
    end

    if left.is_a?(Array) && right.is_a?(Array)
      return comparePair(left, right)
    end

    if left.is_a?(Array)
      return comparePair(left, [right])
    end

    return comparePair([left], right)
  end

  def part1(input)
    sumOfCorrect = 0

    pairs = input.split("\n\n").map{ |x| x.split("\n").map{ |y| JSON.parse(y) }}

    pairs.each_with_index do |pair, index|
      if comparePair(pair[0], pair[1]) == 1
        sumOfCorrect += (index + 1)
      end
    end

    sumOfCorrect
  end


  def part2(input)
    fullArray = []
    signal = 1
    input.split("\n\n").map{ |x| x.split("\n").map{ |y| JSON.parse(y) }}.each do |x| #is there no way to map to multiple values in Ruby without using .flatten?
      fullArray << x[0]
      fullArray << x[1]
    end

    decoder1 = [[2]]
    decoder2 = [[6]]
    fullArray << decoder1
    fullArray << decoder2

    sorted = fullArray.sort do |a,b|
      -compareElement(a,b)
    end

    sorted.each_with_index do |element, index|
      if element == decoder1 || element == decoder2
        signal *= (index + 1)
      end
    end

    signal
  end

  def part2b(input)
    decoder1 = [[2]]
    decoder2 = [[6]]
    lessThanD1 = 1 # technically number of elements < the decoder + 1
    lessThanD2 = 2 # 2 because we already know [[2]] < [[6]]

    pairs = input.split("\n\n").map{ |x| x.split("\n").map{ |y| JSON.parse(y) }}

    pairs.each_with_index do |pair, index|
      if comparePair(pair[0], decoder1) == 1
        lessThanD1 += 1
        lessThanD2 += 1
      elsif comparePair(pair[0], decoder2) == 1
        lessThanD2 += 1
      end

      if comparePair(pair[1], decoder1) == 1
        lessThanD1 += 1
        lessThanD2 += 1
      elsif comparePair(pair[1], decoder2) == 1
        lessThanD2 += 1
      end
    end

    lessThanD1 * lessThanD2
  end
end


Day13.new