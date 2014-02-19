package domain

import (
  "testing"
  "bytes"

  "github.com/stretchr/testify/assert"
)

var (
  buffer *bytes.Buffer
)

func init() {
  buffer = bytes.NewBufferString("ABCDEFGHIGKLMNOPQRSTUVWXYZ1234567890")
}

func TestGetNewUploadIdGeneratesAnId(t *testing.T) {
  upload := Upload{}
  upload.GetNewUploadId()
  assert.NotNil(t, upload.UploadId)
}

func TestAddPartReturnsETag(t *testing.T) {
  bytes := buffer.Bytes()
  upload := Upload{}
  etag := upload.AddPart(4, "ABCDEF", bytes)
  assert.Equal(t, etag, "_guU3FJJuNrrFtOChHxK2MZXPtQ=")
}

