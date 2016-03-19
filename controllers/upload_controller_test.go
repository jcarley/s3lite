package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/onsi/gomega"
)

func addHeaders(req *http.Request) {
	req.Header.Add("Authorization", "GHIJKLMNOPQRSTUV1234567890")
	req.Header.Add("Content-Disposition", "attachment; filename=foobar.mov")
	req.Header.Add("Content-Type", "binary/octel-stream")
	req.Header.Add("Upload-Id", "GHIJKLMNOPQRSTUV1234567890")
	req.Header.Add("Part-Number", "1")
	req.Header.Add("x-amz-acl", "private")
	req.Header.Add("x-amz-server-side-encryption", "AES256")
}

func GetRawData(t *testing.T, buffer []byte) (data map[string]interface{}) {
	err := json.Unmarshal(buffer, &data)
	if err != nil {
		t.Fatalf("Failed to unmarshal buffer: ", err)
	}
	return
}

func TestInitiateMultipartUploadReturnsAnUploadId(t *testing.T) {
	RegisterTestingT(t)

	controller := &UploadController{}

	req, _ := http.NewRequest("POST", "http://bucket-us-west.s3.example.com/uploads/path/to/my/object", nil)
	addHeaders(req)

	w := httptest.NewRecorder()

	controller.InitiateMultipartUpload(w, req)

	data := GetRawData(t, w.Body.Bytes())

	Expect(data["upload_id"]).ToNot(BeNil(), "Should have an upload id")
	Expect(w.Code).To(Equal(http.StatusOK), "Should receive 200 status")
}

func TestUploadPartHasRequiredHeaders(t *testing.T) {
	RegisterTestingT(t)

	cases := []struct {
		Header string
		Err    error
	}{
		{"Upload-Id", MissingUploadIdError},
		{"Part-Number", MissingPartNumberError},
		{"Authorization", MissingAuthorizationKeyError},
		{"Content-Disposition", MissingContentDispostionError},
	}

	for _, tc := range cases {
		controller := &UploadController{}
		req, _ := http.NewRequest("PUT", "http://bucket-us-west.s3.example.com/uploads/path/to/my/object", nil)
		addHeaders(req)
		req.Header.Del(tc.Header)
		w := httptest.NewRecorder()
		controller.UploadPart(w, req)
		data := GetRawData(t, w.Body.Bytes())
		Expect(data["error"]).To(Equal(tc.Err.Error()))
	}

}

func TestUploadPartHasBucketWhenValidSubdomain(t *testing.T) {
	// req, _ := http.NewRequest("POST", "http://bucket-us-west.s3.example.com/uploads/path/to/my/object", nil)
	// addHeaders(req)
	// db := domain.NewInMemoryDatabase()
	// controller := &UploadController{}

	// _, response := controller.InitiateMultipartUpload(req, db)
	// result := GetResultMultipartUploadResult(response)

	// upload := db.GetUploadByUploadId(result.UploadId)

	// assert.Equal(t, "bucket-us-west", upload.Bucket)
}

func TestRecordedPartHasDefaultBucketWhenInValidSubdomain(t *testing.T) {
	// req, _ := http.NewRequest("POST", "http://example.com/uploads/path/to/my/object", nil)
	// addHeaders(req)

	// db := domain.NewInMemoryDatabase()
	// controller := &UploadController{}

	// _, response := controller.InitiateMultipartUpload(req, db)
	// result := GetResultMultipartUploadResult(response)

	// upload := db.GetUploadByUploadId(result.UploadId)

	// assert.Equal(t, "default", upload.Bucket)
}
