# Convex Hull algorithm (Divide and conquer)
The repository contains the implementation of the Convex Hull algorithm (`convex_hull.go`), unit tests (`convex_hull_test.go`), a small driver program (`demo.go`) and some test files.

## How to use it
Use `go build` to build the program. Then you can use it to calculate the convex hull of a set of points, by providing them in a JSON file (for the format, see the test files from the repo). The result will be printed in the same JSON format to the specified output.

Execute the demo program with the following command:

`convex-hull-go input.json output.json`

If you just want to test the algorithm, use the unit tests by calling `go test`.
