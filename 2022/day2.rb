class Day2
  def initialize
    input = File.read('input/day2.txt')

    puts part1(input)
    puts part2(input)
  end

  def convertSymbol(x)
    return {'X' => 'A', 'Y' => 'B', 'Z' => 'C'}[x]
  end

  def winningMoveAgainst(x)
    return {'A' => 'B', 'B' => 'C', 'C' => 'A' }[x]
  end

  def losingMoveAgainst(x)
    return {'A' => 'C', 'B' => 'A', 'C' => 'B' }[x]
  end

  def moveScores(x)
    return {'A' => 1, 'B' => 2, 'C' => 3}[x]
  end
  
  def resultScores(x)
    return {'X' => 0, 'Y' => 3, 'Z' => 6}[x]
  end

  def part1(input)
    total = 0

    games = input.split("\n")
    games.each do |game|
      players = game.split

      player0 = players[0]
      player1 = convertSymbol(players[1])
      total += moveScores(player1)

      if player0 == player1
        total += 3
      elsif player0 != winningMoveAgainst(player1)
        total += 6
      end
    end

    total
  end

  def part2(input)
    total = 0

    games = input.split("\n")
    games.each do |game|
      params = game.split

      opponentsMove = params[0]
      result = params[1]
      
      resultScore = resultScores(result)
      total += resultScore
      if resultScore == 0
        total += moveScores(losingMoveAgainst(opponentsMove))
      elsif  resultScore == 6
        total += moveScores(winningMoveAgainst(opponentsMove))
      else
        total += moveScores(opponentsMove)
      end
    end

    total
  end
end

Day2.new