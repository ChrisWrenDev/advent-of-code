from __future__ import annotations

from pathlib import Path

from aoc_day.solution import solve


def main() -> None:
    input_path = Path(__file__).resolve().parents[2] / "input.txt"
    text = input_path.read_text(encoding="utf-8")
    result = solve(text)
    print(result)


if __name__ == "__main__":
    main()

