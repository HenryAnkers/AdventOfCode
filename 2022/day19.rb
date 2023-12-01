class Day19
  def initialize
    input = File.read('input/day19.txt').split("\n")

    puts part1(input)
    puts part2(input)
  end

  def parseRobotCosts(input)
    input.map do |x|
      robots = x.split(".").map{ |x| x.split(" ")}
      robotCosts = []
      robotCosts << [robots[0][6].to_i, 0, 0]
      robotCosts <<[robots[1][4].to_i, 0, 0]
      robotCosts << [robots[2][4].to_i, robots[2][7].to_i, 0]
      robotCosts << [robots[3][4].to_i, 0, robots[3][7].to_i]

      robotCosts
    end
  end

  def getHashKey(robots, inventory, time)
    hashKeys = ["#{time}"]

    robotString = ""
    robotString += "OR:#{robots["ORE"]}"
    robotString += "CL:#{robots["CLAY"]}"
    robotString += "OB:#{robots["OBSIDIAN"]}"
    robotString += "GE:#{robots["GEODE"]}"
    hashKeys << robotString

    inv = ""
    inv += "OR:#{inventory[0]}"
    inv += "CL:#{inventory[1]}"
    inv += "OB:#{inventory[2]}"
    inv += "GE:#{inventory[3]}"
    hashKeys << inv

    return hashKeys.join("+")
  end

  def canBuildRobot(robotIndex, inventory, robotCosts, robots)
    maxValues = [0,0,0]
    robotCosts.each do |x|
      maxValues[0] = [x[0],maxValues[0]].max
      maxValues[1] = [x[1],maxValues[1]].max
      maxValues[2] = [x[2],maxValues[2]].max
    end

    if robotIndex == 0 && robots["ORE"] >= maxValues[0]
      return false
    end
    if robotIndex == 1 && robots["CLAY"] >= maxValues[1]
      return false
    end
    if robotIndex == 2 && robots["OBSIDIAN"] >= maxValues[2]
      return false
    end

    cost = robotCosts[robotIndex]
    return cost[0] <= inventory[0] && cost[1] <= inventory[1] && cost[2] <= inventory[2]
  end

  def buildRobot(cost, inventory)
    inventory[0] -= cost[0]
    inventory[1] -= cost[1]
    inventory[2] -= cost[2]

    return inventory
  end

  def addNewMineralsToInventory(inventory, robots)
    inventory[0] += robots["ORE"]
    inventory[1] += robots["CLAY"]
    inventory[2] += robots["OBSIDIAN"]
    inventory[3] += robots["GEODE"]

    return inventory
  end

  def trimInventory(inventory, time, robotCosts)  #trim the inventory to the maximum amount of resources we can possibly spend, which is the maximum cost for each resource * remaining amount of time
    maxValues = [0,0,0]
    robotCosts.each do |x|
      maxValues[0] = [x[0],maxValues[0]].max
      maxValues[1] = [x[1],maxValues[1]].max
      maxValues[2] = [x[2],maxValues[2]].max
    end

    inventory[0] = [inventory[0], maxValues[0] * time].min
    inventory[1] = [inventory[1], maxValues[1] * time].min
    inventory[2] = [inventory[2], maxValues[2] * time].min
  end

  def dp(memo, bestAtTime, robotCosts, robots, inventory, time)
    if time <= 10
      trimInventory(inventory, time, robotCosts)
    end

    hashKey = getHashKey(robots, inventory, time) #key for this specific state
    memod = memo[hashKey]
    if memod != -1
      return memod
    end

    if time == 0
      return inventory[3]
    end

    if bestAtTime[time] != 0 && (inventory[3] + (robots["GEODE"] * time) + ((time-1) * time).div(2)) < bestAtTime[time]
      return bestAtTime[time]
    end

    possibleSolutions = []

    if canBuildRobot(3, inventory, robotCosts, robots)
      newInventory = buildRobot(robotCosts[3], inventory[0..-1])
      addNewMineralsToInventory(newInventory, robots)
      newRobots = robots.clone
      newRobots["GEODE"] += 1

      ans = dp(memo, bestAtTime, robotCosts, newRobots, newInventory, time-1)
      memo[hashKey] = ans
      bestAtTime[time] = [bestAtTime[time], ans].max
      return ans
    else  #you only try building nothing if you can't build an geode bot
      naiveInventory = addNewMineralsToInventory(inventory[0..-1], robots)
      possibleSolutions << dp(memo, bestAtTime, robotCosts, robots.clone, naiveInventory, time-1)
    end

    if canBuildRobot(2, inventory, robotCosts, robots)
      newInventory = buildRobot(robotCosts[2], inventory[0..-1])
      addNewMineralsToInventory(newInventory, robots)
      newRobots = robots.clone
      newRobots["OBSIDIAN"] += 1

      possibleSolutions << dp(memo, bestAtTime, robotCosts, newRobots, newInventory, time-1)
    end

    if canBuildRobot(1, inventory, robotCosts, robots)
      newInventory = buildRobot(robotCosts[1], inventory[0..-1])
      addNewMineralsToInventory(newInventory, robots)
      newRobots = robots.clone
      newRobots["CLAY"] += 1
      possibleSolutions << dp(memo, bestAtTime, robotCosts, newRobots, newInventory, time-1)
    end

    if canBuildRobot(0, inventory, robotCosts, robots)
      newInventory = buildRobot(robotCosts[0], inventory[0..-1])
      addNewMineralsToInventory(newInventory, robots)
      newRobots = robots.clone
      newRobots["ORE"] += 1
      possibleSolutions << dp(memo, bestAtTime, robotCosts, newRobots, newInventory, time-1)
    end

    ans = possibleSolutions.max
    bestAtTime[time] = [bestAtTime[time], ans].max
    memo[hashKey] = ans

    return ans
  end

  def part1(input)
    # only 30 blueprints means the algo is probably going to be as slow as we expect..
    # feeling DP again
    # we start with 1 ore collecting robot, we need to build as many geode cracking robots as possible basically
    # at every step we can choose to build a robot if we have the resource
    # as there are 24 minutes I can't see any scenario we wouldn't want to build a geode one when we have the minerals, this cuts down on the number of states
    # we'll cache a state of {"time","numberOfRobotsOfEachType","inventory"}
    # we can make a total of 4 robots every step, gives us like 4^24 possibilities for robots max... and I don't know how much memoization will help :/
    # it feels this can also be solved mathematically. what type of robot should I optimially build next at any given stage? idk how to work this out though

    # this didn't come close to working without cutting down on the number of states we've got to process
    # 1) dont try not building a robot if you can potentially build something 'useful', idk what useful is though. saying if we didnt build a geode robot is a safe bet
    # 2) prune inventory towards the end so you hit more states in the cache
    # 3) keep track of the maximum value you have found at a given time, work out if you can improve on it in your current branch (n + n-1 + n-2... triangle)

    blueprints = parseRobotCosts(input)
    answer = 0

    blueprints.each_with_index do |b, i|
      robotsMap = Hash.new(0)
      robotsMap["ORE"] = 1
      
      opened = dp(Hash.new(-1),Hash.new(0), b, robotsMap, [0,0,0,0], 24)
      answer += (i+1) * opened
    end

    answer
  end

  def part2(input)
    blueprints = parseRobotCosts(input).take(3)
    answer = 1

    blueprints.each_with_index do |b, i|
      robotsMap = Hash.new(0)
      robotsMap["ORE"] = 1

      opened = dp(Hash.new(-1), Hash.new(0), b, robotsMap, [0,0,0,0], 32)
      answer = answer * opened
    end

    answer
  end
end

Day19.new