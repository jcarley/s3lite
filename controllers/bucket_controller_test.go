package controllers

import (
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

func TestCreateBucketReturnsAnBucketId(t *testing.T) {
	RegisterTestingT(t)

	bucket := &domain.Bucket{
		"bucket-us-west",
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
	Expect(w.Code).To(Equal(http.StatusOK), "Should receive 200 status")
}
