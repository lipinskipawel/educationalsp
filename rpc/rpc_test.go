package rpc_test

import (
	"testing"

	"github.com/lipinskipawel/educationalsp/rpc"
)

type EncodingExample struct {
	Testing bool
}

func TestEncode(t *testing.T) {
	expected := "Content-Length: 16\r\n\r\n{\"Testing\":true}"
	actual := rpc.EncodeMessage(EncodingExample{Testing: true})

	if expected != actual {
		t.Fatalf("Expected: %s, Actual: %s", expected, actual)
	}
}

func TestDecode(t *testing.T) {
	incomingMessage := "Content-Length: 16\r\n\r\n{\"Testing\":true}"
	contentLength, err := rpc.DecodeMesage([]byte(incomingMessage))
	if err != nil {
		t.Fatal(err)
	}

	if contentLength != 16 {
		t.Fatalf("Expected: 16, Actual: %d", contentLength)
	}
}
