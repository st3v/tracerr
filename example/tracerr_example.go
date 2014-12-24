package main

import (
	"errors"
	"fmt"

	"github.com/st3v/tracerr"
)

func main() {
	foo := &foo{}

	err := nested(4, func() error {
		return foo.bar()
	})

	if err != nil {
		fmt.Println(err.Error())
	}
}

func nested(depth int, fn func() error) error {
	if depth <= 1 {
		return fn()
	}
	return tracerr.Wrap(nested(depth-1, fn))
}

type foo struct{}

func (f *foo) bar() error {
	return tracerr.Wrap(errors.New("FooBarError"))
}
