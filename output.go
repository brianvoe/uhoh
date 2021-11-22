package uhoh

import (
	"encoding/json"
	"errors"
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
func SetDefaultErrorFormatter(f func(err *Err) string) {
	defaultErrorFormat = f
}

// Error will return the date, type error(if set) original error and describe error(if set)
// as a string in the default error format
func (e *Err) Error() string {
	return defaultErrorFormat(e)
}

// Create MarshalJSON method to handle the json.Marshal for all fields
func (e *Err) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.ToMapStr())
}

// Create UnmarshalJSON method to handle the json.Unmarshal
func (e *Err) UnmarshalJSON(b []byte) error {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}

	if typeErr, ok := m["type"]; ok {
		e.Type = errors.New(typeErr.(string))
	}
	if originalErr, ok := m["original"]; ok {
		e.Original = errors.New(originalErr.(string))
	}
	if describeErr, ok := m["describe"]; ok {
		e.Describe = errors.New(describeErr.(string))
	}
	if stack, ok := m["stack"]; ok {
		e.Stack = stack.([]Frame)
	}
	if date, ok := m["date"]; ok {
		var err error
		e.Date, err = time.Parse(time.RFC3339, date.(string))
		if err != nil {
			return err
		}
	}

	return nil
}

// ToJson converts the output to a json string
func (e *Err) ToJson() []byte {
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
