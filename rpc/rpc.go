package rpc

import (
	"encoding/json"
	"fmt"
)

func EncodeMessage(message any) string {
	content, err := json.Marshal(message)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("Content-Length: %d\r\n\r\n%s", len(content), content)
}
