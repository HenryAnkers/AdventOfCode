class Day10
  def initialize
    input = File.read('input/day10.txt')

    puts part1(input)
    puts part2(input)
  end

  def runInstruction(instructionString, cycle, register)
    instruction = instructionString.split()

    if instruction[0] == "addx"
      return cycle+2, register + instruction[1].to_i
    end

    return cycle+1, register  #noop
  end


  def part1(input)
    cycle = 0
    register = 1
    signalStrengths = {}

    for instruction in input.split("\n")
      prevRegister = register
      cycle, register = runInstruction(instruction, cycle, register)

      if (cycle + 20) % 40 == 1 && !signalStrengths.has_key?(cycle-1) #we passed the 20 + 40N'th cycle during run instruction and didn't fill it in.
        signalStrength = (cycle - 1) * prevRegister
        signalStrengths[cycle-1] = signalStrength
      end

      if (cycle + 20) % 40 == 0
        signalStrength = cycle * prevRegister
        signalStrengths[cycle] = signalStrength
      end
    end

    signalStrengths.values.reduce(:+)
  end

  def part2(input) #if the sprite is on the pixel being drawn then its a '#', otherwise it's ' ' .
    cycle = 0
    register = 1
    pixels = []

    for instruction in input.split("\n")
      prevRegister = register
      prevCycle = cycle

      cycle, register = runInstruction(instruction, cycle, register)

      (prevCycle...cycle).each do |c|
        if (c % 40 == prevRegister - 1) || (c % 40 == prevRegister) || (c % 40 == prevRegister + 1)
          pixels << "#"
        else
          pixels << " "
        end
      end
    end

    pixels.each_slice(40).map do |x|
      x.join()
    end
  end
end


Day10.new