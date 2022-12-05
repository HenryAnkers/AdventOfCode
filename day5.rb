class Day5
  def initialize
    input = File.read('input/day5.txt')
    input = File.read('input/day5.txt')

    puts part1(input)
    puts part2(input)
  end

  def getStacks(rows)
    stacks = [[],[],[],[],[],[],[],[],[]] #assume 9 stacks in input... I might as well just hardcode this

    rows.each do |row|
      (0...9).each { |index|
        nextCharacter = (index * 4) + 1 #The first number will be on character 1, the second character 4...
        if row[nextCharacter] != " " 
          stacks[index] << row[nextCharacter]
        end
      }
    end

  stacks #yeah I should have hardcoded this
  end

  def parseInstruction(instruction, stacks)
    instructionTokens = instruction.split()
    amountToMove = instructionTokens[1].to_i
    stackToMoveFromIndex = instructionTokens[3].to_i - 1
    stackToMoveToIndex = instructionTokens[5].to_i - 1

    return amountToMove, stacks[stackToMoveFromIndex], stacks[stackToMoveToIndex]
  end

  def part1(input)
    inputRows = input.split("\n")
    stackInput = inputRows.slice(0,8)
    stacks = getStacks(stackInput)

    instructionInput = inputRows.drop(10)

    instructionInput.each do |instruction|
      amountToMove, stackToMoveFrom, stackToMoveTo = parseInstruction(instruction, stacks)

      for i in 0...amountToMove
        stackToMoveTo.unshift(stackToMoveFrom.shift)
      end
    end

    answer = stacks.map {|x| x[0]}
    answer.join
  end


  def part2(input)
    inputRows = input.split("\n")
    stackInput = inputRows.slice(0,8)
    stacks = getStacks(stackInput)

    instructionInput = inputRows.drop(10)

    instructionInput.each do |instruction|
      amountToMove, stackToMoveFrom, stackToMoveTo = parseInstruction(instruction, stacks)
      
      stackToMove = []   #shift elements from the top, then add them in reverse order by popping in subsequent loop
      for i in 0...amountToMove
        stackToMove << stackToMoveFrom.shift
      end

      for i in 0 ... amountToMove
        stackToMoveTo.unshift(stackToMove.pop)
      end
    end

    answer = stacks.map {|x| x[0]}
    answer.join
  end
end


Day5.new