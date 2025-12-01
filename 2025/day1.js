const path = require('path');
const fs = require('fs');
const { performance } = require('perf_hooks');

function parseInput(day) {
    const inputPath = `./inputs/day${day}.txt`;
    const text = fs.readFileSync(inputPath, 'utf8');
    return text.split("\n");
}

function part1(input) {
    dialLocation = 50;
    count = 0;

    input.forEach(x => {
        dir = x[0] == "L" ? -1 : 1
        mag = x.slice(1)

        dialLocation += (dir * mag) 
        dialLocation = dialLocation % 100
        if (dialLocation < 0) {
            dialLocation = 100 + dialLocation
        }
        if (dialLocation == 0) {
            count++
        }
    })

    return count
}

function part2(input) {
    dialLocation = 50;
    count = 0;

    input.forEach(x => {
        dir = x[0] == "L" ? -1 : 1
        mag = x.slice(1)

        totalRotations = Math.floor(mag / 100)
        count += totalRotations
        mag = mag % 100

        prevLocation = dialLocation
        dialLocation += (dir * mag) 
        if (dialLocation > 99 || (dialLocation <= 0 && prevLocation != 0)) {
            count += 1
        }

        dialLocation = dialLocation % 100
        if (dialLocation < 0) {
            dialLocation = 100 + dialLocation
        }
    })

    return count
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