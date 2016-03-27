package services

import (
	"errors"
	"fmt"
	"testing"

	"github.com/jcarley/s3lite/domain"
	"github.com/jcarley/s3lite/test"
	. "github.com/onsi/gomega"
)

// type BucketDatastore interface {
// GetBucketById(bucketId string) *Bucket
// CreateBucket(bucket *Bucket) (string, error)
// }

func TestAddBucket(t *testing.T) {
	RegisterTestingT(t)

	cases := []struct {
		SuppliedBucket *domain.Bucket
		ExpectedId     string
		ExpectedErr    error
	}{
		{&domain.Bucket{Name: "bucket-us-west"}, "1234567890", nil},
		{&domain.Bucket{Name: "bucket-us-west"}, "", errors.New("Unknown error")},
		{(*domain.Bucket)(nil), "", errors.New("Must supply a bucket")},
	}

	for _, tc := range cases {
		datastore := test.NewMockBucketDatastore()
		datastore.On("CreateBucket").Return(tc.ExpectedId, tc.ExpectedErr)

		service := NewBucketService(datastore)
		bucket := tc.SuppliedBucket
		err := service.AddBucket(bucket)

		if bucket != nil {
			Expect(bucket.Id).To(Equal(tc.ExpectedId), fmt.Sprintf("Should have Id equal to %s", tc.ExpectedId))
		}
		if tc.ExpectedErr == nil {
			Expect(err).To(BeNil(), "Should have a nil error")
		} else {
			Expect(err).To(MatchError(tc.ExpectedErr), fmt.Sprintf("Should have received error %s", tc.ExpectedErr.Error()))
		}
	}
}
