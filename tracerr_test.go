package tracerr

import (
	"errors"
	"runtime"
	"strings"
	"testing"
)

func TestWrap(t *testing.T) {
	err := Wrap(errors.New("Bogus Error"))

	msg := err.Error()
	if !strings.HasPrefix(msg, "Bogus Error\n") {
		t.Errorf("Unexpected error message: %s", msg)
	}

	lineCount := strings.Count(err.Error(), "\n")
	stackDepth := currentStackDepth()

	if lineCount != stackDepth {
		t.Errorf("Unexpected number of lines in error message. Expected %d. Got %d.", stackDepth, lineCount)
	}
}

func TestError(t *testing.T) {
	err := Error("Bogus Error")

	msg := err.Error()
	if !strings.HasPrefix(msg, "Bogus Error\n") {
		t.Errorf("Unexpected error message: %s", msg)
	}

	lineCount := strings.Count(err.Error(), "\n")
	stackDepth := currentStackDepth()

	if lineCount != stackDepth {
		t.Errorf("Unexpected number of lines in error message. Expected %d. Got %d.", stackDepth, lineCount)
	}
}

func TestErrorf(t *testing.T) {
	err := Errorf("This is a %s Error", "Bogus")

	msg := err.Error()
	if !strings.HasPrefix(msg, "This is a Bogus Error\n") {
		t.Errorf("Unexpected error message: %s", msg)
	}

	lineCount := strings.Count(err.Error(), "\n")
	stackDepth := currentStackDepth()

	if lineCount != stackDepth {
		t.Errorf("Unexpected number of lines in error message. Expected %d. Got %d.", stackDepth, lineCount)
	}
}

func TestNestedWrap(t *testing.T) {
	errMsg := "Bogus Error"
	err1 := Wrap(errors.New(errMsg))
	err2 := Wrap(err1)

	if err1.Error() != err2.Error() {
		t.Errorf("Expected error messages to be the same.\nExpected:\n%s\nGot:\n%s", err1.Error(), err2.Error())
	}
}

func TestWrapForDeepStack(t *testing.T) {
	err := nestedWrapError(maxStackDepth+10, errors.New("Bogus Error"))
	expectedCount := maxStackDepth
	actualCount := strings.Count(err.Error(), " at ")
	if actualCount != expectedCount {
		t.Errorf("Unexpected number of stack frames in error message. Expected %d. Got %d.", expectedCount, actualCount)
	}
}

func TestParseFuncName(t *testing.T) {
	testParseFuncName(t, "package.function", "package", "function")
	testParseFuncName(t, "package.(*type).function", "package", "(*type).function")
	testParseFuncName(t, "path/to/package.function", "path/to/package", "function")
	testParseFuncName(t, "path/to/package.(*type).function", "path/to/package", "(*type).function")
	testParseFuncName(t, "the.path/to/the.package.function", "the.path/to/the.package", "function")
	testParseFuncName(t, "the.path/to/the.package.(*type).function", "the.path/to/the.package", "(*type).function")
}

func testParseFuncName(t *testing.T, fname, expectedPackagePath, expectedFuncSignature string) {
	packagePath, funcSignature := parseFuncName(fname)

	if packagePath != expectedPackagePath {
		t.Errorf("Unexpected package path. Expected '%s'. Got '%s'.", expectedPackagePath, packagePath)
	}

	if funcSignature != expectedFuncSignature {
		t.Errorf("Unexpected function signature. Expected '%s'. Got '%s'.", expectedFuncSignature, packagePath)
	}
}

func currentStackDepth() int {
	return runtime.Callers(2, make([]uintptr, 32))
}

func nestedWrapError(depth int, err error) error {
	if depth > 1 {
		err = nestedWrapError(depth-1, err)
	}
	return Wrap(err)
}
