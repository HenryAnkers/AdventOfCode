// Created by GPT-5.1-Codex-Max
// Improvements on original solution:
// - Parses into a flat splitter mask with start coordinates; avoids mutating string grids each step.
// - Simulates beams row-by-row with compact column lists and per-row de-dupe to keep merges O(k).
// - Part 2 accumulates timeline counts with BigInt and column touch lists so collisions sum without rescans.
// - Reuses buffers across iterations to minimize allocations while keeping control flow identical to the puzzle.

const path = require('path');
const fs = require('fs');
const { performance } = require('perf_hooks');

function parseInput(day) {
    const text = fs.readFileSync(path.join(__dirname, '..', `inputs/day${day}.txt`), 'utf8').trimEnd();
    const lines = text.split('\n');
    const height = lines.length;
    const width = lines[0].length;
    const splitters = new Uint8Array(width * height);

    let startX = -1;
    let startY = -1;

    for (let y = 0; y < height; y++) {
        const line = lines[y];
        for (let x = 0; x < width; x++) {
            const ch = line.charCodeAt(x);
            if (ch === 94) { // '^'
                splitters[y * width + x] = 1;
            } else if (ch === 83) { // 'S'
                startX = x;
                startY = y;
            }
        }
    }

    if (startX === -1) {
        throw new Error('Start position S not found in input');
    }

    return { splitters, width, height, startX, startY };
}

function part1(state) {
    const { splitters, width, height, startX, startY } = state;
    let currentRow = startY + 1;
    if (currentRow >= height) return 0;

    let active = new Int32Array(width);
    let next = new Int32Array(width);
    const seen = new Uint8Array(width);
    let activeLen = 1;
    active[0] = startX;
    let splits = 0;

    while (currentRow < height - 1 && activeLen > 0) {
        seen.fill(0);
        let nextLen = 0;

        for (let i = 0; i < activeLen; i++) {
            const x = active[i];
            const belowIdx = (currentRow + 1) * width + x;

            if (splitters[belowIdx]) {
                splits++;

                const left = x - 1;
                if (left >= 0 && seen[left] === 0) {
                    seen[left] = 1;
                    next[nextLen++] = left;
                }

                const right = x + 1;
                if (right < width && seen[right] === 0) {
                    seen[right] = 1;
                    next[nextLen++] = right;
                }
            } else {
                if (seen[x] === 0) {
                    seen[x] = 1;
                    next[nextLen++] = x;
                }
            }
        }

        const swapCols = active;
        active = next;
        next = swapCols;
        activeLen = nextLen;
        currentRow++;
    }

    return splits;
}

function part2(state) {
    const { splitters, width, height, startX, startY } = state;
    let currentRow = startY + 1;
    if (currentRow >= height) return 0n;

    let activeCols = new Int32Array(width);
    let nextCols = new Int32Array(width);
    let activeCounts = new Array(width);
    let nextCounts = new Array(width);
    let totals = new Array(width).fill(0n);

    let activeLen = 1;
    activeCols[0] = startX;
    activeCounts[0] = 1n;

    while (currentRow < height - 1 && activeLen > 0) {
        for (let i = 0; i < width; i++) totals[i] = 0n;
        let nextLen = 0;

        for (let i = 0; i < activeLen; i++) {
            const x = activeCols[i];
            const ways = activeCounts[i];
            const belowIdx = (currentRow + 1) * width + x;

            if (splitters[belowIdx]) {
                const left = x - 1;
                if (left >= 0) {
                    if (totals[left] === 0n) nextCols[nextLen++] = left;
                    totals[left] += ways;
                }

                const right = x + 1;
                if (right < width) {
                    if (totals[right] === 0n) nextCols[nextLen++] = right;
                    totals[right] += ways;
                }
            } else {
                if (totals[x] === 0n) nextCols[nextLen++] = x;
                totals[x] += ways;
            }
        }

        for (let i = 0; i < nextLen; i++) {
            const x = nextCols[i];
            nextCounts[i] = totals[x];
        }

        let swapCols = activeCols;
        activeCols = nextCols;
        nextCols = swapCols;

        let swapCounts = activeCounts;
        activeCounts = nextCounts;
        nextCounts = swapCounts;

        activeLen = nextLen;
        currentRow++;
    }

    let total = 0n;
    for (let i = 0; i < activeLen; i++) {
        total += activeCounts[i];
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
