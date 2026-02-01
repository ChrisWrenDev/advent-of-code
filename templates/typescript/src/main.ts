import { readFileSync } from "node:fs";
import { resolve } from "node:path";
import { solve } from "./solution.js";

function main(): void {
  const inputPath = resolve(process.cwd(), "../input.txt");
  const input = readFileSync(inputPath, "utf-8");
  const result = solve(input);
  console.log(result);
}

main();

