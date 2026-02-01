from __future__ import annotations


def solve(text: str) -> int:
    # Example: count non-empty lines
    return sum(1 for line in text.splitlines() if line.strip())

