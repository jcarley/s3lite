package test

import (
	"encoding/json"
	"testing"
)

func GetRawData(t *testing.T, buffer []byte) (data map[string]interface{}) {
	err := json.Unmarshal(buffer, &data)
	if err != nil {
		t.Fatalf("Failed to unmarshal buffer: ", err)
	}
	return
}

func SetRawData(t *testing.T, v interface{}) (bytes []byte) {
	bytes, err := json.Marshal(v)
	if err != nil {
		t.Fatalf("Failed to marshal v: ", err)
	}
	return bytes
}
