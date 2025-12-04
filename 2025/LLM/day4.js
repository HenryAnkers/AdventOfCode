// Created by GPT-5.1-Codex-Max
// Improvements on original solution:
// - Parses into a flat Uint8Array grid (1 for '@', 0 otherwise) to shrink memory and simplify index math.
// - Builds neighbor counts via additive contribution (8 writes per '@') to avoid repeated bounds checks in hot paths.
// - Part 1 reads those counts directly for a single O(n) exposed-cell tally.
// - Part 2 computes the k-core with a queue driven by neighbor-count decrements, eliminating full-grid rescans per round.

const path = require('path');
const fs = require('fs');
const { performance } = require('perf_hooks');

function parseInput(day) {
    const text = fs.readFileSync(path.join(__dirname, '..', `inputs/day${day}.txt`), 'utf8').trimEnd();
    const lines = text.split('\n');
    const height = lines.length;
    const width = lines[0].length;
    const flat = new Uint8Array(height * width);

    for (let y = 0; y < height; y++) {
        const line = lines[y];
        const rowOffset = y * width;
        for (let x = 0; x < width; x++) {
            flat[rowOffset + x] = line.charCodeAt(x) === 64 ? 1 : 0; // '@' => 1
        }
    }

    return { flat, width, height };
}

function buildNeighborCounts(state) {
    const { flat, width, height } = state;
    const counts = new Int8Array(flat.length);

    for (let y = 0; y < height; y++) {
        const rowOffset = y * width;
        const yMin = y > 0;
        const yMax = y + 1 < height;

        for (let x = 0; x < width; x++) {
            const idx = rowOffset + x;
            if (flat[idx] === 0) continue;

            const xMin = x > 0;
            const xMax = x + 1 < width;

            // Manually unrolled neighbors to minimize branch work inside inner loops.
            if (yMin) {
                const nRow = rowOffset - width;
                if (xMin) counts[nRow + x - 1]++;
                counts[nRow + x]++;
                if (xMax) counts[nRow + x + 1]++;
            }

            if (xMin) counts[idx - 1]++;
            if (xMax) counts[idx + 1]++;

            if (yMax) {
                const nRow = rowOffset + width;
                if (xMin) counts[nRow + x - 1]++;
                counts[nRow + x]++;
                if (xMax) counts[nRow + x + 1]++;
            }
        }
    }

    return counts;
}

function part1(state) {
    const counts = buildNeighborCounts(state);
    const { flat } = state;
    let exposed = 0;

    for (let i = 0; i < flat.length; i++) {
        if (flat[i] && counts[i] < 4) exposed++;
    }

    return exposed;
}

function part2(state) {
    const { flat, width, height } = state;
    const counts = buildNeighborCounts(state);
    const queue = new Int32Array(flat.length);
    let head = 0;
    let tail = 0;

    for (let i = 0; i < flat.length; i++) {
        if (flat[i] && counts[i] < 4) queue[tail++] = i;
    }

    let removed = 0;

    while (head < tail) {
        const idx = queue[head++];
        if (flat[idx] === 0) continue;

        flat[idx] = 0;
        removed++;

        const y = Math.floor(idx / width);
        const x = idx - y * width;
        const yMin = y > 0;
        const yMax = y + 1 < height;
        const xMin = x > 0;
        const xMax = x + 1 < width;

        const tryEnqueue = (nIdx) => {
            const nextCount = --counts[nIdx];
            if (flat[nIdx] && nextCount === 3) queue[tail++] = nIdx;
        };

        if (yMin) {
            const nRow = idx - width;
            if (xMin) tryEnqueue(nRow - 1);
            tryEnqueue(nRow);
            if (xMax) tryEnqueue(nRow + 1);
        }

        if (xMin) tryEnqueue(idx - 1);
        if (xMax) tryEnqueue(idx + 1);

        if (yMax) {
            const nRow = idx + width;
            if (xMin) tryEnqueue(nRow - 1);
            tryEnqueue(nRow);
            if (xMax) tryEnqueue(nRow + 1);
        }
    }

    return removed;
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
