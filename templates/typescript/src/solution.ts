export function solve(input: string): number {
  // Example: count non-empty lines
  return input
    .split("\n")
    .filter((line) => line.trim().length > 0).length;
}
