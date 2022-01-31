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

// Frame contains the file, function, line of the stack trace
type Frame struct {
	File     string `json:"file"`
	Function string `json:"function"`
	Line     int    `json:"line"`
}

// FirstFrame is the runtime.Frame.File stripped down to just the filename
func (e *Err) FirstFrame() *Frame {
	// Deal with the case where the stack is empty
	if len(e.Stack) == 0 {
		return nil
	}

	return &e.Stack[0]
}

// String returns the stack as a string
func (f *Frame) String() string {
	if f == nil {
		return ""
	}

	return fmt.Sprintf("%s:%d %s", f.File, f.Line, f.Function)
}

// stackInfo returns []stack Frame skipping the number of supplied frames.
func stackInfo(skip int) []Frame {
	pc := make([]uintptr, stackDepth)
	_ = runtime.Callers(skip+2, pc)

	frames := runtime.CallersFrames(pc)

	var stack []Frame
	for {
		rf, hasMore := frames.Next()
		fd := frameDetails(rf)
		if fd != nil {
			stack = append(stack, *fd)
		}

		if !hasMore {
			break
		}
	}

	return stack
}

func frameDetails(rf runtime.Frame) *Frame {
	file := rf.File[strings.LastIndexByte(rf.File, '/')+1:]
	function := rf.Function[strings.LastIndexByte(rf.Function, '.')+1:]
	line := rf.Line
	if file == "" || function == "" || line == 0 {
		return nil
	}

	return &Frame{
		File:     file,
		Function: function,
		Line:     line,
	}
}
