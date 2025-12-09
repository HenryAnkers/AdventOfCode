const path = require('path');
const fs = require('fs');
const { performance } = require('perf_hooks');

function parseInput(day) {
    const inputPath = `./inputs/day${day}.txt`;
    const text = fs.readFileSync(inputPath, 'utf8');
    return text.split("\n").map(line => line.split(",").map(x => Number(x)));
}

function part1(input) {
    let count = 0;

    for (let i = 0; i < input.length; i++) {
        for (let j = i; j < input.length; j++) {
            let diffX = Math.abs(input[i][0] - input[j][0]) + 1;
            let diffY = Math.abs(input[i][1] - input[j][1]) + 1;
            count = Math.max(count, diffX * diffY)
        }
    }

    return count;
}

function part2(input) {
    return ""
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