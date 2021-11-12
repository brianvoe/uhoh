package uhoh

import (
	"errors"
	"fmt"
	"time"
)

func ExampleToJson() {
	originalErr := errors.New("original error")
	describeErr := errors.New("describe error")

	err := New(originalErr, describeErr)
	err.SetDate(time.Date(2021, time.Month(9), 12, 1, 10, 30, 0, time.UTC))
	fmt.Printf("%s", err.ToJson())

	// Output:
	// {"date":"2021-09-12T01:10:30Z","describe":"describe error","file":"output_test.go","function":"ExampleToJson","line":13,"original":"original error"}
}
