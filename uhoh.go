package uhoh

import (
	"time"
)

// Err is the base struct for this package
type Err struct {
	// Errors
	originalErr error
	describeErr error
	typeErr     error

	// Stack traces
	stack []Frame

	// Date/time in which error was created
	date time.Time
}

// New will return an intialized error with a stack trace level
func New(err error) *Err {
	return NewStackLevel(err, 2)
}

// NewStackLevel will return an initialized error with a stack trace at certain level
func NewStackLevel(err error, level int) *Err {
	return &Err{
		originalErr: err,
		stack:       stackInfo(level),
		date:        time.Now().UTC(),
	}
}

// Original will return the original error
func (e *Err) Original() error { return e.originalErr }

// Describe will set the describe error
func (e *Err) Describe() error { return e.describeErr }

// SetDescribe will set the describe error
func (e *Err) SetDescribe(describe error) *Err { e.describeErr = describe; return e }

// Type will return the type error
func (e *Err) Type() error { return e.typeErr }

// SetType will set the type error
func (e *Err) SetType(target error) *Err { e.typeErr = target; return e }

// Date returns the date/time in which the error was created
func (e *Err) Date() time.Time { return e.date }

// SetDate will set the date/time in which the error was created
func (e *Err) SetDate(date time.Time) *Err { e.date = date; return e }

// Unwrap returns the original error
func (e *Err) Unwrap() error { return e.originalErr }

// Is will check to see if target is the original or describe error
func (e *Err) Is(target error) bool {
	return target == e.originalErr || target == e.describeErr || target == e.typeErr
}

// IsOriginal will check to see if target is the original error
func (e *Err) IsOriginal(target error) bool { return target == e.originalErr }

// IsDescribe will check to see if target is the describe error
func (e *Err) IsDescribe(target error) bool { return target == e.describeErr }

// IsType will check to see if target is the type error
func (e *Err) IsType(target error) bool { return target == e.typeErr }
