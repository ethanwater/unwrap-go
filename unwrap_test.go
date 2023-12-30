package unwrap

import (
	"errors"
	"fmt"
	"testing"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func TestUnwrap(t *testing.T) {
	result := Unwrap(divide(4, 2))
	if result != 2 {
		t.Errorf("expected: %v, got: %v", 2, result)
	}
}

//func TestExpect(t *testing.T) {
//	result := Raw(divide(4, 0)).Expect("uh oh")
//	fmt.Println(result)
//}

func TestUnwrapOr(t *testing.T) {
	result := Raw(divide(4, 0)).Or(2)
	fmt.Println(result)
}

func TestUnchecked(t *testing.T) {
	expected := 10
	result := Unchecked(expected, nil)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}

	expectedStr := "hello"
	resultStr := Unchecked(expectedStr, nil)
	if resultStr != expectedStr {
		t.Errorf("Expected %s, got %s", expectedStr, resultStr)
	}
}

//func TestOrElse(t *testing.T) {
//handleErrorFunc := func(value int) int {
//return value * 2
//}

//result := Raw(divide(4, 0)).OrElse(handleErrorFunc)

//if result != 8 {
//t.Errorf("Expected %d, got %d", 8, result)
//}

//err := errors.New("test error")
//resWithError := Result[int]{Value: 10, Err: err}
//expectedErrorValue := handleErrorFunc(resWithError.Value)
//resultWithError := resWithError.OrElse(handleErrorFunc)

//if resultWithError != expectedErrorValue {
//t.Errorf("Expected %d, got %d", expectedErrorValue, resultWithError)
//}
//}
