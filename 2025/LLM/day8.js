// Created by GPT-5.1-Codex-Max
// Improvements on original solution:
// - Parses into a compact Int32Array and precomputes squared distances once; no Math.sqrt calls.
// - Uses a union-find for near-constant-time circuit merges instead of ad-hoc Set unioning.
// - Reuses the sorted edge list for both parts and stops immediately once the final merge happens.
// - Keeps allocations bounded (single edge array, single DSU per part) for faster runtimes.

const path = require('path');
const fs = require('fs');
const { performance } = require('perf_hooks');

function parseInput(day) {
    const text = fs.readFileSync(path.join(__dirname, '..', `inputs/day${day}.txt`), 'utf8').trimEnd();
    const lines = text.split('\n');
    const count = lines.length;
    const coords = new Int32Array(count * 3);

    for (let i = 0; i < count; i++) {
        const [x, y, z] = lines[i].split(',').map(Number);
        const base = i * 3;
        coords[base] = x;
        coords[base + 1] = y;
        coords[base + 2] = z;
    }

    return { coords, count };
}

function buildEdges(coords, count) {
    const total = (count * (count - 1)) / 2;
    const edges = new Array(total);
    let idx = 0;

    for (let i = 0; i < count; i++) {
        const baseI = i * 3;
        const xi = coords[baseI];
        const yi = coords[baseI + 1];
        const zi = coords[baseI + 2];

        for (let j = i + 1; j < count; j++) {
            const baseJ = j * 3;
            const dx = xi - coords[baseJ];
            const dy = yi - coords[baseJ + 1];
            const dz = zi - coords[baseJ + 2];
            const dist2 = dx * dx + dy * dy + dz * dz;
            edges[idx++] = [dist2, i, j];
        }
    }

    edges.sort((a, b) => a[0] - b[0]);
    return edges;
}

class DSU {
    constructor(n) {
        this.parent = new Int32Array(n);
        this.size = new Int32Array(n);
        for (let i = 0; i < n; i++) {
            this.parent[i] = i;
            this.size[i] = 1;
        }
    }

    find(x) {
        let p = x;
        while (this.parent[p] !== p) p = this.parent[p];
        while (this.parent[x] !== x) {
            const next = this.parent[x];
            this.parent[x] = p;
            x = next;
        }
        return p;
    }

    union(a, b) {
        let ra = this.find(a);
        let rb = this.find(b);
        if (ra === rb) return false;

        if (this.size[ra] < this.size[rb]) {
            const tmp = ra;
            ra = rb;
            rb = tmp;
        }

        this.parent[rb] = ra;
        this.size[ra] += this.size[rb];
        return true;
    }
}

function topThreeProduct(dsu, count) {
    const seen = new Set();
    const sizes = [];

    for (let i = 0; i < count; i++) {
        const root = dsu.find(i);
        if (seen.has(root)) continue;
        seen.add(root);
        sizes.push(dsu.size[root]);
    }

    sizes.sort((a, b) => b - a);
    if (sizes.length < 3) return 0;
    return sizes[0] * sizes[1] * sizes[2];
}

function part1(data) {
    const { edges, count } = data;
    const dsu = new DSU(count);
    const limit = Math.min(1000, edges.length);

    for (let i = 0; i < limit; i++) {
        const edge = edges[i];
        const a = edge[1];
        const b = edge[2];
        dsu.union(a, b);
    }

    return topThreeProduct(dsu, count);
}

function part2(data) {
    const { coords, edges, count } = data;
    const dsu = new DSU(count);
    let components = count;

    for (let i = 0; i < edges.length; i++) {
        const edge = edges[i];
        const a = edge[1];
        const b = edge[2];
        if (dsu.union(a, b)) {
            components--;
            if (components === 1) {
                const ax = coords[a * 3];
                const bx = coords[b * 3];
                return ax * bx;
            }
        }
    }

    throw new Error('Failed to connect all junction boxes');
}

const filename = path.basename(__filename);
const day = filename.match(/^day(\d+)/)[1];

let start = performance.now();
const parsed = parseInput(day);
const edges = buildEdges(parsed.coords, parsed.count);
const inputData = { ...parsed, edges };
let end = performance.now();
console.log(`Day ${day} - Parsing + edge build took ${(end - start).toFixed(6)} ms`);

start = performance.now();
const solution1 = part1(inputData);
end = performance.now();
console.log(`Day ${day} - Part 1: ${solution1} (took ${(end - start).toFixed(6)} ms)`);

start = performance.now();
const solution2 = part2(inputData);
end = performance.now();
console.log(`Day ${day} - Part 2: ${solution2} (took ${(end - start).toFixed(6)} ms)`);
