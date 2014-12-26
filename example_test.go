package tracerr_test

import (
	"errors"
	"fmt"

	"github.com/st3v/tracerr"
)

func ExampleWrap() {
	// A regular error
	err := errors.New("Example Error")

	// Wrap it with the current stack trace
	err = tracerr.Wrap(err)

	// Print error message and stack trace
	fmt.Println(err.Error())
	// Output:
	// Example Error
	//   at ExampleWrap (github.com/st3v/tracerr_test/example_test.go:15)
	//   at runExample (testing/example.go:99)
	//   at RunExamples (testing/example.go:36)
	//   at Main (testing/testing.go:436)
	//   at main (main/_testmain.go:59)
	//   at main (runtime/proc.c:249)
	//   at goexit (runtime/proc.c:1445)
}

func ExampleWrap_nested() {
	// Create some foo
	foo := &foo{}

	// Recursive function call with foo.bar() at its end
	err := nested(4, func() error {
		return foo.bar()
	})

	// Print error
	if err != nil {
		fmt.Println(err.Error())
	}
	// Output:
	// Example Error
	//   at (*foo).bar (github.com/st3v/tracerr_test/example_test.go:72)
	//   at func·001 (github.com/st3v/tracerr_test/example_test.go:36)
	//   at nested (github.com/st3v/tracerr_test/example_test.go:63)
	//   at nested (github.com/st3v/tracerr_test/example_test.go:65)
	//   at nested (github.com/st3v/tracerr_test/example_test.go:65)
	//   at nested (github.com/st3v/tracerr_test/example_test.go:65)
	//   at ExampleWrap_nested (github.com/st3v/tracerr_test/example_test.go:37)
	//   at runExample (testing/example.go:99)
	//   at RunExamples (testing/example.go:36)
	//   at Main (testing/testing.go:436)
	//   at main (main/_testmain.go:59)
	//   at main (runtime/proc.c:249)
	//   at goexit (runtime/proc.c:1445)
}

// Recursive function that calls the passed in function at its tail end.
func nested(depth int, fn func() error) error {
	if depth <= 1 {
		return fn()
	}
	return nested(depth-1, fn)
}

type foo struct{}

// Returns an error that has been wrapped with the current stack trace.
func (f *foo) bar() error {
	return tracerr.Wrap(errors.New("Example Error"))
}
