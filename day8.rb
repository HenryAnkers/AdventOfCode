class Day8
  def initialize
    input = File.read('input/day8.txt')

    solve(input)
  end

  def processLeftArray(inputArray, visibleArray, scoreArray)
    inputArray.each_with_index do |row, row_index|
      checkLeftVisible(row, row_index, visibleArray)
      calculateLeftViewScore(row, row_index, scoreArray)
      end
    end

  def checkLeftVisible(row, row_index, visibleArray)
    tallestFound = -1
    row.each_with_index do |tree, col_index|
      if tree.to_i > tallestFound
        visibleArray[row_index][col_index] = 1
        tallestFound = tree.to_i
      end
    end
  end

  def calculateLeftViewScore(row, row_index, scoreArray)
    stack = []   #stack of elements which we haven't found a greater or equal element on the right hand side for, these are obviously in descending order...

    row.each_with_index do |tree, col_index|
      while !stack.empty? && row[stack.last].to_i <= tree.to_i  # we have found a greater element for the smallest item in the stack. iterate over until we're back to square one
        nextSmallerIndex = stack.pop
        scoreArray[row_index][nextSmallerIndex] *= (col_index - nextSmallerIndex)
      end

      stack << col_index # add the element to the stack cos we clearly havent found an element on the RHS (as we havent checked any yet..)
    end

    while !stack.empty? #remaining ones are good until the edge so calculate this distance
      nextSmallerIndex = stack.pop
      scoreArray[row_index][nextSmallerIndex] *= (row.length - 1 - nextSmallerIndex)
    end

    scoreArray[row_index][0] = 0 # finally we just set all outer elements to 0 cos we dont care about them
    scoreArray[row_index][-1] = 0
  end
  
  def solve(input)
    inputArray = input.split("\n").map { |x| x.split("") }
    visibleArray = inputArray.map { |row| Array.new(row.size, 0) }
    viewScoreArray = inputArray.map { |row| Array.new(row.size, 1) }

    processLeftArray(inputArray, visibleArray, viewScoreArray)

    inputArray = inputArray.map(&:reverse)
    visibleArray = visibleArray.map(&:reverse)
    viewScoreArray = viewScoreArray.map(&:reverse)
    processLeftArray(inputArray, visibleArray, viewScoreArray) #checking right to left

    inputArray = inputArray.transpose
    visibleArray = visibleArray.transpose
    viewScoreArray = viewScoreArray.transpose
    processLeftArray(inputArray, visibleArray, viewScoreArray) #checking top to bottom

    inputArray = inputArray.map(&:reverse)
    visibleArray = visibleArray.map(&:reverse)
    viewScoreArray = viewScoreArray.map(&:reverse)
    processLeftArray(inputArray, visibleArray, viewScoreArray) #checking bottom to top

    puts visibleArray.flatten.reduce(:+)
    puts viewScoreArray.flatten.max
  end
end


Day8.new