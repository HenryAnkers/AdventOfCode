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
    let columnStart = 0;
    let opLine = input[input.length - 1];

    for (let i = 1; i < opLine.length; i++) {
        if (i != opLine.length - 1 && opLine[i+1] == " ") {
            continue;
        }

        let symbol = opLine[columnStart];
        let sum = symbol == "+" ? 0 : 1;
        for (let j = 0; j < input.length - 1; j++) {
            let num = Number(input[j].slice(columnStart, i+1));
            if (symbol == "+")  {
                sum += num;
            } else {
                sum *= num;
            }
        }
        columnStart = i+1;
        count += sum;
    }
    
    return count;
}

function part2(input) {
    let count = 0;
    let columnStart = 0;
    let opLine = input[input.length - 1];

    for (let i = 1; i < opLine.length; i++) {
        if (i != opLine.length - 1 && opLine[i+1] == " ") {
            continue;
        }

        let symbol = opLine[columnStart];
        let sum = symbol == "+" ? 0 : 1;
        
        for (let j = columnStart; j < i+1; j++) {
            let num = 0;
            
            for (let k = 0; k < input.length - 1; k++) {
                if (input[k][j] == " ") {
                    continue;
                }
                num *= 10;
                num += Number(input[k][j]);
            }
            if (symbol == "+")  {
                sum += num;
            } else if (num != 0) {
                sum *= num;
            }
        }

        columnStart = i+1;
        count += sum;
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