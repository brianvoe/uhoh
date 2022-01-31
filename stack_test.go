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
			s := tt.err.FirstFrame()

			if s.File != tt.file {
				t.Errorf("TestStack File = %s, want %s", s.File, tt.file)
			}
			if s.Function != tt.function {
				t.Errorf("TestStack Function = %s, want %s", s.Function, tt.function)
			}
			if s.Line != tt.line {
				t.Errorf("TestStack Line = %d, want %d", s.Line, tt.line)
			}
			if s.String() != tt.str {
				t.Errorf("Stack.String() = %v, want %v", s.String(), tt.str)
			}
		})
	}
}

func TestFirstFrameZeroFrameStackDepth(t *testing.T) {
	ogStackDepth := stackDepth

	SetStackDepth(0)
	defer SetStackDepth(ogStackDepth)

	err := New(errors.New("original error"))

	ff := err.FirstFrame()
	if ff != nil {
		t.Error("FirstFrame should be nil")
	}

	// Test that the stack in Err is the correct length
	if len(err.Stack) != 0 {
		t.Error("Stack was not set correctly")
	}
}

func TestStackLevelHigherThanStackDepth(t *testing.T) {
	err := NewStackLevel(errors.New("original error"), 10)

	// Test that stack is empty
	if len(err.Stack) != 0 {
		t.Error("Stack was not set correctly")
	}

	// Try to grab the first frame
	ff := err.FirstFrame()
	if ff != nil {
		t.Error("FirstFrame should be nil")
	}

	// Get string representation of the stack
	str := err.FirstFrame().String()
	if str != "" {
		t.Error("Stack was not set correctly")
	}
}

func TestSetStackDepth(t *testing.T) {
	ogStackDepth := stackDepth

	SetStackDepth(2)
	defer SetStackDepth(ogStackDepth)

	err := New(errors.New("original error"))

	// Test that the stack depth is set correctly
	if stackDepth != 2 {
		t.Error("Stack depth was not set correctly")
	}

	// Test that the stack in Err is the correct length
	if len(err.Stack) != 2 {
		t.Error("Stack was not set correctly")
	}
}

// TestLargeStackDepth - Testing if the stack depth is set larger than the amount of stack frames there are
func TestLargeStackDepth(t *testing.T) {
	ogStackDepth := stackDepth

	SetStackDepth(10)
	defer SetStackDepth(ogStackDepth)

	err := New(errors.New("original error"))

	// Test that the stack depth is set correctly
	if stackDepth != 10 {
		t.Error("Stack depth was not set correctly")
	}

	// Test that the stack in Err is the correct length
	if len(err.Stack) > 5 {
		t.Errorf("Stack was not set correctly. Depth %d", len(err.Stack))
	}
}
