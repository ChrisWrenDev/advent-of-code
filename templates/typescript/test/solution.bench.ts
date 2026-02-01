import { bench, describe } from "vitest";
import { solve } from "../src/solution.js";

const input = Array.from({ length: 10_000 }, (_, i) => `line${i}`).join("\n");

describe("solve benchmark", () => {
  bench("solve", () => {
    solve(input);
  });
});

