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
		line     int
		function string
	}{
		{
			err:      New(originalErr, describeErr),
			name:     "stack0",
			file:     "stack_test.go",
			function: "TestStack",
			line:     19,
		},
		{
			err:      New(originalErr, describeErr),
			name:     "stack1",
			file:     "stack_test.go",
			function: "TestStack",
			line:     26,
		},
		{
			err:      NewStackLevel(originalErr, describeErr, 1),
			name:     "stack1",
			file:     "stack_test.go",
			function: "TestStack",
			line:     33,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.err.File() != tt.file {
				t.Errorf("TestStack File() = %s, want %s", tt.err.File(), tt.file)
			}
			if tt.err.Function() != tt.function {
				t.Errorf("TestStack Function() = %s, want %s", tt.err.Function(), tt.function)
			}
			if tt.err.Line() != tt.line {
				t.Errorf("TestStack Line() = %d, want %d", tt.err.Line(), tt.line)
			}
		})
	}
}
