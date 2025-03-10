package rpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

func EncodeMessage(message any) string {
	content, err := json.Marshal(message)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("Content-Length: %d\r\n\r\n%s", len(content), content)
}

func DecodeMesage(message []byte) (int, error) {
	header, content, found := bytes.Cut(message, []byte{'\r', '\n', '\r', '\n'})
	if !found {
		return 0, errors.New("Did not find separator")
	}

	contentLengthBytes := header[len("Content-Length: "):]
	contentLength, err := strconv.Atoi(string(contentLengthBytes))
	if err != nil {
		return 0, err
	}

	_ = content

	return contentLength, nil
}
