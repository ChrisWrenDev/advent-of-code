from __future__ import annotations

from aoc_day.solution import solve


def test_example() -> None:
    text = "a\nb\nc\n"
    assert solve(text) == 3

