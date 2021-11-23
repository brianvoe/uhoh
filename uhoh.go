package uhoh

import (
	"errors"
	"time"
)

// Err is the base struct for this package
type Err struct {
	// Errors
	Original error `json:"original"`
	Describe error `json:"describe"`
	Type     error `json:"type"`

	// Stack traces
	Stack []Frame `json:"stack"`

	// Date/time in which error was created
	Date time.Time `json:"date"`
}

// New will return an intialized error with a stack trace level
func New(err error) *Err {
	return NewStackLevel(err, 2)
}

// NewStr will return an intialized error with a stack trace level
func NewStr(err string) *Err {
	return NewStackLevel(errors.New(err), 2)
}

// NewStackLevel will return an initialized error with a stack trace at certain level
func NewStackLevel(err error, level int) *Err {
	// Check if error is already an uhoh.Err
	if e, ok := err.(*Err); ok {
		return e
	}

	return &Err{
		Original: err,
		Stack:    stackInfo(level),
		Date:     time.Now().UTC(),
	}
}

// SetDescribe will set the describe error
func (e *Err) SetDescribe(describe error) *Err { e.Describe = describe; return e }

// SetDescribeStr will set the describe error from a string
func (e *Err) SetDescribeStr(describe string) *Err { e.Describe = errors.New(describe); return e }

// SetType will set the type error
func (e *Err) SetType(target error) *Err { e.Type = target; return e }

// SetTypeStr will set the type error from a string
func (e *Err) SetTypeStr(target string) *Err { e.Type = errors.New(target); return e }

// SetDate will set the date/time in which the error was created
func (e *Err) SetDate(date time.Time) *Err { e.Date = date; return e }

// Unwrap returns the original error
func (e *Err) Unwrap() error { return e.Original }

// Is will check to see if target is the original or describe error
func (e *Err) Is(target error) bool {
	return target == e.Original || target == e.Describe || target == e.Type
}

// IsOriginal will check to see if target is the original error
func (e *Err) IsOriginal(target error) bool { return target == e.Original }

// IsDescribe will check to see if target is the describe error
func (e *Err) IsDescribe(target error) bool { return target == e.Describe }

// IsType will check to see if target is the type error
func (e *Err) IsType(target error) bool { return target == e.Type }
