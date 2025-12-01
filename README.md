# Advent of Code - Go

A reusable Advent of Code repository structure in Go, using [mise](https://mise.jdx.dev/) for task running and toolchain management.

## Structure

```
.
├── cmd/aoc/          # CLI entry point
├── internal/runner/  # Solution runner and registry
├── pkg/utils/        # Shared utilities (input parsing, math, grid)
├── years/
│   ├── register.go   # Import all solution packages
│   ├── 2025/
│   │   ├── day01/
│   │   │   ├── solution.go
│   │   │   ├── solution_test.go
│   │   │   └── input.txt
│   │   └── day02/
│   └── 2025/
└── mise.toml
```

## Prerequisites

Install [mise](https://mise.jdx.dev/getting-started.html):

```bash
curl https://mise.run | sh
```

Then trust and install tools:

```bash
mise trust
mise install
```

## Quick Start

```bash
# Build the CLI
mise run build

# Run a specific day
mise run solve 2025 1

# Run all days for a year
mise run solve 2025

# Run tests
mise run test

# Create a new day from template
mise run new-day 2025 2

# List all available tasks
mise tasks
```

## CLI Usage

```bash
# Run today's puzzle (auto-detects during December)
./bin/aoc

# Run specific day
./bin/aoc -year=2025 -day=1

# Run only part 1
./bin/aoc -year=2025 -day=1 -part=1

# Run all days for a year
./bin/aoc -year=2025 -day=0

# Use custom input file
./bin/aoc -year=2025 -day=1 -input=custom.txt
```

## Adding a New Day

1. Create the day directory:

    ```bash
    mise run new-day 2025 2
    ```

2. Add the import to `years/register.go`:

    ```go
    _ "aoc/years/2025/day02"
    ```

3. Add your input to `years/2025/day02/input.txt`

4. Implement your solution in `solution.go`

## Solution Template

Each solution implements the `runner.Solution` interface:

```go
package day02

import (
    "aoc/internal/runner"
    "aoc/pkg/utils"
)

func init() {
    runner.Register(2025, 2, &Solution{})
}

type Solution struct{}

func (s *Solution) Part1(input string) (string, error) {
    lines := utils.Lines(input)
    // Your solution here
    return "answer", nil
}

func (s *Solution) Part2(input string) (string, error) {
    // Your solution here
    return "answer", nil
}
```

## Utilities

The `pkg/utils` package includes common helpers:

**Input parsing:**

-   `Lines(input)` - Split into lines
-   `Ints(input)` - Parse all integers
-   `Grid(input)` - Parse into 2D byte grid
-   `CSV(input)` - Parse comma-separated values

**Math:**

-   `Abs`, `Min`, `Max`, `Sum`
-   `GCD`, `LCM`
-   `Sign`, `Clamp`

**Grid/Points:**

-   `Point{X, Y}` with `Add`, `Sub`, `Manhattan`
-   `Neighbors4()`, `Neighbors8()`
-   Direction constants: `Up`, `Down`, `Left`, `Right`
-   `InBounds()`, `GridAt()`

## Downloading Input

Set your session cookie and use:

```bash
export AOC_SESSION="your-session-cookie"
mise run download 2025 1
```

To get your session cookie, log into adventofcode.com and copy the `session` cookie value from your browser's developer tools.

## Available Tasks

| Task                | Description                                         |
| ------------------- | --------------------------------------------------- |
| `mise run build`    | Build the CLI                                       |
| `mise run solve`    | Run solution(s): `mise run solve <year> [day]`      |
| `mise run test`     | Run all tests                                       |
| `mise run test-day` | Run tests for day: `mise run test-day <year> <day>` |
| `mise run bench`    | Run benchmarks                                      |
| `mise run new-day`  | Create new day: `mise run new-day <year> <day>`     |
| `mise run download` | Download input: `mise run download <year> <day>`    |
| `mise run clean`    | Clean build artifacts                               |
