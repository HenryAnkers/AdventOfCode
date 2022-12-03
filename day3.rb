class Day3
  def initialize
    input = File.read('input/day3.txt')

    puts part1(input)
    puts part2(input)
  end

  def getPriority(char)
    priority = 0
    if char == char.upcase
      priority += 26
    end

    return priority + (char.upcase.ord - "A".upcase.ord) + 1
  end

  def part1(input)
    backpacks = input.split("\n")

    total = 0

    backpacks.each do |backpack|
      compartment1 = backpack[0, backpack.length/2]
      compartment2 = backpack[backpack.length/2, backpack.length]

      shared = compartment1.chars & compartment2.chars
      shared.each do |item|
        total += getPriority(item)
      end
    end

    total
  end


  def part2(input)
    backpacks = input.split("\n")

    total = 0

    while backpacks.length > 0
      backpack1 = backpacks.pop()
      backpack2 = backpacks.pop()
      backpack3 = backpacks.pop()

      badge = (backpack1.chars & backpack2.chars & backpack3.chars)[0]

      total += getPriority(badge)
    end

    total
  end
end

Day3.new