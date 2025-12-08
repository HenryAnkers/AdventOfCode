const path = require('path');
const fs = require('fs');
const { performance } = require('perf_hooks');

function parseInput(day) {
    const inputPath = `./inputs/day${day}.txt`;
    const text = fs.readFileSync(inputPath, 'utf8');
    let input = text.split("\n").map(line => line.split(",").map(x => Number(x)));

    let distances = [];
    for (let i = 0; i < input.length; i++) {
        for (let j = i+1; j < input.length; j++) {
            let a = input[i];
            let b = input[j];
            let distance = Math.sqrt((a[0]-b[0])**2 + (a[1]-b[1])**2 + (a[2]-b[2])**2);
            distances.push([i, j, distance]);
        }
    }
    distances.sort((a, b) => (a[2] - b[2]));
    return [input, distances];
}

function part1(input, distances) {
    let count = 0;
    let circuitMap = {};

    for (let i = 0; i < 1000; i++) {
        let distanceToProcess = distances[i];
        let a = distanceToProcess[0];
        let b = distanceToProcess[1];

        let aCircuit = circuitMap[a];
        let bCircuit = circuitMap[b];
        if (typeof aCircuit == "undefined" && typeof bCircuit == "undefined") {
            let newCircuit = new Set([a,b]);
            circuitMap[a] = newCircuit;
            circuitMap[b] = newCircuit;
        } else if (typeof aCircuit == "undefined") {
            circuitMap[b].add(a);
            circuitMap[a] = circuitMap[b];
        } else if (typeof bCircuit == "undefined") {
            circuitMap[a].add(b);
            circuitMap[b] = circuitMap[a];
        } else {
            let circuit = circuitMap[b].union(circuitMap[a]);
            circuit.forEach(x => {
                circuitMap[x] = circuit;
            });
        };
    }

    let uniqueCircuits = new Set(Object.values(circuitMap));
    let circuitSizes = Array.from(uniqueCircuits).map(x => x.size).sort((a, b) => (b - a));
    count = circuitSizes[0] * circuitSizes[1] * circuitSizes[2];

    return count;
}

function part2(input, distances) {
    let circuitMap = {};    

    for (let i = 0; i < input.length**2; i++) {
        let distanceToProcess = distances[i];
        let a = distanceToProcess[0];
        let b = distanceToProcess[1];

        let aCircuit = circuitMap[a];
        let bCircuit = circuitMap[b];
        if (typeof aCircuit == "undefined" && typeof bCircuit == "undefined") {
            let newCircuit = new Set([a,b]);
            circuitMap[a] = newCircuit;
            circuitMap[b] = newCircuit;
        } else if (typeof aCircuit == "undefined") {
            circuitMap[b].add(a);
            circuitMap[a] = circuitMap[b];
        } else if (typeof bCircuit == "undefined") {
            circuitMap[a].add(b);
            circuitMap[b] = circuitMap[a];
        } else {
            let circuit = circuitMap[b].union(circuitMap[a]);
            circuit.forEach(x => {
                circuitMap[x] = circuit;
            });
        };

        if (circuitMap[a].size == input.length) {
            return input[a][0] * input[b][0];
        }
    }
}

const filename = path.basename(__filename);
const day = filename.split('.')[0].replace('day', '');

let start = performance.now();
const inputData = parseInput(day);
let end = performance.now();
console.log(`Day ${day} - Parsing took ${(end - start).toFixed(6)} ms`);

start = performance.now();
const solution1 = part1(inputData[0], inputData[1]);
end = performance.now();
console.log(`Day ${day} - Part 1: ${solution1} (took ${(end - start).toFixed(6)} ms)`);

start = performance.now();
const solution2 = part2(inputData[0], inputData[1]);
end = performance.now();
console.log(`Day ${day} - Part 2: ${solution2} (took ${(end - start).toFixed(6)} ms)`);