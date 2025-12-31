# 01 - Basics

This directory contains examples covering the fundamental building blocks of the Go programming language. Each subdirectory focuses on a specific core concept.

## Contents

### [Variables](./variables/)

Demonstrates how to declare and initialize variables in Go.

- Standard declaration (`var`)
- Short variable declaration (`:=`)
- Multiple variable declaration
- Basic types usage

### [Data Types](./data-types/)

Explores Go's basic data types and their zero values.

- **Integers**: `int`, `int64`
- **Floats**: `float32`, `float64`
- **Strings**: `string`
- **Booleans**: `bool`
- **Zero Values**: The default values assigned to variables declared without an initial value.

### [Conditionals](./conditionals/)

Shows how to control the flow of execution based on conditions.

- `if`, `else if`, `else` statements
- `if` with short statement (initialization)
- `switch` statements (including cases and default)

### [Loops](./loops/)

Illustrates the various ways to use the `for` loop, which is Go's only looping construct.

- Standard count-controlled loop
- "While"-style loop (condition only)
- Infinite loop with `break`
- Iterating over collections with `range`

## How to Run

To run the code in any of these directories, navigate to the specific folder and use the `go run` command:

```bash
cd variables
go run main.go
```

Repeat for other directories as needed:

```bash
cd ../loops
go run main.go
```
