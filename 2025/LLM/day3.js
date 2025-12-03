// Created by GPT-5.1-Codex-Max
// Improvements on original solution:
// - Parses once into Uint8Array digits (trimmed input) to eliminate per-line split/map overhead.
// - Part 1 uses single-pass prefix/suffix maxima instead of repeated slices and Math.max spreads.
// - Part 2 precomputes next-occurrence tables to pick the lexicographically max 12-digit subsequence in O(10 * targetLen) per line.

const path = require('path');
const fs = require('fs');
const { performance } = require('perf_hooks');

const TARGET_LENGTH = 12;

function parseInput(day) {
    const text = fs.readFileSync(path.join(__dirname, '..', `inputs/day${day}.txt`), 'utf8').trim();
    const lines = text.split('\n');

    const parsed = new Array(lines.length);
    for (let i = 0; i < lines.length; i++) {
        const line = lines[i];
        const digits = new Uint8Array(line.length);
        for (let j = 0; j < line.length; j++) {
            digits[j] = line.charCodeAt(j) - 48;
        }
        parsed[i] = digits;
    }

    return parsed;
}

function part1(entries) {
    let total = 0;

    for (let i = 0; i < entries.length; i++) {
        const digits = entries[i];
        const len = digits.length;

        let maxDigit = digits[0];
        let maxIdx = 0;
        for (let j = 1; j < len - 1; j++) {
            const d = digits[j];
            if (d > maxDigit) {
                maxDigit = d;
                maxIdx = j;
            }
        }

        let suffixMax = digits[maxIdx + 1];
        for (let j = maxIdx + 2; j < len; j++) {
            const d = digits[j];
            if (d > suffixMax) suffixMax = d;
        }

        total += maxDigit * 10 + suffixMax;
    }

    return total;
}

function buildNextIndices(digits) {
    const len = digits.length;
    const next = Array.from({ length: len + 1 }, () => new Int16Array(10));
    next[len].fill(-1);

    for (let i = len - 1; i >= 0; i--) {
        next[i].set(next[i + 1]);
        next[i][digits[i]] = i;
    }

    return next;
}

function part2(entries) {
    let total = 0;

    for (let i = 0; i < entries.length; i++) {
        const digits = entries[i];
        const len = digits.length;
        const next = buildNextIndices(digits);

        let start = 0;
        let remaining = TARGET_LENGTH;
        let value = 0;

        while (remaining > 0) {
            const lastAllowed = len - remaining;
            let chosenDigit = 0;
            let chosenIdx = -1;

            for (let d = 9; d >= 0; d--) {
                const candidateIdx = next[start][d];
                if (candidateIdx !== -1 && candidateIdx <= lastAllowed) {
                    chosenDigit = d;
                    chosenIdx = candidateIdx;
                    break;
                }
            }

            value = value * 10 + chosenDigit;
            start = chosenIdx + 1;
            remaining--;
        }

        total += value;
    }

    return total;
}

const filename = path.basename(__filename);
const day = filename.match(/^day(\d+)/)[1];

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
