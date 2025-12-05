// Created by GPT-5.1-Codex-Max
// Improvements on original solution:
// - Parses blocks into plain Number arrays to preserve wide IDs (input exceeds 32-bit).
// - Sorts ranges once and merges via a linear sweep (overlap-only, no O(n^2) nested checks).
// - Part 1 sorts ingredients and uses a two-pointer scan against merged ranges for O(n log n) prep + O(n + m) counting.
// - Part 2 reuses the merged ranges directly to sum coverage without rescanning inputs.

const path = require('path');
const fs = require('fs');
const { performance } = require('perf_hooks');

function parseInput(day) {
    const text = fs.readFileSync(path.join(__dirname, '..', `inputs/day${day}.txt`), 'utf8').trim();
    const [rangeBlock, ingredientBlock = ''] = text.split('\n\n');

    const rangeLines = rangeBlock.split('\n');
    const ranges = new Array(rangeLines.length * 2);
    for (let i = 0, w = 0; i < rangeLines.length; i++, w += 2) {
        const line = rangeLines[i];
        const dash = line.indexOf('-');
        ranges[w] = Number(line.slice(0, dash));
        ranges[w + 1] = Number(line.slice(dash + 1));
    }

    const ingredientLines = ingredientBlock ? ingredientBlock.split('\n') : [];
    const ingredients = new Array(ingredientLines.length);
    for (let i = 0; i < ingredientLines.length; i++) {
        ingredients[i] = Number(ingredientLines[i]);
    }

    return { ranges, ingredients };
}

function mergeRanges(rangePairs) {
    const pairCount = rangePairs.length >> 1;
    if (pairCount === 0) {
        return { starts: [], ends: [], count: 0 };
    }

    const pairs = new Array(pairCount);
    for (let i = 0, w = 0; i < pairCount; i++, w += 2) {
        pairs[i] = [rangePairs[w], rangePairs[w + 1]];
    }
    pairs.sort((a, b) => (a[0] - b[0]) || (a[1] - b[1]));

    const mergedStarts = [];
    const mergedEnds = [];

    let [curStart, curEnd] = pairs[0];
    for (let i = 1; i < pairs.length; i++) {
        const [start, end] = pairs[i];
        if (start <= curEnd) {
            if (end > curEnd) curEnd = end;
        } else {
            mergedStarts.push(curStart);
            mergedEnds.push(curEnd);
            curStart = start;
            curEnd = end;
        }
    }
    mergedStarts.push(curStart);
    mergedEnds.push(curEnd);

    return {
        starts: mergedStarts,
        ends: mergedEnds,
        count: mergedStarts.length,
    };
}

function part1(ingredients, merged) {
    if (merged.count === 0 || ingredients.length === 0) return 0;

    const sortedIngredients = ingredients.slice().sort((a, b) => a - b);

    let ingIdx = 0;
    let rangeIdx = 0;
    let count = 0;

    while (ingIdx < sortedIngredients.length && rangeIdx < merged.count) {
        const value = sortedIngredients[ingIdx];
        const start = merged.starts[rangeIdx];
        const end = merged.ends[rangeIdx];

        if (value < start) {
            ingIdx++;
        } else if (value > end) {
            rangeIdx++;
        } else {
            count++;
            ingIdx++;
        }
    }

    return count;
}

function part2(merged) {
    let total = 0;
    for (let i = 0; i < merged.count; i++) {
        total += merged.ends[i] - merged.starts[i] + 1;
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
const mergedRanges = mergeRanges(inputData.ranges);
const solution1 = part1(inputData.ingredients, mergedRanges);
end = performance.now();
console.log(`Day ${day} - Part 1: ${solution1} (took ${(end - start).toFixed(6)} ms)`);

start = performance.now();
const solution2 = part2(mergedRanges);
end = performance.now();
console.log(`Day ${day} - Part 2: ${solution2} (took ${(end - start).toFixed(6)} ms)`);
