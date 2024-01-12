//NOTE: the functions in this library DO NOT work with functions that return multiple
//values. Much like in Rust, It can't directly handle multiple return values. If you have
//a Result or Option that contains a tuple or a struct with multiple values, you'd
//need to destructure or access those values explicitly.

//The descriptions of the following methods are directly taken from the Rust's
//standard library descriptions with slight modifications. You can find the
//original documentation in Rust here:
//Result Doc: https://doc.rust-lang.org/stable/std/result/index.html
//Option Doc: https://doc.rust-lang.org/stable/std/option/index.html

//This library shouldn't be depended on for all error handling in production code.
//Some of the functions use unsafe methods in order to handle its output. In situations
//where errors must be more explicitly handled, please use the default golang error
//handling techniques. However, the functions below will be enough in MOST cases.
//As of right now, only Result's unwrap methods are ported.

package unwrap

import (
	"fmt"
	"reflect"
	"unsafe"
)

type Result[T any] struct {
	Ok  T
	Err error
}

// Checks if the Result type is valid and not an error
func (result *Result[T]) IsOk() bool {
	return result.Err != nil
}

// Checks if the Result is an error
func (result *Result[T]) IsErr() bool {
	return !result.IsOk()
}

// Returns a Result type of functions output. This allows for most unwrap fucntions
// to be called.
func Wrap[T any](value T, err error) Result[T] {
	return Result[T]{Ok: value, Err: err}
}

// Returns the contained Ok value, consuming the self value.
// Because this function may panic, its use is generally discouraged. Instead,
// prefer to use pattern matching and handle the error case explicitly, or call
// unwrap_or, unwrap_or_else, or unwrap_or_default.
//
// *Panics if the value is an Err, with a panic message provided by the Err’s value.
// /Usage
// It can be called from a Result type or alone.
func Unwrap[T any](value T, err error) T {
	if err != nil {
		panic(fmt.Errorf("failed to unwrap: %v", err))
	}
	return value
}

func (r Result[T]) Unwrap() T {
	if r.Err != nil {
		panic(fmt.Errorf("failed to unwrap: %v", r.Err))
	}
	return r.Ok
}

// Returns the contained Err value, consuming the self value.
//
// Panics if the value is an Ok, with a custom panic message provided by the Ok’s value.
//
// Usage
// It can be called from a Result type or alone.
func UnwrapErr[T any](ok T, err error) error {
	if err == nil {
		panic(fmt.Errorf("failed to unwrap: error is nil"))
	}
	return err
}
func (r Result[T]) UnwrapErr() error {
	if r.Err == nil {
		panic(fmt.Errorf("failed to unwrap: %v", r.Err))
	}
	return r.Err
}

// Returns the contained Err value, consuming the self value, without checking that the value is not an Ok.
//
// Safety
// Calling this method on an Ok is undefined behavior.
func UnwrapErrUnchecked[T any](_ T, err error) error {
	return *(*error)(unsafe.Pointer(&err))
}
func (r Result[T]) UnwrapErrUnchecked() error {
	return *(*error)(unsafe.Pointer(&r.Err))
}

// Returns the contained Ok value or a provided default.
//
// Arguments passed to unwrap_or are eagerly evaluated; if you are passing the result
// of a function call, it is recommended to use UnwrapOrElse, which is lazily evaluated.
func (r Result[T]) UnwrapOr(defaultValue T) T {
	if r.Err != nil {
		return defaultValue
	}
	return r.Ok
}

func (r Result[T]) UnwrapOrElse(handleErrorFunc func(T) T) T {
	if r.Err != nil {
		return handleErrorFunc(r.Ok)
	}
	return r.Ok
}

// Returns the contained Ok value or a default
//
// Consumes the self argument then, if Ok, returns the contained value, otherwise
// if Err, returns the default value for that type.
func (r Result[T]) UnwrapOrDefault() T {
	//TODO: better handling of types
	if r.Err != nil {
		zero := reflect.Zero(reflect.TypeOf(r.Ok))
		return zero.Interface().(T)
	}
	return r.Ok
}

// Returns the contained Ok value, consuming the self value, without checking that
// the value is not an Err.
//
// **Calling this method on an Err is undefined behavior.
//
// /Usage
// It can be called from a Result type or alone.
func UnwrapUnchecked[T any](output T, _ error) T {
	return *(*T)(unsafe.Pointer(&output))
}

func (r Result[T]) UnwrapUnchecked() T {
	return *(*T)(unsafe.Pointer(&r.Ok))
}

// Returns the contained Ok value, consuming the self value.
//
// Because this function may panic, its use is generally discouraged. Instead,
// prefer to use pattern matching and handle the Err case explicitly, or call
// unwrap_or, unwrap_or_else, or unwrap_or_default.
//
// **Panics if the value is an Err, with a panic message including the passed message,
// and the content of the Err.
func (r Result[T]) Expect(expect string) T {
	if r.Err != nil {
		panic(fmt.Errorf("%s : %v", expect, r.Err))
	}
	return r.Ok
}

// Returns the contained Err value, consuming the self value.
//
// Panics if the value is an Ok, with a panic message including the passed
// message, and the content of the Ok.
func (r Result[T]) ExpectErr(msg string) error {
	if r.Err == nil {
		panic(fmt.Errorf("no error found: %v", msg))
	}
	return r.Err
}
