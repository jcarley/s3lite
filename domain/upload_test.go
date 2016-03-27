package domain

import (
	"bytes"
	"testing"

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
	etag, err := upload.AddPart(4, "ABCDEF", bytes)
	assert.Equal(t, etag, "_guU3FJJuNrrFtOChHxK2MZXPtQ=")
	assert.Equal(t, err, nil)
}
