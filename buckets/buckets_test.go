package buckets

import (
  "encoding/xml"
  "fmt"
  "github.com/jcarley/s3lite/infrastructure"
  "github.com/jcarley/s3lite/webservice"
  "github.com/stretchr/testify/assert"
  "net/http"
  "testing"
)

func addHeaders(req *http.Request) {
  req.Header.Add("Authorization", "GHIJKLMNOPQRSTUV1234567890")
  req.Header.Add("Content-Type", "binary/octel-stream")
  req.Header.Add("Content-Length", "0")
  req.Header.Add("Date", "Wed, 01 Mar  2006 12:00:00 GMT")
}

func testPutBucket(t *testing.T) {
  req, _ := http.NewRequest("PUT", "http://bucket-us-west.s3.example.com", nil)
  addHeaders(req)

}
