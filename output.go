package uhoh

import (
	"encoding/json"
	"time"
)

// ToJson converts the output to a json string
func (e *Err) ToJson() []byte {
	b, _ := json.Marshal(e.errToMapStr())
	return b
}

// errToMapStr converts Err to a map[string]interface{}
func (e *Err) errToMapStr() map[string]interface{} {
	return map[string]interface{}{
		"original": e.original.Error(),
		"describe": e.describe.Error(),
		"file":     e.file,
		"function": e.function,
		"line":     e.line,
		"date":     e.date.Format(time.RFC3339),
	}
}
