package uhoh

import (
	"fmt"
	"runtime"
	"strings"
)

// stackDepth returns the number of frames in the stack
var stackDepth = 3

// SetStackDepth sets the number of frames to skip when creating the stack
func SetStackDepth(depth int) { stackDepth = depth }

// Stack is the file, function, line of the stack trace
type Frame struct {
	file     string
	function string
	line     int
}

// Stack will return the stack trace
func (e *Err) Stack() []Frame { return e.stack }

// FirstFrame is the runtime.Frame.File stripped down to just the filename
func (e *Err) FirstStack() Frame { return e.stack[0] }

// File returns the file name
func (s *Frame) File() string { return s.file }

// Function returns the function name
func (s *Frame) Function() string { return s.function }

// Line returns the line number
func (s *Frame) Line() int { return s.line }

// String returns the stack as a string
func (s *Frame) String() string {
	return fmt.Sprintf("%s:%d %s", s.file, s.line, s.function)
}

// stackInfo returns []stack Frame skipping the number of supplied frames.
func stackInfo(skip int) []Frame {
	pc := make([]uintptr, stackDepth)
	_ = runtime.Callers(skip+2, pc)

	frames := runtime.CallersFrames(pc)

	var stack []Frame
	for {
		rf, hasMore := frames.Next()
		stack = append(stack, *frameDetails(rf))

		if !hasMore {
			break
		}
	}

	return stack
}

func frameDetails(rf runtime.Frame) *Frame {
	return &Frame{
		file:     rf.File[strings.LastIndexByte(rf.File, '/')+1:],
		function: rf.Function[strings.LastIndexByte(rf.Function, '.')+1:],
		line:     rf.Line,
	}
}
