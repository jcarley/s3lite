package controllers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jcarley/s3lite/domain"
	"github.com/jcarley/s3lite/test"
	. "github.com/onsi/gomega"
)

func GetBucketController() *BucketController {
	mockBucketService := test.NewMockBucketService()
	return NewBucketController(mockBucketService)
}

func TestCreateBucketReturnsBucketId(t *testing.T) {
	RegisterTestingT(t)

	bucket := &domain.Bucket{
		Name: "bucket-us-west",
	}

	bucketBytes := test.SetRawData(t, bucket)

	reader := strings.NewReader(string(bucketBytes))
	req, _ := http.NewRequest("POST", "http://s3.example.com/buckets", reader)

	addHeaders(req)

	w := httptest.NewRecorder()

	controller := GetBucketController()
	controller.CreateBucket(w, req)

	data := test.GetRawData(t, w.Body.Bytes())

	Expect(data["bucket_id"]).ToNot(BeNil(), "Should have an bucket id")
	Expect(data["name"]).To(Equal("bucket-us-west"), "Should have an bucket id")
	Expect(w.Code).To(Equal(http.StatusOK), "Should receive 200 status")
}

func TestDeleteBucketRemovesBucket(t *testing.T) {
	RegisterTestingT(t)

	id := "1234567890"

	path := fmt.Sprintf("http://s3.example.com/buckets/%s", id)
	req, _ := http.NewRequest("DELETE", path, nil)

	addHeaders(req)

	w := httptest.NewRecorder()

	controller := GetBucketController()
	controller.DeleteBucket(w, req)

	data := test.GetRawData(t, w.Body.Bytes())

	Expect(data["status"]).To(Equal("success"), "Should have a success status")
	Expect(w.Code).To(Equal(http.StatusOK), "Should receive 200 status")
}
