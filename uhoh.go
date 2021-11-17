package uhoh

import (
	"time"
)

// Err is the base struct for this package
type Err struct {
	// Original error
	original error

	// Describe error
	describe error

	// Stack trace items
	file     string
	function string
	line     int

	// Date/time in which error was created
	date time.Time
}

// New will return an intialized error with a stack trace level
func New(original error, describe error) *Err {
	return NewStackLevel(original, describe, 2)
}

// NewStackLevel will return an initialized error with a stack trace level
func NewStackLevel(original error, describe error, level int) *Err {
	file, function, line := stackLevel(level)
	return &Err{
		original: original,
		describe: describe,
		file:     file,
		function: function,
		line:     line,
		date:     time.Now().UTC(),
	}
}

// Original will return the original error
func (e *Err) Original() error { return e.original }

// Describe will set the describe error
func (e *Err) Describe() error { return e.describe }

// SetDescribe will set the describe error
func (e *Err) SetDescribe(describe error) *Err { e.describe = describe; return e }

// Date returns the date/time in which the error was created
func (e *Err) Date() time.Time { return e.date }

// SetDate will set the date/time in which the error was created
func (e *Err) SetDate(date time.Time) *Err { e.date = date; return e }

// Error will return the desbribe error if it exists, otherwise the original error
func (e *Err) Error() string {
	if e.describe != nil {
		return e.describe.Error()
	}
	return e.original.Error()
}

// Unwrap returns the original error
func (e *Err) Unwrap() error { return e.original }

// Is will check to see if target is the original or describe error
func (e *Err) Is(target error) bool { return target == e.original || target == e.describe }

// IsOriginal will check to see if target is the original error
func (e *Err) IsOriginal(target error) bool { return target == e.original }

// IsDescribe will check to see if target is the describe error
func (e *Err) IsDescribe(target error) bool { return target == e.describe }
