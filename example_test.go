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
}

func ExampleError() {
	// Create error using convinience method
	err := tracerr.Error("Example Error")

	// Print error message and stack trace
	fmt.Println(err.Error())
}

func ExampleErrorf() {
	// Create error using convinience method
	err := tracerr.Errorf("%s %s", "Example", "Error")

	// Print error message and stack trace
	fmt.Println(err.Error())
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
