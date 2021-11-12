package uhoh

import (
	"runtime"
	"strings"
)

// File is the runtime.Frame.File stripped down to just the filename
func (e *Err) File() string { return e.file }

// Function is the runtime.Frame.Function stripped down to just the function name
func (e *Err) Function() string { return e.function }

// Line is the line of the runtime.Frame and exposed for convenience.
func (e *Err) Line() int { return e.line }

// stackLevel returns a stack Frame skipping the number of supplied frames.
// This is primarily used by other libraries who use this package
// internally as the additional.
func stackLevel(skip int) (file, function string, line int) {
	var frame [3]uintptr
	runtime.Callers(skip+2, frame[:])
	rf, _ := runtime.CallersFrames(frame[:]).Next()
	return rf.File[strings.LastIndexByte(rf.File, '/')+1:], rf.Function[strings.LastIndexByte(rf.Function, '.')+1:], rf.Line
}
