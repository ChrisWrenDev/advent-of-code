# Advent of Code

Welcome to my **Advent of Code** solutions!

[Advent of Code](https://adventofcode.com/) is an annual coding challenge event that runs every December, offering a new programming puzzle each day leading up to Christmas. The puzzles are designed to be fun, challenging, and a great way to sharpen your coding skills. Participants from all around the world compete to solve these puzzles as quickly as possible.

This repository is a collection of my solutions to the Advent of Code puzzles. Each folder contains my approach and code for the corresponding day's challenge.

## Progress

```shell
2015 [-------------------------] 00/25
2016 [-------------------------] 00/25
2017 [-------------------------] 00/25
2018 [-------------------------] 00/25
2019 [-------------------------] 00/25
2020 [-------------------------] 00/25
2021 [-------------------------] 00/25
2022 [-------------------------] 00/25
2023 [-------------------------] 00/25
2024 [-------------------------] 00/25
2025 [-------------------------] 00/12
```

## AOC

This repository includes a single Bash entrypoint, ./aoc, which is responsible for:

- scaffolding new Advent of Code days
- running solutions (optionally across multiple languages)
- delegating to each language’s native test runner

The script does not replace language toolchains (Cargo, Go, npm, etc.). It simply orchestrates them so the repo can be driven with one consistent interface.

```shell
# creating a new day
./aoc new YEAR DAY

# running solutions
./aoc run YEAR DAY

# running tests
./aoc test

# limit to specific languages
./aoc new YEAR DAY --lang rust,go

# overwrite exisitng files
./aoc new YEAR DAY --force
```

Make sure the script is executable:

```shell
chmod +x aoc
```

## Typical Workflow

```shell
# scaffold a new day
./aoc new 2023 5

# paste input into 2015/day01/input.txt

# work on solutions
./aoc run 2015 1 --lang python

# compare implementations
./aoc run 2015 1 --lang python,rust,go

# make sure nothing broke
./aoc test
```
