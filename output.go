package uhoh

import (
	"encoding/json"
	"time"
)

// Error will return the desbribe error if it exists, otherwise the original error
func (e *Err) Error() string {
	if e.describeErr != nil {
		return e.describeErr.Error()
	}
	return e.originalErr.Error()
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
	if e.originalErr != nil {
		m["original"] = e.originalErr.Error()
	}
	if e.describeErr != nil {
		m["describe"] = e.describeErr.Error()
	}
	m["stack"] = e.stack
	m["date"] = e.date.Format(time.RFC3339)

	return m
}
