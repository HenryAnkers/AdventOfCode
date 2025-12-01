const path = require('path');
const fs = require('fs');
const { performance } = require('perf_hooks');

function parseInput(day) {
    const inputPath = `./inputs/day${day}.txt`;
    const text = fs.readFileSync(inputPath, 'utf8');
    return text.split("\n")
}

function part1(input) {
    return ""
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