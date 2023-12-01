class Day11
  def initialize
    input = File.read('input/day11.txt')

    puts solve(1,input)
    puts solve(2,input)
  end


  def parseInput(input)
    monkeys = input.split("\n\n")
    parsedMonkeys = []

    monkeys.each_with_index do |monkey|
      monkeyLines = monkey.split("\n")
      items = monkeyLines[1].split(":")[1].split(",").map{ |x| x.to_i }
      operation = monkeyLines[2].split(":")[1].split("=")[1]
      divisibleBy = monkeyLines[3].split()[-1].to_i
      ifDivisible = monkeyLines[4].split()[-1].to_i
      ifNotDivisible = monkeyLines[5].split()[-1].to_i

      parsedMonkeys << {items: items, operation: operation, divisibleBy: divisibleBy, ifDivisible: ifDivisible, ifNotDivisible: ifNotDivisible}
    end

    parsedMonkeys
  end

  def solve(part,input)
    monkeys = parseInput(input)
    lcm = monkeys.map{ |m| m[:divisibleBy]}.reduce(:*)
    numInspections = Hash.new(0)
    numOps = part == 1 ? 20 : 10000

    (0...numOps).each do
      for m_index in (0...monkeys.length)
        monkey = monkeys[m_index]
        items = monkey[:items]
        operation = monkey[:operation]
        operationTokens = operation.split()

        items.each do |i|
          numInspections[m_index] += 1

          operationTokens = operation.split()
          if operationTokens[2] == "old"
            operationTokens[2] = i.to_s
          end

          newValue = -1
          if operationTokens[1] == "+"
            newValue = (i + operationTokens[2].to_i).div(part == 1 ? 3 : 1) % lcm
          else
            newValue = (i * operationTokens[2].to_i).div(part == 1 ? 3 : 1) % lcm
          end

          if newValue % monkey[:divisibleBy] == 0
            monkeys[monkey[:ifDivisible]][:items] << newValue
          else
            monkeys[monkey[:ifNotDivisible]][:items] << newValue
          end
        end

        monkey[:items] = []
      end

    end

    numInspections.values.sort[-1] *numInspections.values.sort[-2]
  end
end


Day11.new