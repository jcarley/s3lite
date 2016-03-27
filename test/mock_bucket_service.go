package test

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/jcarley/s3lite/domain"
)

type MockBucketService struct {
	MockService
}

func NewMockBucketService() *MockBucketService {
	return &MockBucketService{
		MockService{
			callChain:     make(map[string]struct{ Count int }),
			methodWatches: make(map[string]*MethodWatch),
		}}
}

func (this *MockBucketService) AddBucket(bucket *domain.Bucket) error {

	this.AddCallChainFunc("AddBucket")

	data := []byte("1234567890")
	hash := md5.Sum(data)
	bucket.Id = hex.EncodeToString(hash[:])

	return nil
}

func (this *MockBucketService) DeleteBucketById(id string) (err error) {
	methodName := "DeleteBucketById"

	this.AddCallChainFunc(methodName)

	if methodWatch, ok := this.MethodWatches()[methodName]; ok {
		this.GetReturnArg(methodWatch, 0, &err, nil)
		return
	}

	return nil
}
