# advent-of-code-2015
Advent of code 2015 https://adventofcode.com/2015

# Rust Developer - Add New Day
1. `mkdir day-N/rust/`
2. `cargo init`

# Rust Developer - How to Run a Day
1. mkdir input
2. Add input txt files
- sample.txt: sample input
- input.txt: test input
3. Run `cargo run`

# Go Developer - Setup
Project Structure
advent-of-code-2015/
├── MODULE.bazel                            # Root BUILD file (optional)
├── go_libs/
│   └── file_processor/
│       ├── BUILD.bazel                    # Defines the go_library target
│       ├── file_processor.go              # Main package file
│       └── file_processor_test.go         # Tests (optional)
└── day-1/
    └── go/
        ├── BUILD.bazel                    # Defines the go_binary target
        └── main.go                        # Your solution

# Go Developer - Add New

Step-by-Step Setup
1. WORKSPACE File (Root Level)
- Define your workspace name
- Load the Go rules for Bazel (rules_go)
- Load Gazelle (a tool that auto-generates BUILD files)
- Register the Go SDK
- Initialize the Go module with a module name (e.g., github.com/yourname/advent-of-code-2015)

2. go_libs/file_processor/BUILD.bazel
This defines shared library as a go_library target:
- Name it something like file_processor
- Specify the srcs (your .go files)
- Set importpath (e.g., github.com/yourname/advent-of-code-2015/go_libs/file_processor)
- Set visibility to ["//visibility:public"] so other packages can use it

Your solution file that:
- Imports the file_processor package using the full import path
- Uses functions like file_processor.LoadRaw("input.txt")

Key Bazel Concepts
Targets: Each BUILD.bazel file defines targets (like go_library, go_binary, go_test)
Labels: You reference other packages using labels like //go_libs/file_processor (the // means "from workspace root")
Import Paths: Go needs fully-qualified import paths. In Bazel, you set these explicitly in the importpath attribute
Visibility: By default, targets are private. You need to set visibility = ["//visibility:public"] on your library so other packages can depend on it

Building and Running with Bazel

Once set up, you'd:
- Build: bazel build //day-1/go:go (or whatever you name your binary target)
- Run: bazel run //day-1/go:go
- Test: bazel test //go_libs/file_processor:file_processor_test

Gazelle (Optional but Recommended)

Gazelle can auto-generate and update BUILD.bazel files:
- Add a gazelle target in your root BUILD.bazel
- Run bazel run //:gazelle to auto-generate BUILD files
- It will detect your Go files and create appropriate targets
- You can add # gazelle: directives in comments to customize behavior

Advantages of This Setup

1. Hermetic builds: Bazel manages all dependencies
2. Caching: Only rebuilds what changed
3. Monorepo-friendly: Easy to share code across multiple days/languages
4. Consistent: Same build system for Go, Rust (with rules_rust), etc.
