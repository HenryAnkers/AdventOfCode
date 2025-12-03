// Created by GPT-5.1-Codex-Max
// Improvements on original solution:
// - Parses once into [dir, magnitude] tuples to avoid repeated string slicing and implicit globals.
// - Uses pure arithmetic normalization (no branching modulo quirks) for the 0–99 dial.
// - Counts full 100-step rotations up front and handles wrap detection in a single pass.

const path = require('path');
const fs = require('fs');
const { performance } = require('perf_hooks');

function parseInput(day) {
    const inputPath = `./inputs/day${day}.txt`;
    const text = fs.readFileSync(inputPath, 'utf8').trim();
    const lines = text.split('\n');
    return lines.map(line => {
        const dir = line.charCodeAt(0) === 76 ? -1 : 1; // 'L' => -1, 'R' => 1
        const mag = Number(line.slice(1));
        return [dir, mag];
    });
}

function normalizeDial(pos) {
    pos %= 100;
    return pos < 0 ? pos + 100 : pos;
}

function part1(instructions) {
    let dial = 50;
    let count = 0;

    for (let i = 0; i < instructions.length; i++) {
        const [dir, mag] = instructions[i];
        dial = normalizeDial(dial + dir * mag);
        if (dial === 0) count++;
    }

    return count;
}

function part2(instructions) {
    let dial = 50;
    let count = 0;

    for (let i = 0; i < instructions.length; i++) {
        const [dir, mag] = instructions[i];

        const fullTurns = Math.floor(mag / 100);
        count += fullTurns;
        const step = mag % 100;

        const prev = dial;
        dial += dir * step;
        if (dial > 99 || (dial <= 0 && prev !== 0)) {
            count += 1;
        }

        dial = normalizeDial(dial);
    }

    return count;
}

const filename = path.basename(__filename);
const day = filename.split('.')[0].replace('day', '');

let start = performance.now();
const inputData = parseInput(day.replace('_gpt-5', ''));
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
