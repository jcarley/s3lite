package services

import "github.com/jcarley/s3lite/domain"

type InvalidArgumentError string

func (i InvalidArgumentError) Error() string {
	return "invalid argument: " + string(i)
}

type BucketServicer interface {
	AddBucket(bucket *domain.Bucket) error
	DeleteBucketById(id string) error
}

type BucketService struct {
	datastore domain.BucketDatastore
}

func NewBucketService(datastore domain.BucketDatastore) *BucketService {
	return &BucketService{
		datastore: datastore,
	}
}

func (this *BucketService) AddBucket(bucket *domain.Bucket) error {
	if bucket == nil {
		return InvalidArgumentError("bucket can not be nil")
	}

	id, err := this.datastore.CreateBucket(bucket)
	if err != nil {
		return err
	}

	bucket.Id = id

	return nil
}

func (this *BucketService) DeleteBucketById(id string) error {

	err := this.datastore.DeleteBucketById(id)
	if err != nil {
		return err
	}

	return nil
}
