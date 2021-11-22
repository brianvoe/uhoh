package uhoh

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func Example() {
	// Errors
	originalErr := errors.New("original error")
	describeErr := errors.New("describe error")

	// Create error
	err := New(originalErr)
	err.SetDescribe(describeErr)
	err.SetType(ErrGeneral)

	// Can set date if need be
	err.SetDate(time.Date(2021, time.Month(9), 12, 1, 20, 30, 0, time.UTC))

	// Output info
	fmt.Println(err.Error()) // Will prioritize describe error
	fmt.Println(err.Original)
	fmt.Println(err.Describe)
	fmt.Println(err.Stack)

	// Output:
	// 2021-09-12T01:20:30Z | general error | original error | describe error
	// original error
	// describe error
	// [{uhoh_test.go Example 16} {run_example.go runExample 64} {example.go runExamples 44}]
}

func BenchmarkNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		original := errors.New("original")
		describe := errors.New("describe")

		New(original).SetDescribe(describe).SetType(ErrGeneral)
	}
}

func ExampleNew() {
	// Original error
	originalErr := errors.New("original error")

	err := New(originalErr)
	fmt.Println(err.Error())

	// Output:
	// original error
}

func ExampleNewStackLevel() {
	// Original error
	originalErr := errors.New("original error")

	// Create error with stack level
	err := NewStackLevel(originalErr, 1)
	fmt.Println(err.Error())

	// Output:
	// original error
}

func ExampleErr_Original() {
	// Errors
	originalErr := errors.New("original error")

	err := New(originalErr)
	fmt.Println(err.Error())

	// Output:
	// original error
}

func ExampleErr_Describe() {
	// Errors
	originalErr := errors.New("original error")
	describeErr := errors.New("describe error")

	err := New(originalErr).SetDescribe(describeErr)
	fmt.Println(err.Error())

	// Output:
	// describe error
}

func ExampleErr_Type() {
	// Errors
	originalErr := errors.New("original error")

	err := New(originalErr).SetType(ErrGeneral)
	fmt.Println(err.Type)

	// Output:
	// uhoh
}

func ExampleErr_SetDescribe() {
	// Errors
	originalErr := errors.New("original error")

	err := New(originalErr)
	err.SetDescribe(errors.New("new describe error"))
	fmt.Println(err.Error())

	// Output:
	// new describe error
}

func ExampleErr_Date() {
	// Errors
	originalErr := errors.New("original error")

	err := New(originalErr)
	err.SetDate(time.Date(2021, time.Month(9), 12, 1, 20, 30, 0, time.UTC))
	fmt.Println(err.Date)

	// Output:
	// 2021-09-12 01:20:30 +0000 UTC
}

func ExampleErr_Unwrap() {
	// Errors
	originalErr := errors.New("original error")
	describeErr := errors.New("describe error")

	err := New(originalErr).SetDescribe(describeErr)
	fmt.Println(err.Unwrap().Error())

	// Output:
	// original error
}

func TestIs(t *testing.T) {
	original := errors.New("original")
	describe := errors.New("describe")
	err := New(original).SetDescribe(describe).SetType(ErrGeneral)

	// Check to make sure Is is either original or describe
	if !errors.Is(err, original) && !errors.Is(err, describe) && !errors.Is(err, ErrGeneral) {
		t.Error("Error should be original, describe or type error")
	}

	// Check if is original error
	if !err.IsOriginal(original) {
		t.Error("Error did not match original")
	}

	// Check if is describe error
	if !err.IsDescribe(describe) {
		t.Error("Error did not match describe")
	}

	// Check if is type error
	if !err.IsType(ErrGeneral) {
		t.Error("Error did not match type")
	}
}
