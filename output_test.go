package uhoh

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func ExampleErr_ToJson() {
	originalErr := errors.New("original error")
	describeErr := errors.New("describe error")

	err := New(originalErr, describeErr)
	err.SetDate(time.Date(2021, time.Month(9), 12, 1, 10, 30, 0, time.UTC))
	fmt.Printf("%s", err.ToJson())

	// Output:
	// {"date":"2021-09-12T01:10:30Z","describe":"describe error","file":"output_test.go","function":"ExampleErr_ToJson","line":14,"original":"original error"}
}

func TestToMapStrNil(t *testing.T) {
	var err *Err
	if err.ToMapStr() != nil {
		t.Errorf("ToMapStr() = %v, want %v", err.ToMapStr(), nil)
	}
}
