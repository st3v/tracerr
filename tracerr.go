package tracerr

import (
	"fmt"
	"path/filepath"
	"regexp"
	"runtime"
)

type traceError struct {
	err   error
	stack []*stackFrame
}

const maxStackDepth = 64

func Wrap(err error) error {
	if err == nil {
		return err
	}

	if _, ok := err.(interface {
		stackTrace() []string
	}); ok {
		return err
	}

	return &traceError{
		err:   err,
		stack: trace(2, maxStackDepth),
	}
}

func (t *traceError) Error() string {
	str := t.err.Error()
	for _, frame := range t.stackTrace() {
		str += fmt.Sprintf("\n  at %s", frame)
	}
	return str
}

func (t *traceError) stackTrace() []string {
	stack := make([]string, len(t.stack))
	for i, frame := range t.stack {
		stack[i] = fmt.Sprintf("%s (%s:%d)", frame.function, frame.file, frame.line)
	}
	return stack
}

type stackFrame struct {
	file     string
	line     int
	function string
}

func trace(skip, maxDepth int) []*stackFrame {
	pcs := make([]uintptr, maxDepth)

	count := runtime.Callers(skip+1, pcs)

	frames := make([]*stackFrame, count)

	for i, pc := range pcs[0:count] {
		frames[i] = newStackFrame(pc)
	}

	return frames
}

func newStackFrame(pc uintptr) *stackFrame {
	fn := runtime.FuncForPC(pc)
	file, line := fn.FileLine(pc)
	packagePath, funcSignature := parseFuncName(fn.Name())
	_, fileName := filepath.Split(file)

	return &stackFrame{
		file:     filepath.Join(packagePath, fileName),
		line:     line,
		function: funcSignature,
	}
}

func parseFuncName(name string) (packagePath, signature string) {
	regEx := regexp.MustCompile("([^\\(]*)\\.(.*)")
	parts := regEx.FindStringSubmatch(name)
	packagePath = parts[1]
	signature = parts[2]
	return
}
