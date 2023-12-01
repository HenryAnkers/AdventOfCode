require 'set'

class Day16
  def initialize
    input = File.read('input/day16.txt').split("\n")

    solve(input)
  end


    def DP(memo, nodes, currentNode, time, score, openedValves)
      memod = memo["#{currentNode[:name]},#{time},#{score}"]
      if memod != -1
        return memod, openedValves
      end

      if time <= 1
        return score, openedValves
      end

      currentMax = score
      newOpenedValves = openedValves[0..-1]

      currentNode[:neighbours].each do |n|
        nextNode = nodes[n]

        potentialMax, potentialOV = DP(memo, nodes, nextNode, time-1, score, openedValves[0..-1])

        if currentNode[:flow] > 0 && !openedValves.include?(currentNode[:name])
          openPotentialMax, openPotentialOV = DP(memo, nodes, nextNode, time-2, score + (currentNode[:flow] * (time-1)), openedValves + [currentNode[:name]])
          if openPotentialMax > potentialMax
            potentialMax = openPotentialMax
            potentialOV = openPotentialOV
          end
        end

        if potentialMax > currentMax
          currentMax = potentialMax
          newOpenedValves = potentialOV
        end
      end

      memo["#{currentNode[:name]},#{time},#{score}"] = currentMax
      return currentMax, newOpenedValves
    end


  def solve(input)
    # at every step we have two choices
    # 1) open the value and get (RM * VALVE) valve where RM is remaining minutes
    # 2) move to a different valve and don't get (RM * VALUE) but you do get an extra minute...
    # can we model this as dynamic programming, VALVE,MINUTES? there is 50 * 30 max calculations if we memo?

    # no because we can arrive at the same valve at the same time but have a different score calculated
    # BUT just adding in score as an extra memo parameter works though lol

    nodes = Hash.new() #store an object consisting of a value and neighbours (just strings we can use to lookup nodes with)

    input.each do |line|
      name = line[6..7]
      flow = line.split("rate=")[1].split(";")[0].to_i
      neighbours = line.split(";")[1].split(" ").drop(4).join(" ").split(", ")

      nodes[name] = { name: name, flow: flow, neighbours: neighbours }
    end


    potentialAnswer, visited = DP(Hash.new(-1), nodes, nodes["AA"], 30, 0, [])
    puts potentialAnswer

    #part 2
    # don't know how to do this apart from bruteforce
    # the two optimum paths might not be the two local optimal paths (so my greedy initial path is A-B-C, elephants is therefore D-E, but in reality most efficient is A-B-E  D-C or something) or we can just run it twice
    # dont know how to tell, not sure we can just run them seperately...
    # THERES ONLY LIKE 7 RETURNED NODES IN MY FIRST ANSWER SO THEY PROBABLY ARE DISJOINT??? 
    #
    # IT WORKS!!!!
    potentialAnswer1, visited1 = DP(Hash.new(-1), nodes, nodes["AA"], 26, 0, [])
    potentialAnswer2, visited2 = DP(Hash.new(-1), nodes, nodes["AA"], 26, 0, visited1)

    puts potentialAnswer1 + potentialAnswer2
  end
end


Day16.new