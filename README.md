# Unexpected Type Conversion in Go Comparisons

This repository demonstrates a subtle but potentially problematic aspect of Go's type system: implicit type conversions during comparisons.  While convenient, it can lead to unexpected results if not handled carefully.  The example showcases how an integer `0` will compare equal to a floating-point `0.0`.  This might not always be the desired behaviour, especially when dealing with numerical precision.

The `bug.go` file contains the code exhibiting this behavior.  The `bugSolution.go` file offers potential solutions and best practices to avoid this kind of problem.

## How to Reproduce

1. Clone this repository.
2. Navigate to the project directory.
3. Run `go run bug.go`