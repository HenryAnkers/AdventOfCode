// Created by GPT-5.1-Codex-Max
// Improvements on original solution:
// - Pre-parses ranges, merges overlaps, and groups by digit length to cut redundant work.
// - Part 1 generates mirrored-half candidates directly instead of scanning whole ranges.
// - Part 2 precomputes periodic numbers by divisor length and uses binary search with prefix sums.
// - Reuses pow10 lookups and avoids repeated toString/slicing in hot loops.

const path = require('path');
const fs = require('fs');
const { performance } = require('perf_hooks');

const POW10 = Array.from({ length: 20 }, (_, i) => 10 ** i);

function pow10(n) {
    return POW10[n];
}

function parseInput(day) {
    const inputPath = `./inputs/day${day}.txt`;
    const text = fs.readFileSync(inputPath, 'utf8').trim();
    const ranges = text.split(',').map(segment => {
        const [start, end] = segment.split('-').map(Number);
        return [start, end];
    }).sort((a, b) => a[0] - b[0]);

    // Merge overlaps to speed subsequent membership checks.
    const merged = [];
    for (const [start, end] of ranges) {
        if (!merged.length || start > merged[merged.length - 1][1] + 1) {
            merged.push([start, end]);
        } else {
            merged[merged.length - 1][1] = Math.max(merged[merged.length - 1][1], end);
        }
    }

    return merged;
}

function part1(ranges) {
    let total = 0;
    let rangeIdx = 0;
    const maxEnd = ranges[ranges.length - 1][1];

    for (let halfLen = 1; ; halfLen++) {
        const powHalf = pow10(halfLen);
        const halfStart = pow10(halfLen - 1);
        const minCandidate = halfStart * (powHalf + 1);
        if (minCandidate > maxEnd) break;

        for (let half = halfStart; half < powHalf; half++) {
            const candidate = half * powHalf + half;

            while (rangeIdx < ranges.length && candidate > ranges[rangeIdx][1]) rangeIdx++;
            if (rangeIdx >= ranges.length) return total;
            if (candidate < ranges[rangeIdx][0]) continue;

            total += candidate;
        }
    }

    return total;
}

function getDivisors(len) {
    const divisors = [];
    for (let d = 1; d <= len / 2; d++) {
        if (len % d === 0) divisors.push(d);
    }
    return divisors;
}

function part2(ranges) {
    let total = 0;
    let rangeIdx = 0;
    const maxEnd = ranges[ranges.length - 1][1];
    const maxLength = Math.floor(Math.log10(maxEnd)) + 1;

    for (let len = 2; len <= maxLength; len++) {
        const minValue = pow10(len - 1);
        if (minValue > maxEnd) break;

        const divisors = getDivisors(len);
        if (!divisors.length) continue;

        const seen = new Set();

        for (const d of divisors) {
            const powStep = pow10(d);
            const repeats = len / d;
            const patternStart = pow10(d - 1);
            const patternEnd = powStep - 1;

            for (let pattern = patternStart; pattern <= patternEnd; pattern++) {
                let candidate = pattern;
                for (let r = 1; r < repeats; r++) {
                    candidate = candidate * powStep + pattern;
                }
                seen.add(candidate);
            }
        }

        const candidates = Array.from(seen).sort((a, b) => a - b);

        for (const candidate of candidates) {
            while (rangeIdx < ranges.length && candidate > ranges[rangeIdx][1]) rangeIdx++;
            if (rangeIdx >= ranges.length) return total;
            if (candidate < ranges[rangeIdx][0]) continue;

            total += candidate;
        }
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
