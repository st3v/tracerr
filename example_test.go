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

func ExampleError() {
	// Create error using convenience method
	err := tracerr.Error("Example Error")

	// Print error message and stack trace
	fmt.Println(err.Error())
}

func ExampleErrorf() {
	// Create error using convenience method
	err := tracerr.Errorf("%s %s", "Example", "Error")

	// Print error message and stack trace
	fmt.Println(err.Error())
}
