class Day7
  def initialize
    input = File.read('input/day7.txt')

    puts part1(input)
    puts part2(input)
  end


  def transverseFilesystem(input)
    instructions = input.split("\n")
    directoryValues = Hash.new(0)
    currentDirectories = []

    instructions.each do |i|
      tokens = i.split()

      if tokens[0] == "$"
        if tokens[1] == "cd"
          directoryName = tokens[2]

          if directoryName == ".."
            currentDirectories.pop
          else            #Input doesn't do 'cd /' apart from first line so don't need to code it
            currentDirectories << directoryName
          end
        end
      else
        filesize = tokens[0].to_i

        currentDirectory = ""
        currentDirectories.each do |directory|
          if currentDirectory != ""
            currentDirectory += "/"
          end
          currentDirectory += directory
          directoryValues[currentDirectory] += filesize
        end
      end
    end

    directoryValues
  end


  def part1(input)
    directoryValues = transverseFilesystem(input)

    totalValue = 0
    directoryValues.values.each do |value|
      if value <= 100000
        totalValue += value
      end
    end

    totalValue
  end


  def part2(input)
    directoryValues = transverseFilesystem(input)
    totalSpace = 70000000
    targetFreeSpace = 30000000
    totalSpaceUsed = directoryValues['/']
    amountToRemove = targetFreeSpace - (totalSpace - totalSpaceUsed)


    directoryToRemoveSize  = totalSpaceUsed   #worst case = just remove everything
    directoryValues.values.each do |value|
      if value > amountToRemove
        directoryToRemoveSize = [value, directoryToRemoveSize].min
      end
    end
  end
end


Day7.new