const path = require('path');
const fs = require('fs');
const { performance } = require('perf_hooks');

function parseInput(day) {
    const inputPath = `./inputs/day${day}.txt`;
    const text = fs.readFileSync(inputPath, 'utf8');
    return text.split("\n").map(line => line.split(""));
}

function part1(input) {
    let count = 0;
    input[1][input[0].indexOf("S")] = "|";
    let currentY = 1;
    let queueX = [input[0].indexOf("S")];

    while (currentY < input.length - 1) {
        let nextQueue = [];

        queueX.forEach(x => {
            let nextCoord = input[currentY+1][x];
            
            if (nextCoord == ".") {
                input[currentY+1][x] = "|";
                nextQueue.push(x);
            } else if (nextCoord == "^") {
                count++;
                if (input[currentY+1][x-1] == ".") {
                    input[currentY+1][x-1] = "|";
                    nextQueue.push(x-1);
                } 
                if (input[currentY+1][x+1] == ".") {
                    input[currentY+1][x+1] = "|";
                    nextQueue.push(x+1);
                } 
            } 
        });

        currentY += 1;
        queueX = nextQueue;
    }

    return count;
}


function part2(input) {
    let count = 0;
    input[1][input[0].indexOf("S")] = 1;
    let currentY = 1;
    let queueX = [input[0].indexOf("S")];
    
    while (currentY < input.length - 1) {
        let nextQueue = [];

        queueX.forEach(x => {
            let currentWays = input[currentY][x];   
            let nextCoord = input[currentY+1][x];
            
            if (nextCoord == ".") {
                input[currentY+1][x] = currentWays;
                nextQueue.push(x);
            } else if (nextCoord == "^") {
                if (input[currentY+1][x-1] == ".") {
                    input[currentY+1][x-1] = currentWays;
                    nextQueue.push(x-1);
                } else {
                    input[currentY+1][x-1] += currentWays;
                }

                if (input[currentY+1][x+1] == ".") {
                    input[currentY+1][x+1] = currentWays;
                    nextQueue.push(x+1);
                } else {
                    input[currentY+1][x+1] += currentWays;
                }
            } else {
                input[currentY+1][x] += currentWays;
            }
        });

        currentY += 1;
        queueX = nextQueue;
    }

    queueX.forEach(x => { 
        count += input[currentY][x];
    });

    return count;
}

const filename = path.basename(__filename);
const day = filename.split('.')[0].replace('day', '');

let start = performance.now();
let inputData = parseInput(day);
let end = performance.now();
console.log(`Day ${day} - Parsing took ${(end - start).toFixed(6)} ms`);

start = performance.now();
const solution1 = part1(inputData);
end = performance.now();
console.log(`Day ${day} - Part 1: ${solution1} (took ${(end - start).toFixed(6)} ms)`);

start = performance.now();
inputData = parseInput(day);
const solution2 = part2(inputData);
end = performance.now();
console.log(`Day ${day} - Part 2: ${solution2} (took ${(end - start).toFixed(6)} ms)`);
