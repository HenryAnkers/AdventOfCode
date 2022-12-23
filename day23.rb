class Day23
  def initialize
    input = File.read('input/day23.txt')

    puts solve(input.split("\n"))
  end

  def parseInput(input)
    map = Hash.new(false)
    input.each_with_index do |line, y|
      line.split("").each_with_index do |token, x|
        if token == "#"
          map[[x,y]] = true
        end
      end
    end

    map
  end

  def solve(input)
    # first half - if no elves in the first 8 tiles around me, then continue
    # if none in y-1, then propose north one step
    # if none in y+1, then propose south
    # if none in x-1, then propose west
    # if none in x+1, then propose east
    # second half - move to the tile if and only if nobody else proposed the same tile, or stand still
    # move the first direction (north) to the end of the list (so now south, west, east, north)

    elfMap = parseInput(input)
    moves = [[0,-1], [0,1], [-1,0], [1,0]]
    checks = [
      [[0,-1], [1,-1],[-1,-1]],
      [[0,1], [1,1], [-1,1]],
      [[-1,0],[-1,-1],[-1,1]],
      [[1,0],[1,-1],[1,1]]
    ]
    currentMove = 0

    (0..1000000).each do |round|

      if round == 10
        minX, maxX, minY, maxY = 0,0,0,0
        elfMap.keys.each do |k|
          x,y = k
          minX = [x,minX].min
          maxX = [x,maxX].max
          minY = [y,minY].min
          maxY = [y,maxY].max
        end

        puts "#{(((maxX - minX) + 1) * ((maxY - minY) + 1)) - elfMap.keys.count}"
      end
      # [x,y] = elf x,y has proposed a move
      # -1 = no move proposed
      # 0 = two elves tried to propose, can't do anything
      proposedMoves = Hash.new(-1)
      elfMap.keys.each do |elf|
        x,y = elf
        hasElfNearby = false

        checks.each do |direction|
          direction.each do |c|
            xd, yd = c
            if elfMap[[x + xd, y + yd]]
              hasElfNearby = true
              break
            end
          end
        end

        if hasElfNearby
          tempMove = currentMove
          (0..3).each do |i|
            tempMove = (currentMove + i) % 4

            hasElfInProposed = false

            checks[tempMove].each do |c|
              xd, yd = c
              if elfMap[[x + xd, y + yd]]
                hasElfInProposed = true
                break
              end
            end

            if !hasElfInProposed
              xp, yp = moves[tempMove]
              proposed = [x + xp, y + yp]
              if proposedMoves[proposed] == -1  # only scenario we can do it
                proposedMoves[proposed] = elf
              else                              # another elf already tried
                proposedMoves[proposed] = 0
              end

              break
            end
          end
        end
      end

      if proposedMoves.keys.length == 0
        return round + 1
      end

      proposedMoves.keys.each do |k|
        if proposedMoves[k] != 0
          olfElfPosition = proposedMoves[k]
          newElfPosition = k

          elfMap[newElfPosition] = true
          elfMap.delete(olfElfPosition)
        end
      end

      currentMove = (currentMove + 1) % 4
      minX, maxX, minY, maxY = 0,0,0,0
      elfMap.keys.each do |k|
        x,y = k
        minX = [x,minX].min
        maxX = [x,maxX].max
        minY = [y,minY].min
        maxY = [y,maxY].max
      end

      # debug
      # (0..maxY).each do |yi|
      #   line = ""
      #   (0..maxX).each do |xi|
      #     if elfMap[[xi,yi]]
      #       line += "#"
      #     else
      #       line += "."
      #     end
      #   end
      #   puts line
      # end
      # puts ""
    end
  end
end

#3709 low

Day23.new