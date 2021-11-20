package uhoh

import (
	"errors"
	"testing"
)

func TestStack(t *testing.T) {
	originalErr := errors.New("original error")
	describeErr := errors.New("describe error")
	tests := []struct {
		err      *Err
		name     string
		file     string
		function string
		line     int
		str      string
	}{
		{
			err:      New(originalErr).SetDescribe(describeErr),
			name:     "stack0",
			file:     "stack_test.go",
			function: "TestStack",
			line:     20,
			str:      "stack_test.go:20 TestStack",
		},
		{
			err:      New(originalErr).SetDescribe(describeErr),
			name:     "stack1",
			file:     "stack_test.go",
			function: "TestStack",
			line:     28,
			str:      "stack_test.go:28 TestStack",
		},
		{
			err:      NewStackLevel(originalErr, 1).SetDescribe(describeErr),
			name:     "stack1",
			file:     "stack_test.go",
			function: "TestStack",
			line:     36,
			str:      "stack_test.go:36 TestStack",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := tt.err.FirstStack()

			if s.File() != tt.file {
				t.Errorf("TestStack File() = %s, want %s", s.File(), tt.file)
			}
			if s.Function() != tt.function {
				t.Errorf("TestStack Function() = %s, want %s", s.Function(), tt.function)
			}
			if s.Line() != tt.line {
				t.Errorf("TestStack Line() = %d, want %d", s.Line(), tt.line)
			}
			if s.String() != tt.str {
				t.Errorf("Stack.String() = %v, want %v", s.String(), tt.str)
			}
		})
	}
}
