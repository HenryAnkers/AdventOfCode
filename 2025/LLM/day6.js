// Created by GPT-5.1-Codex-Max
// Improvements on original solution:
// - Detects problem spans via true blank columns across all rows (handles ragged spacing robustly).
// - Pads rows once and accumulates digits directly to avoid substring/Number churn.
// - Finds operators by scanning within each span, so misaligned symbols are still read correctly.
// - Part 2 builds column-wise values and skips only genuinely empty columns, preserving zeros.

const path = require('path');
const fs = require('fs');
const { performance } = require('perf_hooks');

function parseInput(day) {
    const text = fs.readFileSync(path.join(__dirname, '..', `inputs/day${day}.txt`), 'utf8').trimEnd();
    const rawLines = text.split('\n');
    const height = rawLines.length - 1; // last row is operators

    let width = 0;
    for (let i = 0; i < rawLines.length; i++) {
        if (rawLines[i].length > width) width = rawLines[i].length;
    }

    const rows = new Array(rawLines.length);
    for (let i = 0; i < rawLines.length; i++) {
        rows[i] = rawLines[i].padEnd(width, ' ');
    }

    // Identify problem spans: contiguous columns that are not entirely blank.
    const spans = [];
    let x = 0;
    while (x < width) {
        let isBlank = true;
        for (let y = 0; y <= height; y++) {
            if (rows[y].charCodeAt(x) !== 32) {
                isBlank = false;
                break;
            }
        }
        if (isBlank) {
            x++;
            continue;
        }

        const start = x;
        x++;
        while (x < width) {
            let blankCol = true;
            for (let y = 0; y <= height; y++) {
                if (rows[y].charCodeAt(x) !== 32) {
                    blankCol = false;
                    break;
                }
            }
            if (blankCol) break;
            x++;
        }
        spans.push([start, x - 1]);
    }

    return { rows, spans, height, opRow: rows[height] };
}

function part1(data) {
    const { rows, spans, height, opRow } = data;
    let total = 0n;

    for (let s = 0; s < spans.length; s++) {
        const [start, end] = spans[s];
        let op = '+';
        for (let x = start; x <= end; x++) {
            const c = opRow.charAt(x);
            if (c === '+' || c === '*') {
                op = c;
                break;
            }
        }
        let acc = op === '+' ? 0n : 1n;

        for (let y = 0; y < height; y++) {
            const line = rows[y];
            let value = 0n;
            let seen = false;

            for (let x = start; x <= end; x++) {
                const code = line.charCodeAt(x);
                if (code !== 32) {
                    value = value * 10n + BigInt(code - 48);
                    seen = true;
                }
            }

            if (!seen) continue;
            acc = op === '+' ? acc + value : acc * value;
        }

        total += acc;
    }

    return total;
}

function part2(data) {
    const { rows, spans, height, opRow } = data;
    let total = 0n;

    for (let s = 0; s < spans.length; s++) {
        const [start, end] = spans[s];
        let op = '+';
        for (let x = start; x <= end; x++) {
            const c = opRow.charAt(x);
            if (c === '+' || c === '*') {
                op = c;
                break;
            }
        }
        let acc = op === '+' ? 0n : 1n;

        for (let x = start; x <= end; x++) {
            let value = 0n;
            let seen = false;

            for (let y = 0; y < height; y++) {
                const code = rows[y].charCodeAt(x);
                if (code !== 32) {
                    value = value * 10n + BigInt(code - 48);
                    seen = true;
                }
            }

            if (!seen) continue;
            acc = op === '+' ? acc + value : acc * value;
        }

        total += acc;
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
