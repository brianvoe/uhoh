package uhoh

import (
	"encoding/json"
	"time"
)

// ToJson converts the output to a json string
func (e *Err) ToJson() []byte {
	b, _ := json.Marshal(e.ToMapStr())
	return b
}

// ToMapStr converts Err to a map[string]interface{}
func (e *Err) ToMapStr() map[string]interface{} {
	m := make(map[string]interface{})
	if e.original != nil {
		m["original"] = e.original.Error()
	}
	if e.describe != nil {
		m["describe"] = e.describe.Error()
	}
	m["file"] = e.file
	m["function"] = e.function
	m["line"] = e.line
	m["date"] = e.date.Format(time.RFC3339)

	return m
}
