const path = require('path');
const fs = require('fs');
const { performance } = require('perf_hooks');

function parseInput(day) {
    const inputPath = `./inputs/day${day}.txt`;
    const text = fs.readFileSync(inputPath, 'utf8');
    return text.split("\n\n").map(line => line.split("\n"));
}

function part1(input) {
    let count = 0;
    let ranges = input[0].map(r => r.split("-").map(x => Number(x)));
    let ingredients = input[1].map(i => Number(i));

    ingredients.forEach(ingredient => {
        for (let i = 0; i < ranges.length; i++) {
            let range = ranges[i];
            if (ingredient >= range[0] && ingredient <= range[1]) {
                count++
                return;
            }
        }
    });

    return count;
}

function part2(input) {
    let ranges = input[0].map(r => r.split("-").map(x => Number(x)));

    mainLoop:
    for (let i = 0; i < ranges.length; i++) {
        let range = ranges[i];
        
        let pointer = i + 1;
        while (pointer < ranges.length) {
            let mR = ranges[pointer];
            let rangesAreExclusive = (mR[1] < range[0]) || (mR[0] > range[1]);
            if (!rangesAreExclusive) {
                    let newMin = Math.min(range[0], mR[0]);
                    let newMax = Math.max(range[1], mR[1]);
                    ranges[pointer] = [newMin, newMax];
                    ranges[i] = null;
                    continue mainLoop;
                }
            pointer++;
        }
    }

    return ranges
        .filter(x => x != null)
        .reduce((c, r) => c + (r[1] - r[0] + 1), 0);
}

function part2Bruteforce(input) {
    let count = 0;
    let ranges = input[0].map(r => r.split("-").map(x => Number(x)));za

    let min = Math.min(...(ranges.map(r => r[0])));
    let max = Math.max(...(ranges.map(r => r[1])));

    mainLoop:
    for (let n = min; n <= max; n++) {
        for (let i = 0; i < ranges.length; i++) {
            let range = ranges[i];
            if (n >= range[0] && n <= range[1]) {
                count++
                continue mainLoop;
            }
        }
    }

    return count;
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