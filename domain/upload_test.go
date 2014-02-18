package domain

import (
  "testing"
  "bytes"

  "github.com/stretchr/testify/assert"
)

func TestGetNewUploadIdGeneratesAnId(t *testing.T) {
  upload := Upload{}
  upload.GetNewUploadId()
  assert.NotNil(t, upload.UploadId)
}


func TestAddPartReturnsETag(t *testing.T) {
  buffer := bytes.NewBufferString("ABCDEFGHIGKLMNOPQRSTUVWXYZ1234567890")
  bytes := buffer.Bytes()
  upload := Upload{}
  etag := upload.AddPart(4, "ABCDEF", bytes)
  assert.Equal(t, etag, "_guU3FJJuNrrFtOChHxK2MZXPtQ=")
}
