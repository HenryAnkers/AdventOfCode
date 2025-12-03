// Helper to compare original vs LLM solutions for all dayN.js files (excluding dayX).
// Prints per-day timings using in-program parse/part1/part2 measurements.
const { execFileSync } = require('child_process');
const fs = require('fs');
const path = require('path');

const ROOT = __dirname;
const LLM_DIR = path.join(ROOT, 'LLM');

function listDays() {
    return fs.readdirSync(ROOT)
        .filter(name => /^day(\d+)\.js$/.test(name) && name !== 'dayX.js')
        .map(name => name.match(/^day(\d+)\.js$/)[1])
        .sort((a, b) => Number(a) - Number(b));
}

function run(scriptPath) {
    const output = execFileSync('node', [scriptPath], { cwd: ROOT, encoding: 'utf8' });
    return output.trim();
}

function parseTimings(output) {
    const resultMatch = /Part 1:\s*([0-9]+).*Part 2:\s*([0-9]+)/s.exec(output);
    const timeMatch = [
        /Parsing took ([0-9.]+) ms/,
        /Part 1: [0-9]+ \(took ([0-9.]+) ms\)/,
        /Part 2: [0-9]+ \(took ([0-9.]+) ms\)/,
    ].map(r => {
        const m = output.match(r);
        return m ? Number(m[1]) : null;
    });

    if (!resultMatch || timeMatch.some(t => t === null)) return null;
    return {
        results: [resultMatch[1], resultMatch[2]],
        parts: timeMatch,
        totalMs: timeMatch.reduce((a, b) => a + b, 0),
    };
}

function formatLine(label, parsed) {
    const [parseMs, p1Ms, p2Ms] = parsed.parts;
    return {
        label,
        total: parsed.totalMs.toFixed(3),
        parts: [parseMs, p1Ms, p2Ms].map(v => v.toFixed(3)),
    };
}

function compareDay(day) {
    const origOutput = run(`day${day}.js`);
    const llmPath = path.join('LLM', `day${day}.js`);
    const llmExists = fs.existsSync(path.join(LLM_DIR, `day${day}.js`));
    const llmOutput = llmExists ? run(llmPath) : null;

    const origParsed = parseTimings(origOutput);
    const llmParsed = llmOutput ? parseTimings(llmOutput) : null;

    const match = origParsed && llmParsed &&
        origParsed.results[0] === llmParsed.results[0] &&
        origParsed.results[1] === llmParsed.results[1];

    const origLine = formatLine('Original', origParsed);
    const llmLine = llmParsed ? formatLine('LLM', llmParsed) : null;

    const labelWidth = 8;
    const totalWidth = Math.max(
        origLine.total.length,
        llmLine ? llmLine.total.length : 0
    );

    const prefix = `DAY ${day}: `;
    const spacer = ' '.repeat(prefix.length);
    const fmt = (line) =>
        `${line.label.padEnd(labelWidth)} ${line.total.padStart(totalWidth)} ms [${line.parts.join(' + ')}]`;

    console.log(prefix + fmt(origLine));
    if (llmParsed) {
        console.log(`${spacer}${fmt(llmLine)}${match ? '' : '  (mismatch!)'}`);
    } else {
        console.log('         LLM missing');
    }

    return {
        origTotal: origLine.total,
        llmTotal: llmParsed ? llmLine.total : null,
    };
}

function main() {
    const days = listDays();
    let origFaster = 0;
    let llmFaster = 0;

    days.forEach(day => {
        const result = compareDay(day);
        if (result && result.llmTotal !== null) {
            const o = Number(result.origTotal);
            const l = Number(result.llmTotal);
            if (o < l) origFaster++;
            else if (l < o) llmFaster++;
        }
        if (day !== days[days.length - 1]) console.log('');
    });

    console.log('');
    console.log(`Totals: Original faster on ${origFaster} day(s), LLM faster on ${llmFaster} day(s).`);
}

main();
