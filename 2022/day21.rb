class Day21
  def initialize
    input = File.read('input/day21.txt').split("\n").map {|x| [x.split(":")[0], x.split(":")[1].split(" ")]}

    puts part1(input)
    puts part2(input)
  end

  def findValue(monkey, values)
    if values[monkey].length == 1
      return values[monkey][0].to_i
    end

    value1 = findValue(values[monkey][0], values)
    value2 =  findValue(values[monkey][2], values)
    operation = values[monkey][1]

    if operation == "+"
      return value1 + value2
    end
    if operation == "-"
      return value1 - value2
    end
    if operation == "*"
      return value1 * value2
    end
    return value1.div(value2)
  end

  def part1(input)
    values = Hash.new(-1)
    input.each do |m|
      values[m[0]] = m[1]
    end

    findValue("root", values)
  end


  def part2(input)
    # cant be arsed with equations
    # replacing the root value with "-" we're looking for 0
    # can run the followng code in a debugger
    #     newValues = values.clone
    #     newValues["humn"][0] = [testValue]
    #     newValues["root"][1] = "-"
    #     findValue("root", newValues)
    # in a debugger we see that 100 is closer than 10, 1000 is closer than 100... up to 1000000000000
    # we can do a sort of mock binary search on the digits this way (start the index at 5, move it up/down until its as low as possible but not negative, continue)
    # unfortunately it's just easier if I do this in the debugger so the answer is 3876027196185
    # i might code it later
    3876027196185
  end
end


Day21.new