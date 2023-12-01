class Day1
  def initialize
    input = File.read('input/day1.txt')

    puts part1(input)
    puts part2(input)
  end

  def getElfCalories(input)
    elvesFood = input.split("\n\n")

    elfCalories = []
    elvesFood.each do |elf|
      total = elf.split.map(&:to_i).sum
      elfCalories.push(total)
    end

    elfCalories.sort { |a, b| b <=> a } # highest to lowest
  end

  def part1(input)
    getElfCalories(input)[0]
  end

  def part2(input)
    getElfCalories(input).take(3).sum
  end
end

Day1.new
