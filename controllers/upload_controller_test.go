package controllers

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	services "github.com/jcarley/s3lite/services/testing"
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

func GetUploadController() *UploadController {
	mockUploadService := services.NewMockUploadService()
	return NewUploadController(mockUploadService)
}

func TestCreateMultipartUploadReturnsAnUploadId(t *testing.T) {
	RegisterTestingT(t)

	controller := GetUploadController()

	req, _ := http.NewRequest("POST", "http://bucket-us-west.s3.example.com/uploads/path/to/my/object", nil)
	addHeaders(req)

	w := httptest.NewRecorder()

	controller.CreateMultipartUpload(w, req)

	data := GetRawData(t, w.Body.Bytes())

	Expect(data["upload_id"]).ToNot(BeNil(), "Should have an upload id")
	Expect(w.Code).To(Equal(http.StatusOK), "Should receive 200 status")
}

func TestCreateMultipartUploadCreatesAnUploadRecord(t *testing.T) {
	RegisterTestingT(t)

	controller := GetUploadController()

	req, _ := http.NewRequest("POST", "http://bucket-us-west.s3.example.com/uploads/path/to/my/object", nil)
	addHeaders(req)

	w := httptest.NewRecorder()

	controller.CreateMultipartUpload(w, req)

	service := controller.service.(*services.MockUploadService)

	Expect(service.Called("CreateUpload").Times(1)).To(BeTrue())
}

func TestUploadPartHasRequiredHeaders(t *testing.T) {
	RegisterTestingT(t)

	cases := []struct {
		Header string
		Err    error
	}{
		{"Upload-Id", MissingUploadIdError},
		{"Part-Number", InvalidPartNumberError},
		{"Authorization", MissingAuthorizationKeyError},
		{"Content-Disposition", MissingContentDispostionError},
	}

	for _, tc := range cases {
		controller := GetUploadController()
		req, _ := http.NewRequest("PUT", "http://bucket-us-west.s3.example.com/uploads/path/to/my/object", nil)
		addHeaders(req)
		req.Header.Del(tc.Header)
		w := httptest.NewRecorder()
		controller.UploadPart(w, req)
		data := GetRawData(t, w.Body.Bytes())
		Expect(data["message"]).To(Equal(tc.Err.Error()))
	}

}

func TestUploadPartHasBucket(t *testing.T) {
	RegisterTestingT(t)

	cases := []struct {
		Url      string
		Expected string
	}{
		{"http://bucket-us-west.s3.example.com/uploads/path/to/my/object", "bucket-us-west"},
		{"http://example.com/uploads/path/to/my/object", "default"},
	}

	for _, tc := range cases {
		req, _ := http.NewRequest("PUT", tc.Url, nil)
		controller := GetUploadController()
		actual := controller.parseBucket(req)
		Expect(tc.Expected).To(Equal(actual))
	}
}

func TestUploadPartRequiresBody(t *testing.T) {
	RegisterTestingT(t)

	cases := []struct {
		Reader io.Reader
	}{
		{strings.NewReader("")},
		{nil},
	}

	for _, tc := range cases {
		reader := tc.Reader
		req, _ := http.NewRequest("PUT", "http://bucket-us-west.s3.example.com/uploads/path/to/my/object", reader)
		addHeaders(req)
		w := httptest.NewRecorder()

		controller := GetUploadController()
		controller.UploadPart(w, req)

		data := GetRawData(t, w.Body.Bytes())
		Expect(data["message"]).To(Equal(MissingContentBodyError.Error()))
	}

}

func TestUploadPartAddsThePart(t *testing.T) {
	RegisterTestingT(t)

	body := strings.NewReader("AAAAAAAAAAAAAAAAAAAAAAAA")
	req, _ := http.NewRequest("PUT", "http://bucket-us-west.s3.example.com/uploads/path/to/my/object", body)
	addHeaders(req)

	w := httptest.NewRecorder()

	controller := GetUploadController()
	controller.UploadPart(w, req)

	service := controller.service.(*services.MockUploadService)

	Expect(service.Called("AddPart").Times(1)).To(BeTrue())
}

func TestUploadPartReturnsAnEtag(t *testing.T) {
	RegisterTestingT(t)

	body := strings.NewReader("AAAAAAAAAAAAAAAAAAAAAAAA")
	req, _ := http.NewRequest("PUT", "http://bucket-us-west.s3.example.com/uploads/path/to/my/object", body)
	addHeaders(req)

	w := httptest.NewRecorder()

	controller := GetUploadController()

	service := controller.service.(*services.MockUploadService)
	service.On("AddPart").Return("12345", nil)

	controller.UploadPart(w, req)

	data := GetRawData(t, w.Body.Bytes())

	Expect(data["status"]).To(Equal("success"))
	Expect(data["message"]).To(Equal("12345"))
}
