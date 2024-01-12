package unwrap

import (
	"errors"
	"testing"
)

func TestUnwrap(t *testing.T) {
	testcase1 := Unwrap(divide(4, 2))
	if testcase1 != 2 {
		t.Errorf("expected: %v, got: %v", 2, testcase1)
	}
	testcase2 := Wrap(divide(4, 2)).Unwrap()
	if testcase2 != 2 {
		t.Errorf("expected: %v, got: %v", 2, testcase2)
	}
}

func TestUnwrapErr(t *testing.T) {
	if UnwrapErr(divide(4, 0)) == nil {
		t.Errorf("failed")
	}
	if Wrap(divide(4, 0)).UnwrapErr() == nil {
		t.Errorf("failed")
	}
}

func TestUnwrapErrUnchecked(t *testing.T) {
	if UnwrapErrUnchecked(divide(4, 0)) == nil {
		t.Errorf("failed")
	}
	if Wrap(divide(4, 0)).UnwrapErrUnchecked() == nil {
		t.Errorf("failed")
	}
}

func TestUnwrapOr(t *testing.T) {
	result := Wrap(divide(4, 0)).UnwrapOr(2)
	if result != 2 {
		t.Errorf("Expected %d, got %d", 2, result)
	}
}

// This test sucks, fix this. or fix the func
func TestUnwrapOrElse(t *testing.T) {
	errorTestFunc := func(x int) (int, error) {
		return x, errors.ErrUnsupported
	}

	alternativeFunc := func(x int) int {
		return x * x
	}

	result := Wrap(errorTestFunc(4)).UnwrapOrElse(alternativeFunc)
	if result != 16 {
		t.Errorf("Expected %d, got %d", 16, result)
	}
}

func TestUnwrapOrDefault(t *testing.T) {
	result := Wrap(divide(4, 0)).UnwrapOrDefault()
	if result != 0 {
		t.Errorf("Expected %d, got %d", 2, result)
	}
}

func TestUnwrapUnchecked(t *testing.T) {
	expected := 10
	result := UnwrapUnchecked(expected, nil)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}

	expectedStr := "hello"
	resultStr := UnwrapUnchecked(expectedStr, nil)
	if resultStr != expectedStr {
		t.Errorf("Expected %s, got %s", expectedStr, resultStr)
	}
}

func TestExpect(t *testing.T) {
	result := Wrap(divide(4, 2)).Expect("uh oh")
	if result != 2 {
		t.Errorf("Expected %d, got %d", 2, result)
	}
}

func TestExpectErr(t *testing.T) {
	result := Wrap(divide(4, 0)).ExpectErr("uh oh")
	if result == nil {
		t.Errorf("failed")
	}
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division failure")
	}
	return a / b, nil
}
