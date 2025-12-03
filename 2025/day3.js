const path = require('path');
const fs = require('fs');
const { performance } = require('perf_hooks');

function parseInput(day) {
    const inputPath = `./inputs/day${day}.txt`;
    const text = fs.readFileSync(inputPath, 'utf8');
    return text.split("\n")
}

function part1(input) {
    let count = 0;

    input.forEach(x => {
        let numbers = x.split('').map(Number)
        let firstNumberCandidates = numbers.slice(0, numbers.length-1)
        let maxFirst = Math.max(...firstNumberCandidates)
        let maxFirstIndex = firstNumberCandidates.indexOf(maxFirst)

        let secondNumberCandidates = numbers.slice(maxFirstIndex+1)
        let maxSecond = Math.max(...secondNumberCandidates)

        count += 10*maxFirst + maxSecond
       
    });

    return count
}

function part2(input) {
    let count = 0;

    input.forEach(x => {
        let numbers = x.split('').map(Number)
        let numberToFind = 12
        let startIndex = 0
        let tempCount = 0

        while (numberToFind > 0) {
            let candidates = numbers.slice(startIndex, numbers.length-(numberToFind-1))
            let maxForThisDigit  = Math.max(...candidates)
            let maxForThisDigitIndex = candidates.indexOf(maxForThisDigit)

            startIndex += maxForThisDigitIndex + 1
            tempCount += (10**(numberToFind-1))*maxForThisDigit
            numberToFind -= 1
        }

        count += tempCount 
    });

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