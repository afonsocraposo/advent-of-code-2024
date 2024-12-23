# Advent of Code 2024 🎄✨
```
     Merry Christmas!
 ------------------------
   \
    \
     \   ,_---~~~~~----._         
  _,,_,*^____      _____``*g*\"*, 
 / __/ /'     ^.  /      \ ^@q   f 
[  @f | @))    |  | @))   l  0 _/  
 \`/   \~____ / __ \_____/    \   
  |           _l__l_           I   
  }          [______]           I  
  ]            | | |            |  
  ]             ~ ~             |  
  |                            |   
   |                           |   
```

Welcome to **Advent of Code 2024**! 🎅✨ This repo features my solutions to the daily programming puzzles, written in **Golang**, with a clean architecture for festive coding!

---

## Why Go? 🐹

- **Coolness**: I like Go and don't get the chance to use it on my day-to-day
- **Speed**: No waiting around for Santa's sleigh.
- **Simplicity**: Go keeps the solutions straightforward and delightful.
- **Concurrency**: Perfect for parallelizing Santa's busy workshop tasks.

---

## Repo Structure 📂

Here's how the magic is organized:

```
.
├── LICENSE               # Licensing information
├── cmd
│   └── main.go           # Entry point for running solutions
├── go.mod                # Go module configuration
├── internal
│   ├── day1
│   │   └── day1.go       # Solution logic for Day 1
│   ├── day2
│   │   └── day2.go       # Solution logic for Day 2
│   └── utils       # Some utils to help
│       ├── filereader
│       │   └── filereader.go  # File parsing utilities
│       ├── math
│       │   └── math.go        # Math helpers
│       └── matrix
│           └── matrix.go      # Matrix manipulation helpers
└── puzzles                # Puzzle input files (ignored in Git)
    └── day2
        ├── example1.txt
        └── input1.txt
```

---

## Running a Solution 🏃

To solve the puzzles for a specific day, run the following command:

```bash
go run cmd/main.go <day-number>
```

### Example
```bash
go run cmd/main.go 2
```

### Output
```plaintext
Advent of Code 2024
DAY 2
Part 1:
The solution is: 218
Part 2:
The solution is: 290
```

---

## Adding New Days 🗓️

1. Create a new directory under `internal` for the day:
   ```
   mkdir -p internal/day3
   ```

2. Write your solution in `internal/day3/day3.go`.

3. Update `cmd/main.go` to include a case for the new day.
