package uhoh

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func ExampleSetDefaultErrorFormatter() {
	// Set original formatter to be used to reset it later
	originalFormatter := defaultErrorFormat

	SetDefaultErrorFormatter(func(err *Err) string {
		return err.Original.Error()
	})

	originalErr := errors.New("original error")
	describeErr := errors.New("describe error")

	err := New(originalErr).SetDescribe(describeErr).SetType(ErrGeneral)
	fmt.Printf("%s", err.Error())

	// Set the default error formatter back to the default
	SetDefaultErrorFormatter(originalFormatter)

	// Output:
	// original error
}

func ExampleErr_Error() {
	originalErr := errors.New("original error")
	describeErr := errors.New("describe error")

	err := New(originalErr).SetDescribe(describeErr).SetType(ErrGeneral)
	err.SetDate(time.Date(2021, time.Month(9), 12, 1, 10, 30, 0, time.UTC))
	fmt.Printf("%s", err.Error())

	// Output:
	// 2021-09-12T01:10:30Z | general error | original error | describe error
}

func TestErr_Error(t *testing.T) {
	originalErr := errors.New("original error")
	describeErr := errors.New("describe error")

	err := New(originalErr).SetDescribe(describeErr).SetType(ErrGeneral)
	err.SetDate(time.Date(2021, time.Month(9), 12, 1, 10, 30, 0, time.UTC))
	if err.Error() != "2021-09-12T01:10:30Z | general error | original error | describe error" {
		t.Errorf("Error() = %v, want %v", err.Error(), "2021-09-12T01:10:30Z | general error | original error | describe error")
	}
}

func TestErr_ErrorNil(t *testing.T) {
	var err *Err
	if err.Error() != "" {
		t.Errorf("Error() = %v, want %v", err.Error(), "")
	}
}

func ExampleErr_ToJson() {
	originalErr := errors.New("original error")
	describeErr := errors.New("describe error")

	err := New(originalErr).SetDescribe(describeErr).SetType(ErrGeneral)
	err.SetDate(time.Date(2021, time.Month(9), 12, 1, 10, 30, 0, time.UTC))
	fmt.Printf("%s", err.ToJSON())

	// Output:
	// {"date":"2021-09-12T01:10:30Z","describe":"describe error","original":"original error","stack":[{"file":"output_test.go","function":"ExampleErr_ToJson","line":65},{"file":"run_example.go","function":"runExample","line":64},{"file":"example.go","function":"runExamples","line":44}],"type":"general error"}
}

func TestToMapStrNil(t *testing.T) {
	var err *Err
	if err.ToMapStr() != nil {
		t.Errorf("ToMapStr() = %v, want %v", err.ToMapStr(), nil)
	}
}
