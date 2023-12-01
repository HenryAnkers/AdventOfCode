class Day12
  def initialize
    input = File.read('input/day12.txt').split("\n")

    puts part1(input)
    puts part2(input)
  end

  def BFS(input, distance, nodesToVisit)
    while nodesToVisit.length > 0
      nextLocation = nodesToVisit.shift()
      x = nextLocation[0]

      y = nextLocation[1]

      [[x+1,y],[x-1,y],[x,y+1],[x,y-1]].each do |coord|
        x_i = coord[0]
        y_i = coord[1]

        if (x_i >= 0 && y_i >= 0 && x_i < input.length && y_i < input[x].length)
          if input[x_i][y_i] == "E" && "z".ord - input[x][y].ord <= 1
            return distance[x][y] + 1
          end

          if distance[x_i][y_i] == -1 && input[x_i][y_i].ord - input[x][y].ord <= 1
            distance[x_i][y_i] = distance[x][y] + 1
            nodesToVisit << [x_i,y_i]
          end
        end
      end
    end
  end


  def part1(input)
    distance  = input.map { |row| Array.new(row.length, -1) }
    nodesToVisit = []

    startingLocation = [0,0]
    input.each_with_index do |d, x|  #5 minutes to code BFS, 1 hour to debug it to find out the input doesn't always start at [0,0]..........
      d.split("").each_with_index do |c, y|
        if c == "S"
          nodesToVisit << [x,y]
          input[x][y] = "a"
          distance[x][y] = 0
        end
      end
    end

    BFS(input, distance, nodesToVisit)
  end


  def part2(input)
    distance  = input.map { |row| Array.new(row.length, -1) }
    nodesToVisit = []

    startingLocation = [0,0]
    input.each_with_index do |d, x|
      d.split("").each_with_index do |c, y|
        if c == "a" || c == "S"
          nodesToVisit << [x,y]
          input[x][y] = "a"
          distance[x][y] = 0
        end
      end
    end

    BFS(input, distance, nodesToVisit)
  end
end


Day12.new