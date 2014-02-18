package controllers

import (
  "fmt"
  "testing"
  "net/http"
  "encoding/xml"
  "github.com/jcarley/s3lite/domain"
  "github.com/jcarley/s3lite/webservice"
  "github.com/stretchr/testify/assert"
)

func addHeaders(req *http.Request) {
  req.Header.Add("Authorization", "GHIJKLMNOPQRSTUV1234567890")
  req.Header.Add("Content-Disposition", "attachment; filename=foobar.mov")
  req.Header.Add("Content-Type", "binary/octel-stream")
  req.Header.Add("x-amz-acl", "private")
  req.Header.Add("x-amz-server-side-encryption", "AES256")
}

func GetResultMultipartUploadResult(response string) *webservice.InitiateMultipartUploadResult {
  result := webservice.InitiateMultipartUploadResult{}

  err := xml.Unmarshal([]byte(response), &result)
  if err != nil {
    fmt.Printf("error: %v", err)
    return nil
  }

  return &result
}


func TestInitiateMultipartUploadReturnsAnUploadId(t *testing.T) {

  req, _ := http.NewRequest("POST", "http://bucket-us-west.s3.example.com/uploads/path/to/my/object", nil)
  addHeaders(req)
  db := domain.NewInMemoryDatabase()
  controller := &UploadController{}

  status, response := controller.InitiateMultipartUpload(req, db)
  result := GetResultMultipartUploadResult(response)

  assert.Equal(t, status, 201)
  assert.NotNil(t, result)
  assert.NotNil(t, result.UploadId)
}

func TestRecordedPartHasUploadId(t *testing.T) {

  req, _ := http.NewRequest("POST", "http://bucket-us-west.s3.example.com/uploads/path/to/my/object", nil)
  addHeaders(req)
  db := domain.NewInMemoryDatabase()
  controller := &UploadController{}

  _, response := controller.InitiateMultipartUpload(req, db)
  result := GetResultMultipartUploadResult(response)

  upload := db.GetUploadByUploadId(result.UploadId)

  assert.NotNil(t, result)
  assert.NotNil(t, upload)
  assert.Equal(t, result.UploadId, upload.UploadId)
}

func TestRecordedPartHasFilename(t *testing.T) {
  req, _ := http.NewRequest("POST", "http://bucket-us-west.s3.example.com/uploads/path/to/my/object", nil)
  addHeaders(req)
  db := domain.NewInMemoryDatabase()
  controller := &UploadController{}

  _, response := controller.InitiateMultipartUpload(req, db)
  result := GetResultMultipartUploadResult(response)

  upload := db.GetUploadByUploadId(result.UploadId)

  assert.Equal(t, "foobar.mov", upload.Filename)
}

func TestRecordedPartHasBucketWhenValidSubdomain(t *testing.T) {
  req, _ := http.NewRequest("POST", "http://bucket-us-west.s3.example.com/uploads/path/to/my/object", nil)
  addHeaders(req)
  db := domain.NewInMemoryDatabase()
  controller := &UploadController{}

  _, response := controller.InitiateMultipartUpload(req, db)
  result := GetResultMultipartUploadResult(response)

  upload := db.GetUploadByUploadId(result.UploadId)

  assert.Equal(t, "bucket-us-west", upload.Bucket)
}

func TestRecordedPartHasDefaultBucketWhenInValidSubdomain(t *testing.T) {
  req, _ := http.NewRequest("POST", "http://example.com/uploads/path/to/my/object", nil)
  addHeaders(req)

  db := domain.NewInMemoryDatabase()
  controller := &UploadController{}

  _, response := controller.InitiateMultipartUpload(req, db)
  result := GetResultMultipartUploadResult(response)

  upload := db.GetUploadByUploadId(result.UploadId)

  assert.Equal(t, "default", upload.Bucket)
}

func TestRecordedPartHasKey(t *testing.T) {
  req, _ := http.NewRequest("POST", "http://bucket-us-west.s3.example.com/path/to/my/object", nil)
  addHeaders(req)
  db := domain.NewInMemoryDatabase()
  controller := &UploadController{}

  _, response := controller.InitiateMultipartUpload(req, db)
  result := GetResultMultipartUploadResult(response)

  upload := db.GetUploadByUploadId(result.UploadId)

  assert.Equal(t, "path/to/my/object", upload.Key)
}


