package uhoh

import (
	"encoding/json"
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

// Create MarshalJSON for Frame for all fields
func (f *Frame) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("{\"file\":\"%s\",\"function\":\"%s\",\"line\":%d}", f.File, f.Function, f.Line)), nil
}

// Create UnmarshalJSON for Frame for all fields
func (f *Frame) UnmarshalJSON(b []byte) error {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}

	f.File = m["file"].(string)
	f.Function = m["function"].(string)
	f.Line = int(m["line"].(float64))

	return nil
}

// FirstFrame is the runtime.Frame.File stripped down to just the filename
func (e *Err) FirstStack() Frame { return e.Stack[0] }

// String returns the stack as a string
func (s *Frame) String() string {
	return fmt.Sprintf("%s:%d %s", s.File, s.Line, s.Function)
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
		File:     rf.File[strings.LastIndexByte(rf.File, '/')+1:],
		Function: rf.Function[strings.LastIndexByte(rf.Function, '.')+1:],
		Line:     rf.Line,
	}
}
