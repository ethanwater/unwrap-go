package unwrap

import (
	"fmt"
	"unsafe"
)

func Unwrap[T any](value T, err error) T {
	if err != nil {
		panic(fmt.Errorf("failed to unwrap: %v", err))
	}
	return value
}

func Unchecked[T any](output T, _ error) T {
	return *(*T)(unsafe.Pointer(&output))
}

type Result[T any] struct {
	Value T
	Err   error
}

func Raw[T any](value T, err error) Result[T] {
	return Result[T]{Value: value, Err: err}
}

func (r Result[T]) Expect(expect string) T {
	if r.Err != nil {
		panic(fmt.Errorf("%s : %v", expect, r.Err))
	}
	return r.Value
}

func (r Result[T]) Or(defaultValue T) T {
	if r.Err != nil {
		return defaultValue
	}
	return r.Value
}

func (r Result[T]) OrElse(handleErrorFunc func(T) T) T {
	if r.Err != nil {
		return handleErrorFunc(r.Value)
	}
	return r.Value
}
