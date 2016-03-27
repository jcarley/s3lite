package test

import "github.com/jcarley/s3lite/domain"

type MockBucketDatastore struct {
	MockService
}

func NewMockBucketDatastore() *MockBucketDatastore {
	return &MockBucketDatastore{
		MockService{
			callChain:     make(map[string]struct{ Count int }),
			methodWatches: make(map[string]*MethodWatch),
		}}
}

func (this *MockBucketDatastore) GetBucketById(bucketId string) *domain.Bucket {
	return nil
}

func (this *MockBucketDatastore) CreateBucket(bucket *domain.Bucket) (id string, err error) {

	methodName := "CreateBucket"

	this.AddCallChainFunc(methodName)

	if methodWatch, ok := this.MethodWatches()[methodName]; ok {
		this.GetReturnArg(methodWatch, 0, &id, nil)
		this.GetReturnArg(methodWatch, 1, &err, nil)
		return
	}

	return "", nil
}

func (this *MockBucketDatastore) DeleteBucketById(bucketId string) (err error) {

	methodName := "DeleteBucketById"

	this.AddCallChainFunc(methodName)

	if methodWatch, ok := this.MethodWatches()[methodName]; ok {
		this.GetReturnArg(methodWatch, 0, &err, nil)
		return
	}

	return nil
}
