package buckets

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func addHeaders(req *http.Request) {
	req.Header.Add("Authorization", "GHIJKLMNOPQRSTUV1234567890")
	req.Header.Add("Content-Disposition", "attachment; filename=foobar.mov")
	req.Header.Add("Content-Type", "binary/octel-stream")
	req.Header.Add("x-amz-acl", "private")
	req.Header.Add("x-amz-server-side-encryption", "AES256")
}

func TestCreateBucket_ReturnsBucketName(t *testing.T) {
	req, _ := http.NewRequest("POST", "http://bucket-us-west.s3.example.com/uploads/path/to/my/object", nil)
	addHeaders(req)

	status, bucketName := CreateBucket(req, nil)

	assert.Equal(t, status, 201)
	assert.Equal(t, bucketName, "bucket-us-west")
}
