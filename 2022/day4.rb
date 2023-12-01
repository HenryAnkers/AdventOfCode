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

  def gptSolution(input)
    # Parse the input string into an array of ranges
    ranges = input.split("\n").map do |line|
      # Split each line on the comma and parse the numbers
      a, b = line.split(",").map do |range|
        # Split each range on the dash and parse the numbers
        range.split("-").map(&:to_i)
      end
      # Create a range object for each pair of numbers
      a_range = a[0]..a[1]
      b_range = b[0]..b[1]
      [a_range, b_range]
    end

    # Count how many pairs of ranges have one range completely containing the other
    count_complete_containment = ranges.count do |a_range, b_range|
      a_range.cover?(b_range) || b_range.cover?(a_range)
    end

    # Count how many pairs of ranges have a non-empty intersection
    count_non_empty_intersection = ranges.count do |a_range, b_range|
      !(a_range.to_a & b_range.to_a).empty?
    end

    puts "Number of pairs with complete containment: #{count_complete_containment}"
    puts "Number of pairs with non-empty intersection: #{count_non_empty_intersection}"
  end
end


Day4.new