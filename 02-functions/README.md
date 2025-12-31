# 02 - Functions

This directory explores how functions work in Go, from basic syntax to more advanced features like multiple return values and defer/panic/recover.

## Contents

### [Functions](./functions/)

Covers basic function declaration and usage.

- Defining functions
- Function parameters
- Return values
- Package-level vs local scope

### [Multiple Returns](./multiple-returns/)

Demonstrates Go's ability to return multiple values from a single function.

- Returning multiple values
- Named return values
- Error handling patterns using multiple returns
- The blank identifier `_` to ignore values

### [Defer, Panic, and Recover](./defer-panic-recover/)

Explains Go's mechanism for resource management and exceptional error handling.

- `defer`: Delaying function execution until the surrounding function returns (useful for cleanup).
- `panic`: Stopping normal execution when an unrecoverable error occurs.
- `recover`: Regaining control of a panicking goroutine.

## How to Run

Navigate to a specific subdirectory and run the `main.go` file:

```bash
cd functions
go run main.go
```
