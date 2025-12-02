const path = require('path');
const fs = require('fs');
const { performance } = require('perf_hooks');

function parseInput(day) {
    const inputPath = `./inputs/day${day}.txt`;
    const text = fs.readFileSync(inputPath, 'utf8');
    return text.split(",")
}

function part1(input) {
    let count = 0;

    input.forEach(x => {
        let startInt = parseInt(x.split("-")[0])
        let endInt = parseInt(x.split("-")[1])

        for (let i = startInt; i <= endInt; i++) {
            let stringTarget = i.toString()
            if (stringTarget.length % 2) {
                continue;
            }

            let midPoint = Math.floor(stringTarget.length / 2);
            let firstHalf = stringTarget.slice(0, midPoint)
            let endHalf = stringTarget.slice(midPoint)

            if (firstHalf === endHalf) {
                count += i
            }
        }
    });

    return count
}

function part2(input) {
    let count = 0;

    input.forEach(x => {
        let startInt = parseInt(x.split("-")[0])
        let endInt = parseInt(x.split("-")[1])

        mainLoop:
        for (let i = startInt; i <= endInt; i++) {
            let stringTarget = i.toString()
            let midPoint = Math.floor(stringTarget.length / 2);

            for (let j = 1; j <= midPoint; j++) {
                if (stringTarget.length % j) {
                    continue;
                }

                let potentialPattern = stringTarget.slice(0, j) 
                let potentialOccurrences = stringTarget.length / j
                if (potentialPattern.repeat(potentialOccurrences) === stringTarget) {
                    count += i
                    continue mainLoop
                }        
            }
        }
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