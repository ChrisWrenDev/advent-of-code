import { describe, expect, it } from "vitest";
import { solve } from "../src/solution.js";

describe("solve", () => {
  it("counts non-empty lines", () => {
    const input = "a\nb\nc\n";
    expect(solve(input)).toBe(3);
  });
});

