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

    for (let y = 0; y < input.length; y++) {
        for (let x = 0; x < input[y].length; x++) {
            if (input[y][x] != "@") {
                continue;
            }
            let adjacentCount = 0

            for (let dy = -1; dy < 2; dy++) {
                for (let dx = -1; dx < 2; dx++) {
                    let ty = y + dy;
                    let tx = x + dx;

                    if (ty < 0 || ty >= input.length) {
                        continue;
                    }
                    if (tx < 0 || tx >= input[y].length) {
                        continue;
                    }
                    if (dx == 0 && dy == 0) {
                        continue;
                    }

                    if (input[ty][tx] == "@") {
                        adjacentCount++;
                    }
                }
            }

            if (adjacentCount < 4) {
                count++;
            }
        }
    }

    return count;
}

function part2(input) {
    let totalCount = 0;
    let roundCount = 0;

    mainLoop:
    while (true) {
        for (let y = 0; y < input.length; y++) {
            for (let x = 0; x < input[y].length; x++) {
                if (input[y][x] != "@") {
                    continue;
                }
                let adjacentCount = 0

                for (let dy = -1; dy < 2; dy++) {
                    for (let dx = -1; dx < 2; dx++) {
                        let ty = y + dy;
                        let tx = x + dx;

                        if (ty < 0 || ty >= input.length) {
                            continue;
                        }
                        if (tx < 0 || tx >= input[y].length) {
                            continue;
                        }
                        if (dx == 0 && dy == 0) {
                            continue;
                        }

                        if (input[ty][tx] == "@") {
                            adjacentCount++;
                        }
                    }
                }

                if (adjacentCount < 4) {
                    totalCount++;
                    roundCount++
                    input[y][x] = "."
                }
            }
        }
        if (!roundCount) {
            break mainLoop
        } else {
            roundCount = 0
        }
    }
    

    return totalCount;
}

const filename = path.basename(__filename);
const day = filename.split('.')[0].replace('day', '');

let start = performance.now();
const inputData = parseInput(day);
let end = performance.now();
console.log(`Day ${day} - Parsing took ${(end - start).toFixed(6)} ms`);

start = performance.now();
const solution1 = part1(inputData);
end = performance.now();
console.log(`Day ${day} - Part 1: ${solution1} (took ${(end - start).toFixed(6)} ms)`);

start = performance.now();
const solution2 = part2(inputData);
end = performance.now();
console.log(`Day ${day} - Part 2: ${solution2} (took ${(end - start).toFixed(6)} ms)`);
