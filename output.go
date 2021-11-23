package uhoh

import (
	"encoding/json"
	"strings"
	"time"
)

// defaultErrorFormat variable function to format for the Error output
var defaultErrorFormat = func(err *Err) string {
	if err == nil {
		return ""
	}

	// Put together an array of date, type, original, and describe
	var s []string
	s = append(s, err.Date.Format(time.RFC3339))
	if err.Type != nil {
		s = append(s, err.Type.Error())
	}
	if err.Original != nil {
		s = append(s, err.Original.Error())
	}
	if err.Describe != nil {
		s = append(s, err.Describe.Error())
	}

	// Take string array and join with ": "
	return strings.Join(s, " | ")
}

// SetDefaultErrorFormatter sets the defaultErrorFormatter function
func SetDefaultErrorFormatter(function func(err *Err) string) {
	defaultErrorFormat = function
}

// Error will return the date, type error(if set) original error and describe error(if set)
// as a string in the default error format
func (e *Err) Error() string {
	return defaultErrorFormat(e)
}

// ToJson converts the output to a json string
func (e *Err) ToJSON() []byte {
	b, _ := json.Marshal(e.ToMapStr())
	return b
}

// ToMapStr converts Err to a map[string]interface{}
func (e *Err) ToMapStr() map[string]interface{} {
	if e == nil {
		return nil
	}

	m := make(map[string]interface{})
	if e.Type != nil {
		m["type"] = e.Type.Error()
	}
	if e.Original != nil {
		m["original"] = e.Original.Error()
	}
	if e.Describe != nil {
		m["describe"] = e.Describe.Error()
	}
	if len(e.Stack) > 0 {
		m["stack"] = e.Stack
	}
	m["date"] = e.Date.Format(time.RFC3339)

	return m
}
