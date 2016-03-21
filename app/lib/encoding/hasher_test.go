package encoding

import (
  "bytes"
  "testing"

  "github.com/stretchr/testify/assert"
)


func TestMD5Set(t *testing.T) {
  h := NewHasher("MD5", nil)
  assert.Equal(t, h.Algorithm, "MD5")
}

func TestSHA1Set(t *testing.T) {
  h := NewHasher("SHA1", nil)
  assert.Equal(t, h.Algorithm, "SHA1")
}

func TestAdler32Set(t *testing.T) {
  h := NewHasher("ADLER32", nil)
  assert.Equal(t, h.Algorithm, "ADLER32")
}

func TestHashStringReturnsCorrectHash(t *testing.T) {
  buffer := bytes.NewBufferString("ABCDEFGHIGKLMNOPQRSTUVWXYZ1234567890")
  h := NewHasher("SHA1", buffer)
  hashString := h.HashString()
  assert.Equal(t, hashString, "fNCakRJEKjfBPy85eNs5xeoEiOE=")
}
