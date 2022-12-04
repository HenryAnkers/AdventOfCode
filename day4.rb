class Day4
  def initialize
    input = File.read('input/day4.txt')

    puts part1(input)
    puts part2(input)
  end

  def getElfAssignments(input)
    assignmentPairs = input.split("\n")

    assignmentPairs.map { |pair| pair.split(",").map { |elf| elf.split("-").map(&:to_i)} }
  end

  def part1(input)
    total = 0
    assignments = getElfAssignments(input)
    assignments.each do |assignment|
      elf1 = assignment[0]
      elf2 = assignment[1]

      if (elf1[0] <= elf2[0] && elf1[1]  >= elf2[1]) || ( elf1[0] >= elf2[0] && elf1[1] <= elf2[1]) #one assignment has a smaller min and larger max than the other
        total += 1
      end
    end

    return total
  end


  def part2(input)
    total = 0
    assignments = getElfAssignments(input)
    assignments.each do |assignment|
      elf1 = assignment[0]
      elf2 = assignment[1]

      unless (elf1[0] < elf2[0] && elf1[1] < elf2[0]) || (elf2[0] < elf1[0] && elf2[1] < elf1[0]) #one assignment starts and ends before the other assignment
        total += 1
      end
    end

    return total
  end
end


Day4.new